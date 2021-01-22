// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/testutil"
	"github.com/tellor-io/telliot/pkg/util"
)

func TestAmpl(t *testing.T) {
	util.CreateTestClient(&client, mockAPI)
	cfg := config.OpenTestConfig(t)
	DB, cleanup := db.OpenTestDB(t)
	testClient := rpc.NewMockClient()
	defer t.Cleanup(cleanup)

	proxy, err := db.OpenLocal(cfg, DB)
	testutil.Ok(t, err)

	mock := clock.NewMock()
	clck = mock
	mock.Set(time.Now())
	if _, err := BuildIndexTrackers(cfg, proxy, testClient); err != nil {
		testutil.Ok(t, err)
	}
	amplTrackers := indexes["AMPL/USD"]
	btcTrackers := indexes["BTC/USD"]
	amplBtcTrackers := indexes["AMPL/BTC"]

	indexers := []*IndexTracker{}
	indexers = append(indexers, amplTrackers...)
	indexers = append(indexers, btcTrackers...)
	indexers = append(indexers, amplBtcTrackers...)
	for i := 0; i < 288; i++ {
		for _, indexer := range indexers {
			// Ignore on-chain trackers, as they could be tested in other test cases.
			if indexer.Type == ethereumIndexType {
				continue
			}
			if err := indexer.Exec(context.Background()); err != nil {
				testutil.Ok(t, err)
			}
		}
		mock.Add(10 * time.Minute)
	}

	// reset mocks
	client = http.Client{}
	clck = clock.New()
}

func contains(req *http.Request, segment string) bool {
	return strings.Contains(req.URL.String(), segment)
}

