package pow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
)

//job is the response from the pool server contianing job data for this worker
type Job struct {
	Challenge     string `json:"challenge"`
	Difficulty    uint64    `json:"difficulty"`
	EndNonce      uint64    `json:"end_nonce"`
	JobID         uint64    `json:"job_id"`
	PublicAddress string `json:"public_address"`
	StartNonce    uint64    `json:"start_nonce"`
	WorkID        uint64    `json:"work_id"`
}

type Pool struct {
	log           *util.Logger
	url           string
	publicAddress string
	currJobID     uint64
	group         *MiningGroup
	jobDuration   config.Duration

}

func CreatePool(cfg *config.Config, group *MiningGroup) *Pool {

	return &Pool{
		url:           cfg.PoolURL,
		currJobID:     0,
		publicAddress: cfg.PublicAddress,
		log:           util.NewLogger("pow", "Pool"),
		group:         group,
		jobDuration:   cfg.PoolJobDuration,
	}
}

func (p *Pool) GetWork() *Work {
	// HTTP GET work from the server
	if p.currJobID > 0 {
		return nil
	}
	//target 15s pool chunks by default
	jobSize := p.jobDuration.Seconds()*p.group.HashRateEstimate()
	step := p.group.PreferredWorkMultiple()
	nsteps := uint64(math.Round(jobSize /float64(step)))
	if nsteps == 0 {
		nsteps = 1
	}
	n := nsteps * step

	url := p.url + fmt.Sprintf("/job?public_address=%s&job_size=%d", p.publicAddress, n)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		p.log.Error("failed to get work from pool: %s", err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.log.Error("failed to read response: %s", err.Error())
		return nil
	}

	var j = new(Job)
	err = json.Unmarshal(body, &j)
	if err != nil {
		p.log.Error("Error decoding job json: %s", err.Error())
		return nil
	}
	diff := new(big.Int)
	diff.SetUint64(j.Difficulty)
	newChallenge := &MiningChallenge{
		Challenge:  decodeHex(j.Challenge),
		Difficulty: diff,
		RequestID:  big.NewInt(1),
	}
	p.currJobID = j.JobID
	return &Work{Challenge: newChallenge, PublicAddr: j.PublicAddress, Start: j.StartNonce, N: j.EndNonce - j.StartNonce}
}

func (p *Pool) Submit(ctx context.Context, result *Result) {
	// HTTP POST valid nonces to the pool server
	nonce := result.Nonce
	if result.Nonce == "" {
		nonce = "0"
	}
	values := map[string]string{
		"job_id": fmt.Sprintf("%v", p.currJobID),
		"nonce":  nonce,
	}
	jsonValue, _ := json.Marshal(values)

	url := p.url + "/job"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.SetBasicAuth("demo", "demo")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		p.log.Error("Error posting nonce:", err.Error())
	} else {
		resp.Body.Close()
	}
	p.currJobID = 0
}
