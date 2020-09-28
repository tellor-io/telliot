// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/tellor-io/TellorMiner/util"
)

type StratumRequest struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     uint64   `json:"id"`
}

type StratumResponse struct {
	ID     uint64           `json:"id"`
	Result *json.RawMessage `json:"result"`
	Error  error
	// Method is used in notifications.
	Method string           `json:"method"`
	Params *json.RawMessage `json:"params"`
}

type StratumClient struct {
	socket  net.Conn
	seq     uint64
	timeout time.Duration
	log     *util.Logger
	msgChan chan *StratumResponse
	running bool
}

func StratumConnect(host string, msgChan chan *StratumResponse) (*StratumClient, error) {
	var client StratumClient
	var err error
	client.socket, err = net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	client.seq = 0
	client.msgChan = msgChan
	client.log = util.NewLogger("pow", "Pool")
	client.SetTimeout(10)
	client.log.Info("connect to pool success")
	client.running = true
	go client.Listen()
	return &client, nil
}

func (c *StratumClient) SetTimeout(timeout int64) {
	c.timeout = time.Duration(timeout) * time.Second
}

func (c *StratumClient) Listen() {
	defer func() {
		if err := c.socket.Close(); err != nil {
			fmt.Println("error closing the connection", err)
		}
	}()
	if err := c.socket.SetReadDeadline(time.Time{}); err != nil {
		fmt.Println("error setting connection deadline", err)
	}

	for {
		result, err := bufio.NewReader(c.socket).ReadString('\n')
		if err != nil {
			c.log.Error("failed to read: %s", err.Error())
			c.running = false
			break
		}

		response := &StratumResponse{}
		// c.log.Info("get response from pool %s", result)
		err = json.Unmarshal([]byte(result), &response)
		if err != nil {
			c.log.Error("failed to get response from pool: %s", err.Error())
			continue
		}
		// c.log.Info("get response : %v", response)
		c.msgChan <- response
	}
}

func (c *StratumClient) Send(request *StratumRequest) *StratumResponse {
	c.seq++
	request.ID = c.seq

	if request.Params == nil {
		request.Params = make([]string, 0)
	}
	encoded, err := json.Marshal(request)
	msg := string(encoded) + "\n"

	if err != nil {
		return &StratumResponse{Error: err}
	}
	// .log.Info("send msg to pool: %s", msg)
	_, err = c.socket.Write([]byte(msg))
	if err != nil {
		c.log.Error("failed to send msg to pool: %s", err.Error())
		return &StratumResponse{Error: err}
	}

	return &StratumResponse{ID: 1, Error: nil}
}

func (c *StratumClient) Request(method string, params ...string) *StratumResponse {
	return c.Send(&StratumRequest{Method: method, Params: params})
}
