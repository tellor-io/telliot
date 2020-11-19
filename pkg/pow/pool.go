// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"context"
	"encoding/json"
	"math"
	"math/big"
	"math/rand"
	"strconv"

	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

type StratumPool struct {
	log           *util.Logger
	url           string
	minerAddress  string
	minerPassword string
	group         *MiningGroup
	stratumClient *StratumClient
	input         chan *Work
	currChallenge *MiningChallenge
	currWork      *Work
	currJobID     string
}

type MiningNotify struct {
	JobID             string
	Challenge         string
	PoolAddress       string
	LowDifficulty     *big.Int
	MedianDifficulty  *big.Int
	NetworkDifficulty *big.Int
	CleanJob          bool
}

type MiningSetDifficulty struct {
	Difficulty *big.Int
}

func (n *MiningNotify) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.JobID, &n.Challenge, &n.PoolAddress, &n.LowDifficulty, &n.MedianDifficulty, &n.NetworkDifficulty, &n.CleanJob}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return errors.Errorf("wrong number of fields in MiningNotify: %v != %v", g, e)
	}
	return nil
}

func (n *MiningSetDifficulty) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.Difficulty}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return errors.Errorf("wrong number of fields in MiningSetDifficulty: %v != %v", g, e)
	}
	return nil
}

func CreatePool(cfg *config.Config, group *MiningGroup) *StratumPool {
	return &StratumPool{
		url:           cfg.PoolURL,
		minerAddress:  cfg.PublicAddress + "." + cfg.Worker,
		minerPassword: cfg.Password,
		log:           util.NewLogger("pow", "StratumPool"),
		group:         group,
	}
}

func (p *StratumPool) GetWork(input chan *Work) (*Work, bool) {
	if p.stratumClient != nil && p.stratumClient.running {
		p.log.Warn("stratum client already running")
		return nil, false
	}

	p.input = input

	msgChan := make(chan *StratumResponse)
	stratumClient, err := StratumConnect(p.url, msgChan)
	if err != nil {
		p.log.Error("stratum connect error: %s", err.Error())
		return nil, false
	}

	p.stratumClient = stratumClient
	p.stratumClient.Request(
		"mining.subscribe",
		"TellorStratum/1.0.0")

	subscribed := false
	nonce1 := ""

	go func() {
		for msg := range msgChan {
			if !subscribed {
				r, err := json.Marshal(msg.Result)
				if err != nil {
					p.log.Error("parse subscribe result error: %s", err.Error())
					return
				}
				result := string(r)
				nonce1 = fmt.Sprintf("%x", []byte(result[7:15]))
				subscribed = true

				p.stratumClient.Request(
					"mining.authorize",
					p.minerAddress, p.minerPassword)
			}

			if msg.Method == "mining.notify" {
				params, err := json.Marshal(msg.Params)
				if err != nil {
					p.log.Error("mining.notify msg parse error: %s", err.Error())
					return
				}

				var miningNotify MiningNotify
				if err := json.Unmarshal([]byte(string(params)), &miningNotify); err != nil {
					p.log.Error("mining.notify params msg parse error: %s", err.Error())
				}

				p.log.Info("mining.notify: %#v", miningNotify)

				newChallenge := &MiningChallenge{
					Challenge:  decodeHex(miningNotify.Challenge),
					Difficulty: miningNotify.MedianDifficulty,
					// Difficulty: big.NewInt(10000000),
					// Difficulty: big.NewInt(6377077812),
					RequestIDs: [5]*big.Int{big.NewInt(1)},
				}

				p.currChallenge = newChallenge
				p.currJobID = miningNotify.JobID
				job := &Work{
					Challenge:  newChallenge,
					PublicAddr: miningNotify.PoolAddress + nonce1,
					Start:      uint64(rand.Int63()),
					N:          math.MaxInt64}
				p.currWork = job
				input <- job

			} else if msg.Method == "mining.set_difficulty" {
				// Not implmented
				p.log.Error("mining.set_difficulty not implemented")
			}
		}
	}()

	return nil, false
}

func (p *StratumPool) Submit(ctx context.Context, result *Result) (*types.Transaction, error) {
	nonce := result.Nonce
	p.stratumClient.Request(
		"mining.submit",
		p.minerAddress,
		fmt.Sprintf("%v", p.currJobID), nonce)

	if p.input != nil {
		result.Work.Start = uint64(rand.Int63())
		p.input <- result.Work
	}

	noncePrs, err := strconv.ParseUint(result.Nonce, 0, 64)
	if err != nil {
		return nil, err
	}
	return types.NewTransaction(noncePrs, common.HexToAddress(p.minerAddress), big.NewInt(0), 0, big.NewInt(0), []byte{}), nil
}
