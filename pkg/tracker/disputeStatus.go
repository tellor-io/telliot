// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"fmt"
	"log"

	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tellor "github.com/tellor-io/TellorMiner/abi/contracts"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

const DisputeTrackerName = "DisputeTracker2"

type DisputeTracker struct {
}

func (b *DisputeTracker) String() string {
	return DisputeTrackerName
}

func (b *DisputeTracker) Exec(ctx context.Context) error {
	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	_conAddress := cfg.ContractAddress

	//convert to address
	contractAddress := common.HexToAddress(_conAddress)

	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		fmt.Println("instance Error, disputeStatus")
		return err
	}

	status, _, err := instance.GetStakerInfo(nil, fromAddress)

	if err != nil {
		fmt.Println("instance Error, disputeStatus")
		return err
	}
	enc := hexutil.EncodeBig(status)
	log.Printf("Staker Status: %v", enc)
	err = DB.Put(db.DisputeStatusKey, []byte(enc))
	if err != nil {
		fmt.Printf("Problem storing dispute info: %v\n", err)
		return err
	}
	// Issue #50, bail out of not able to mine
	// if status.Cmp(big.NewInt(1)) != 0 {
	// 	log.Fatalf("Miner is not able to mine with status %v. Stopping all mining immediately", status)
	// }

	//add all whitelisted miner addresses as well since they will be coming in
	//asking for dispute status
	for _, addr := range cfg.ServerWhitelist {
		address := common.HexToAddress(addr)
		//fmt.Println("Getting staker info for address", addr)
		status, _, err := instance.GetStakerInfo(nil, address)
		if err != nil {
			fmt.Printf("Could not get staker dispute status for miner address %s: %v\n", addr, err)
		}
		fmt.Printf("Whitelisted Miner %s Dispute Status: %v\n", addr, status)
		dbKey := fmt.Sprintf("%s-%s", strings.ToLower(address.Hex()), db.DisputeStatusKey)
		err = DB.Put(dbKey, []byte(hexutil.EncodeBig(status)))
		if err != nil {
			fmt.Printf("Problem storing staker dispute status: %v\n", err)
		}
	}
	//fmt.Println("Finished updated dispute status")
	return nil
}
