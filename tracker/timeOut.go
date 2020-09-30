package tracker

import (
	"context"
	"fmt"
	"log"
	// "math/big"
	"strings"
	"encoding/hex"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	
)

//TimeOutTracker struct
type TimeOutTracker struct {
}

func (b *TimeOutTracker) String() string {
	return "TimeOutTracker"
}

//Exec - Places the Dispute Status in the database
func (b *TimeOutTracker) Exec(ctx context.Context) error {
	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress
	_conAddress := cfg.ContractAddress

	//convert to address
	contractAddress := common.HexToAddress(_conAddress)

	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		fmt.Println("instance Error, disputeStatus")
		return err
	}
	address := "000000000000000000000000" + _fromAddress[2:] 
	decoded, err := hex.DecodeString(address)
	if err != nil {
		log.Fatal(err)
	}
	hash := solsha3.SoliditySHA3(decoded)
	var data [32]byte
	copy(data[:], hash)
	status,err := instance.GetUintVar(nil,data)
	
	if err != nil {
		fmt.Println("instance Error, disputeStatus")
		return err
	}
	enc := hexutil.EncodeBig(status)
	err = DB.Put(db.TimeOutKey, []byte(enc))
	if err != nil {
		fmt.Printf("Problem storing dispute info: %v\n", err)
		return err
	}
	//Issue #50, bail out of not able to mine
	// if status.Cmp(big.NewInt(1)) != 0 {
	// 	log.Fatalf("Miner is not able to mine with status %v. Stopping all mining immediately", status)
	// }

	//add all whitelisted miner addresses as well since they will be coming in
	//asking for dispute status
	for _, addr := range cfg.ServerWhitelist {
		address := "000000000000000000000000" + addr[2:] 
		decoded, err := hex.DecodeString(address)
		if err != nil {
			log.Fatal(err)
		}
		hash := solsha3.SoliditySHA3(decoded)
		var data [32]byte
		copy(data[:], hash)
		status,err := instance.GetUintVar(nil,data)
		if err != nil {
			fmt.Printf("Could not get staker timeOut status for miner address %s: %v\n", addr, err)
		}
		if(status.Int64() > 0){
				fmt.Printf("Whitelisted Miner %s Last Time Mined: %v\n", addr, time.Unix(status.Int64(), 0))
		}
		from := common.HexToAddress(addr)
		dbKey := fmt.Sprintf("%s-%s", strings.ToLower(from.Hex()), db.TimeOutKey)
		err = DB.Put(dbKey, []byte(hexutil.EncodeBig(status)))
		if err != nil {
			fmt.Printf("Problem storing staker last time mined: %v\n", err)
		}
	}
	//fmt.Println("Finished updated dispute status")
	return nil
}
