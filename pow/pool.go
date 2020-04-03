package pow

import (
	// "bytes"
	"context"
	"encoding/json"
	"math"
	"math/rand"
	"math/big"

	"fmt"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
)

type StratumPool struct {
	log           *util.Logger
	url           string
	minerAddress  string
	worker        string
	group         *MiningGroup
	stratumClient *StratumClient
	input 				chan *Work
	currChallenge *MiningChallenge
	currJobID     string
}

type MiningNotify struct {
	JobID  			string
	Challenge 	string
	PoolAddress string
	Difficulty *big.Int
	CleanJob 		bool
}

func (n *MiningNotify) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.JobID, &n.Challenge, &n.PoolAddress, &n.Difficulty, &n.CleanJob}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in MiningNotify: %d != %d", g, e)
	}
	return nil
}


func CreatePool(cfg *config.Config, group *MiningGroup) *StratumPool {
	return &StratumPool{
		url:           cfg.PoolURL,
		minerAddress:  cfg.PublicAddress + "." + cfg.Worker,
		log:           util.NewLogger("pow", "StratumPool"),
		group:         group,
	}
}

func (p *StratumPool) GetWork(input chan *Work) *Work {
	if p.stratumClient != nil && p.stratumClient.running {
		p.log.Warn("stratum client already running")
		return nil
	}

	p.input = input

	msgChan := make(chan *StratumResponse)
	stratumClient, err := StratumConnect(p.url, msgChan)
	if err != nil {
		p.log.Error("stratum connect error: %s", err.Error())
		return nil
	}

	p.stratumClient = stratumClient
	p.stratumClient.Request(
		"mining.subscribe",
		"TellorStratum/1.0.0")

	subscribed := false
	nonce1 := ""

	go func() {
		for {
			select {
			case msg := <-msgChan:

				p.log.Info("received message from pool: %#v", msg)

				if !subscribed {
					r, err := json.Marshal(msg.Result)
					if err != nil {
						p.log.Error("parse subscribe result error: %s", err.Error())
						return
					}
					result := string(r)
					nonce1 = fmt.Sprintf("%x", []byte(result[7:15]))
					p.log.Info("nonce1 is : %v", nonce1)
					subscribed = true

					p.stratumClient.Request(
						"mining.authorize",
						p.minerAddress, "123")
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
						Difficulty: miningNotify.Difficulty,
						// Difficulty: big.NewInt(10000000),
						// Difficulty: big.NewInt(6377077812),
						RequestID:  big.NewInt(1),
					}

					p.currChallenge = newChallenge
					p.currJobID = miningNotify.JobID
					job := &Work{
						Challenge: newChallenge,
						PublicAddr: miningNotify.PoolAddress + nonce1,
						Start: uint64(rand.Int63()),
						N: math.MaxInt64}
					input <- job
					p.log.Info("send new work to hash %#v", job)
				}
			}
		}
	}()

	return nil
}

func (p *StratumPool) Submit(ctx context.Context, result *Result) {
	nonce := result.Nonce
	submission := fmt.Sprintf("%s, %s, %s", p.minerAddress, p.currJobID, nonce)
	p.log.Warn("mining.submit: %s", submission)
	p.stratumClient.Request(
		"mining.submit",
		p.minerAddress,
		fmt.Sprintf("%v", p.currJobID), nonce)

	if p.input != nil {
		result.Work.Start = uint64(rand.Int63())
		p.input <- result.Work
	}
}
