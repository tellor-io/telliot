// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package gasPrice

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
	"github.com/tellor-io/telliot/pkg/web"
)

const ComponentName = "gasTracker"

// GasTracker is the struct that maintains the latest gasprices.
// note the prices are actually stored in the DB.
type GasTracker struct {
	netID  int64
	client *ethclient.Client
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

func New(logger log.Logger, client *ethclient.Client) (*GasTracker, error) {
	ctx, cncl := context.WithTimeout(context.Background(), 15*time.Second)
	defer cncl()
	netID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get network id")
	}

	return &GasTracker{
		netID:  netID.Int64(),
		client: client,
		logger: log.With(logger, "component", ComponentName),
	}, nil

}

func (self *GasTracker) Query(ctx context.Context) (int64, error) {

	var gasPrice *big.Int
	var err error

	if self.netID == 1 {
		ctx, cncl := context.WithTimeout(ctx, 15*time.Second)
		defer cncl()
		resp, err := web.Fetch(ctx, "https://ethgasstation.info/json/ethgasAPI.json")
		if err != nil {
			level.Error(self.logger).Log("msg", "fetching eth gas price falling back to client suggested price", "err", err)
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
