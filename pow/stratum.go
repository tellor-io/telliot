package pow

import (
    "net";
    "bufio";
    "time";
    "encoding/json"
    "github.com/tellor-io/TellorMiner/util"

    // "strconv"
    // "io/ioutil"
)

type StratumRequest struct{
    Method string         `json:"method"`
    Params []string       `json:"params"`
    ID     uint64         `json:"id"`
}


/*method, params is used in notifications*/

type StratumResponse struct{
    ID     uint64             `json:"id"`
    Result *json.RawMessage   `json:"result"`
    Error  error              `json:"error,string"`
    Method string             `json:"method"`
    Params *json.RawMessage   `json:"params"`
}

type StratumClient struct
{
    socket net.Conn
    seq uint64
    timeout time.Duration
    log *util.Logger
    msgChan chan *StratumResponse
    running bool
}

func StratumConnect(host string, msgChan chan *StratumResponse) (*StratumClient, error) {
    var client StratumClient
    var err error
    client.socket, err = net.Dial("tcp", host)
    if err!=nil {return nil,err}
    client.seq = 0
    client.msgChan = msgChan
    client.log = util.NewLogger("pow", "Pool")
    client.SetTimeout(10)
    client.log.Info("connect to pool success")
    client.running = true
    go client.Listen()
    return &client, nil
}

func (c *StratumClient) SetTimeout(timeout int64){
    c.timeout = time.Duration(timeout) * time.Second
}

func (c *StratumClient) Listen(){
    defer c.socket.Close()
    c.socket.SetReadDeadline(time.Time{})
    for {
        result, err := bufio.NewReader(c.socket).ReadString('\n')
        if err !=nil {
            c.log.Error("failed to read: %s", err.Error())
            c.running = false
            break
        }

        response := &StratumResponse{}
        //c.log.Info("get response from pool %s", result)
        err = json.Unmarshal([]byte(result), &response)
        if err!=nil{
            c.log.Error("failed to get response from pool: %s", err.Error())
            continue
        }
        //c.log.Info("get response : %v", response)
        c.msgChan <- response
    }
}

func (c *StratumClient) Send(request *StratumRequest) *StratumResponse{
    c.seq++
    request.ID = c.seq

    if(request.Params==nil){
        request.Params = make([]string, 0)
    }
    encoded,err:= json.Marshal(request)
    msg := string(encoded) + "\n"

    if err!=nil {return &StratumResponse{Error:err}}
    //c.log.Info("send msg to pool: %s", msg)
    _, err = c.socket.Write([]byte(msg));
    if err!=nil{
        c.log.Error("failed to send msg to pool: %s", err.Error())
        return &StratumResponse{Error:err}
    }

    return &StratumResponse{ID: 1, Error: nil}
}

func (c *StratumClient) Request(method string, params ...string) *StratumResponse{
    return c.Send(&StratumRequest{Method:method, Params:params});
}
