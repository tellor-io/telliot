// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package gasStation

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/format"
	"github.com/tellor-io/telliot/pkg/web"
)

const ComponentName = "gasPriceGasStation"

type Config struct {
	TimeWait format.Duration
}

type GasStation struct {
	netID  int64
	cfg    Config
	client *ethclient.Client
	logger log.Logger
}

// GasStation is what ETHGasStation returns from queries. Not all fields are filled in.
type GasStationModel struct {
	Fast    float32 `json:"fast"`
	Fastest float32 `json:"fastest"`
	Average float32 `json:"average"`
}

func New(logger log.Logger, cfg Config, client *ethclient.Client) (*GasStation, error) {
	ctx, cncl := context.WithTimeout(context.Background(), 15*time.Second)
	defer cncl()
	netID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get network id")
	}

	return &GasStation{
		netID:  netID.Int64(),
		cfg:    cfg,
		client: client,
		logger: log.With(logger, "component", ComponentName),
	}, nil

}

func (self *GasStation) Query(ctx context.Context) (gasPriceFinal *big.Int, errFinal error) {
	if self.netID != 1 {
		gasPrice, err := self.client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "getting suggested gas price")
		}
		return gasPrice, nil
	}

	defer func() {
		if errFinal != nil {
			level.Error(self.logger).Log("msg", "fetching eth gas price falling back to client suggested price", "err", errFinal)
			gasPrice, err := self.client.SuggestGasPrice(ctx)
			if err != nil {
				errFinal = errors.Wrapf(errFinal, "failed to get price from chain client:%v", err)
				return
			}
			gasPriceFinal = gasPrice
		}
	}()

	ctx, cncl := context.WithTimeout(ctx, 15*time.Second)
	defer cncl()
	resp, err := web.Fetch(ctx, "https://ethgasstation.info/json/ethgasAPI.json")
	if err != nil {
		return nil, errors.Wrap(err, "fetch price from provider")
	}

	gpModel := GasStationModel{}
	err = json.Unmarshal(resp, &gpModel)
	if err != nil {
		return nil, errors.Wrap(err, "provider response json unmarshal")
	}

	var gasPrice float32
	switch t := self.cfg.TimeWait.Duration; {
	case t < 5*time.Minute:
		gasPrice = gpModel.Average
	case t < 2*time.Minute:
		gasPrice = gpModel.Fast
	case t < 30*time.Second:
		gasPrice = gpModel.Fastest
	}

	gasPriceB := big.NewInt(int64(gasPrice / 10))
	return big.NewInt(0).Mul(gasPriceB, big.NewInt(params.GWei)), nil
}