func mockAPI(req *http.Request) *http.Response {
	var reqBody string
	switch {
	// Ampleforth/BTC
	case contains(req, "https://api.kucoin.com/api/v1/market/stats?symbol=AMPL-BTC"):
		reqBody = `{"code":"200000","data":{"time":1598943308068,"symbol":"AMPL-BTC","buy":"0.00018337","sell":"0.00018422","changeRate":"0.2536","changePrice":"0.00003726","high":"0.00019398","low":"0.00013484","vol":"2104825.78672","volValue":"354.2575905596620409","last":"0.00018416","averagePrice":"0.00015212"}}`
	case contains(req, "https://api-pub.bitfinex.com/v2/tickers?symbols=tAMPBTC"):
		reqBody = `[["tAMPBTC",0.00017945,35968.96882899,0.00018267,22481.02421014,0.00003498,0.2375,0.00018226,201846.40489022,0.00019297,0.00013474]]`
	// Ampleforth/USD
	case contains(req, "https://api.coingecko.com/api/v3/simple/price?ids=ampleforth&vs_currencies=usd"):
		reqBody = `{"ampleforth":{"usd":2.15}}`
	case contains(req, "https://api.kucoin.com/api/v1/market/stats?symbol=AMPL-USDT"):
		reqBody = `{"code":"200000","data":{"time":1598899422047,"symbol":"AMPL-USDT","buy":"2.15601","sell":"2.16","changeRate":"0.3035","changePrice":"0.50301","high":"2.21758","low":"1.56113","vol":"9111879.72267762","volValue":"16728059.486419401907","last":"2.16","averagePrice":"1.79088323"}}`
	case contains(req, "https://api-pub.bitfinex.com/v2/tickers?symbols=tAMPUSDT"):
		reqBody = `[["tAMPUST",2.1658,30036.857127949996,2.1697,37123.11378477,0.5024,0.3031,2.16,106302.90115003,2.2746,1.5466]]`
	case contains(req, "https://api-pub.bitfinex.com/v2/tickers?symbols=tAMPUSD"):
		reqBody = `[["tAMPUSD",1.9711,30215.46857231,2.0035,34257.8592384,0.1128,0.0606,1.9729,558433.75188253,2.255,1.5608]]`
	case contains(req, "https://min-api.cryptocompare.com/data/price?fsym=AMPL&tsyms=USD"):
		reqBody = `{"USD":2.193}`
	case contains(req, "https://api-pub.bitfinex.com/v2/tickers?symbols=tAMPUST"):
		reqBody = `[["tAMPUST",0.6608,47736.78043824,0.66429,74484.21865398,-0.00608,-0.0092,0.65544,75030.61631492,0.67534,0.62829]]`
	// Bitcoin
	case contains(req, "https://api.pro.coinbase.com/products/BTC-USD/ticker"):
		reqBody = `{"trade_id":101831363,"price":"11694.23","size":"0.02110914","time":"2020-09-01T03:58:00.728546Z","bid":"11693.47","ask":"11694.23","volume":"10139.22817438"}`
	case contains(req, "https://api.binance.com/api/v1/klines?symbol=BTCUSDT&interval=1d&limit=1"):
		reqBody = `[[1598918400000,"11649.51000000","11703.00000000","11515.00000000","11689.87000000","9417.31309700",1599004799999,"109431843.37221315",160621,"4348.88606800","50544369.97251384","0"]]`
	case contains(req, "https://api.coindesk.com/v1/bpi/currentprice.json"):
		reqBody = `{"time":{"updated":"Sep 1, 2020 03:56:00 UTC","updatedISO":"2020-09-01T03:56:00+00:00","updateduk":"Sep 1, 2020 at 04:56 BST"},"disclaimer":"This data was produced from the CoinDesk Bitcoin Price Index (USD). Non-USD currency data converted using hourly conversion rate from openexchangerates.org","chartName":"Bitcoin","bpi":{"USD":{"code":"USD","symbol":"&#36;","rate":"11,687.4855","description":"United States Dollar","rate_float":11687.4855},"GBP":{"code":"GBP","symbol":"&pound;","rate":"8,928.9000","description":"British Pound Sterling","rate_float":8928.9},"EUR":{"code":"EUR","symbol":"&euro;","rate":"9,908.8489","description":"Euro","rate_float":9908.8489}}}`
	case contains(req, "https://www.bitstamp.net/api/v2/ticker/btcusd"):
		reqBody = `{"high": "11780.00", "last": "11698.78", "timestamp": "1598932679", "bid": "11691.96", "vwap": "11676.81", "volume": "4914.20345226", "low": "11515.00", "ask": "11699.08", "open": "11658.48"}`
	case contains(req, "https://api.coinpaprika.com/v1/tickers/btc-bitcoin"):
		reqBody = `{"id":"btc-bitcoin","name":"Bitcoin","symbol":"BTC","rank":1,"circulating_supply":18476362,"total_supply":18476375,"max_supply":21000000,"beta_value":1.01929,"last_updated":"2020-09-01T03:57:04Z","quotes":{"USD":{"price":11691.65276996,"volume_24h":17505564862.454,"volume_24h_change_24h":11.25,"market_cap":216019208956,"market_cap_change_24h":-0.06,"percent_change_15m":-0.06,"percent_change_30m":0.2,"percent_change_1h":0.47,"percent_change_6h":0.33,"percent_change_12h":-0.52,"percent_change_24h":-0.06,"percent_change_7d":-0.15,"percent_change_30d":-2.3,"percent_change_1y":18.8,"ath_price":20089,"ath_date":"2017-12-17T12:19:00Z","percent_from_price_ath":-41.8}}}`
	case contains(req, "https://api.cryptowat.ch/markets/coinbase-pro/btcusd/price"):
		reqBody = `{"result":{"price":11693.93},"allowance":{"cost":4063246,"remaining":3950410844,"upgrade":"For unlimited API access, create an account at https://cryptowat.ch"}}`
	case contains(req, "https://api.kraken.com/0/public/Ticker?pair=XBTUSD"):
		reqBody = `{"error":[],"result":{"XXBTZUSD":{"a":["11688.10000","1","1.000"],"b":["11688.00000","19","19.000"],"c":["11688.00000","0.00004306"],"v":["533.73252072","2539.33630153"],"p":["11623.45636","11677.73296"],"t":[2108,11581],"l":["11532.80000","11532.80000"],"h":["11700.00000","11785.00000"],"o":"11656.90000"}}}`
	case contains(req, "https://api-pub.bitfinex.com/v2/tickers?symbols=tBTCUSD"):
		reqBody = `[["tBTCUSD",11712,77.33582934,11713,95.85281801,-11.02013226,-0.0009,11712.97986774,2611.46682368,11790,11552]]`
	case contains(req, "api.coingecko.com") && contains(req, "bitcoin"):
		reqBody = `{"bitcoin":{"usd":11684.26}}`
	default:
		reqBody = `{"USD":2.193}`
	}
	// Test request parameters
	return &http.Response{
		StatusCode: 200,
		// Send response to be tested
		Body: ioutil.NopCloser(bytes.NewBufferString(reqBody)),
		// Must be set to non-nil value or it panics
		Header: make(http.Header),
	}
}
