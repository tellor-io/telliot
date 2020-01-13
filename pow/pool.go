package pow

import (
	"fmt"
  "context"
  "math/big"
  "bytes"
  "net/http"
  "io/ioutil"
  "encoding/json"
	"github.com/tellor-io/TellorMiner/config"
  "github.com/tellor-io/TellorMiner/util"
)

//job is the response from the pool server contianing job data for this worker
type Job struct {
	Challenge       string `json:"challenge"`
  Difficulty     int `json:"difficulty"`
  EndNonce       int `json:"end_nonce"`
  JobID          int `json:"job_id"`
	PublicAddress   string `json:"public_address"`
	StartNonce     int `json:"start_nonce"`
  WorkID        int `json:"work_id"`
}

type Pool struct {
	log              *util.Logger
  url string
  publicAddress string
  currJobID int
}

func CreatePool(cfg *config.Config) *Pool {

	return &Pool{
		url: cfg.PoolURL,
    currJobID: 0,
    publicAddress: cfg.PublicAddress,
		log:util.NewLogger("pow", "Pool"),
	}
}


func (p *Pool) GetWork() *Work {
  // HTTP GET work from the server
  p.log.Info("Getting work from the pool server.")
  if p.currJobID > 0 {
    return nil
  }
  var j = new(Job)
	url := p.url + fmt.Sprintf("/job?public_address=%s&job_size=10000000", p.publicAddress)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	cli := &http.Client{}
	resp, err := cli.Do(req)
	p.log.Info("response Status:", resp.Status)
	p.log.Info("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &j)
  p.log.Info("Response Body Job Data: %v", string(body))
  p.log.Info("Response Body Job Data: %v", j)
	if err != nil {
			p.log.Error("Error decoding nonce:", err.Error())
			return nil
	}
	newChallenge := &MiningChallenge{
		Challenge:  []byte(j.Challenge),
		Difficulty: big.NewInt(int64(j.Difficulty)),
		RequestID:  big.NewInt(1),
	}
	p.currJobID = j.JobID
	return &Work{Challenge:newChallenge, PublicAddr:j.PublicAddress, Start:uint64(j.StartNonce), N:uint64(j.EndNonce - j.StartNonce)}
}

func (p *Pool) Submit(ctx context.Context, result *Result) {
  // HTTP POST valid nonces to the pool server
  nonce := result.Nonce
  if result.Nonce == "" {
    nonce = "0"
  }
  values := map[string]string{
		"job_id": fmt.Sprintf("%v",p.currJobID),
		"nonce": nonce}
  p.log.Info(fmt.Sprintf("Submit: %v", values))
	jsonValue, _ := json.Marshal(values)

	url := p.url + "/job"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.SetBasicAuth("demo", "demo")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		p.log.Error("Problem submitting txn", err)
			p.log.Error("Error posting nonce:", err.Error())
			return
	} else {
		p.log.Info("Successfully submitted solution")
		p.log.Info("Posted job to pool server.")
		body, _ := ioutil.ReadAll(resp.Body)
		//ml.log.Info("response Status:", resp.Status)
	  //ml.log.Info("response Headers:", resp.Header)
	  p.log.Info("response Body:", string(body))
    p.currJobID = 0
	}
}
