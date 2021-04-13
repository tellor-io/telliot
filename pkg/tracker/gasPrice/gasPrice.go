// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package gasPrice

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/timestamp"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/tsdb/tsdbutil"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/http"
)

const ComponentName = "gasTracker"

// GasTracker is the struct that maintains the latest gasprices.
// note the prices are actually stored in the DB.
type GasTracker struct {
	db     *tsdb.DB
	client contracts.ETHClient
	logger log.Logger
}

// GasPriceModel is what ETHGasStation returns from queries. Not all fields are filled in.
type GasPriceModel struct {
	Fast    float32 `json:"fast"`
	Fastest float32 `json:"fastest"`
	Average float32 `json:"average"`
}

func (self *GasTracker) String() string {
	return "GasTracker"
}

func New(logger log.Logger, db *tsdb.DB, client contracts.ETHClient) *GasTracker {
	return &GasTracker{
		db:     db,
		client: client,
		logger: log.With(logger, "component", ComponentName),
	}

}

func (self *GasTracker) set(ctx context.Context) (int64, error) {
	netID, err := self.client.NetworkID(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get network id")
	}

	var gasPrice *big.Int

	if big.NewInt(1).Cmp(netID) == 0 {
		ctx, cncl := context.WithTimeout(ctx, 15*time.Second)
		defer cncl()
		resp, err := http.Fetch(ctx, self.logger, "https://ethgasstation.info/json/ethgasAPI.json", time.Second)
		if err != nil {
			gasPrice, err = self.client.SuggestGasPrice(ctx)
			if err != nil {
				level.Warn(self.logger).Log("msg", "get suggested gas price", "err", err)
			}
		} else {
			gpModel := GasPriceModel{}
			err = json.Unmarshal(resp, &gpModel)
			if err != nil {
				level.Warn(self.logger).Log("msg", "eth gas station json", "err", err)
				gasPrice, err = self.client.SuggestGasPrice(ctx)
				if err != nil {
					level.Warn(self.logger).Log("msg", "getting suggested gas price", "err", err)
				}
			} else {
				gasPrice = big.NewInt(int64(gpModel.Fast / 10))
				gasPrice = gasPrice.Mul(gasPrice, big.NewInt(params.GWei))
				level.Info(self.logger).Log("msg", "using ETHGasStation fast price", "price", gasPrice)
			}
		}
	} else {
		gasPrice, err = self.client.SuggestGasPrice(ctx)
		if err != nil {
			level.Warn(self.logger).Log("msg", "getting suggested gas price", "err", err)
		}
	}

	appender := self.db.Appender(ctx)

	_, err = appender.Append(0, labels.Labels{db.GasPriceLabel}, timestamp.FromTime(time.Now().Round(0)), float64(gasPrice.Int64()))
	if err != nil {
		return 0, errors.Wrap(err, "append to db")
	}

	if err := appender.Commit(); err != nil {
		return 0, err
	}

	return gasPrice.Int64(), nil
}

func (self *GasTracker) Query(ctx context.Context, mint time.Time) (float64, error) {
	q, err := self.db.Querier(ctx, timestamp.FromTime(mint.Round(0)), timestamp.FromTime(time.Now()))
	if err != nil {
		return 0, err
	}
	defer q.Close()
	s := q.Select(false, nil, labels.MustNewMatcher(labels.MatchEqual, db.GasPriceLabel.Name, db.GasPriceLabel.Value))

	var i int
	var samples []tsdbutil.Sample
	for s.Next() {
		if i > 0 {
			return 0, errors.New("returned more then one series")
		}
		series := s.At()
		_samples, err := storage.ExpandSamples(series.Iterator(), nil)
		if err != nil {
			return 0, err
		}
		samples = _samples
		i++
	}

	if len(samples) == 0 {
		res, err := self.set(ctx)
		if err != nil {
			return 0, errors.Wrap(err, "getting fresh gas price")
		}
		return float64(res), nil
	}

	return samples[len(samples)-1].V(), nil
}
