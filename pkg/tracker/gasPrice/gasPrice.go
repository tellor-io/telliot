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
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/http"
)

const ComponentName = "gasTracker"

// GasTracker is the struct that maintains the latest gasprices.
// note the prices are actually stored in the DB.
type GasTracker struct {
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

func New(logger log.Logger, client contracts.ETHClient) *GasTracker {
	return &GasTracker{
		client: client,
		logger: log.With(logger, "component", ComponentName),
	}

}

func (self *GasTracker) Query(ctx context.Context) (int64, error) {
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

	return gasPrice.Int64(), nil
}
