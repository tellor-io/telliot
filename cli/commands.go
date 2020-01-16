package cli

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jawher/mow.cli"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/contracts1"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/ops"
	"github.com/tellor-io/TellorMiner/rpc"
	"github.com/tellor-io/TellorMiner/util"
	"log"
	"os"
)

var ctx context.Context

func ErrorHandler(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		cli.Exit(-1)
	}
}

func ErrorWrapper(f func() error) func() {
	return func() {
		ErrorHandler(f())
	}
}

func buildContext() error {
	cfg := config.GetConfig()
	//create an rpc client
	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}
	//create an instance of the tellor master contract for on-chain interactions
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	masterInstance, err := contracts.NewTellorMaster(contractAddress, client)
	transactorInstance, err := contracts1.NewTellorTransactor(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	ctx = context.WithValue(context.Background(), tellorCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, tellorCommon.MasterContractContextKey, masterInstance)
	ctx = context.WithValue(ctx, tellorCommon.TransactorContractContextKey, transactorInstance)

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return fmt.Errorf("problem getting private key: %s", err.Error())
	}
	ctx = context.WithValue(ctx, tellorCommon.PrivateKey, privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}

	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	ctx = context.WithValue(ctx, tellorCommon.PublicAddress, publicAddress)

	//Issue #55, halt if client is still syncing with Ethereum network
	s, err := client.IsSyncing(ctx)
	if err != nil {
		return fmt.Errorf("could not determine if Ethereum client is syncing: %v\n", err)
	}
	if s {
		return fmt.Errorf("ethereum node is still sycning with the network")
	}
	return nil
}

func AddDBToCtx() error {
	cfg := config.GetConfig()
	//create a db instance
	os.RemoveAll(cfg.DBFile)
	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		return err
	}


	var dataProxy db.DataServerProxy
	if true { //EVANTODO
		proxy, err := db.OpenLocalProxy(DB)
		if err != nil {
			return err
		}
		dataProxy = proxy
	} else {
		proxy, err := db.OpenRemoteDB(DB)
		if err != nil {
			return err
		}
		dataProxy = proxy
	}
	ctx = context.WithValue(ctx, tellorCommon.DataProxyKey, dataProxy)
	ctx = context.WithValue(ctx, tellorCommon.DBContextKey, DB)
	return nil
}


func App() *cli.Cli {


	app := cli.App("tellor", "Tellor Coin Miner")

	//app wide config options
	configPath := app.StringOpt("config", "config.json", "Path to the primary JSON config file")
	logPath := app.StringOpt("logConfig", "loggingConfig.json", "Path to a JSON logging config file")

	app.Before = func() {
		ErrorHandler(config.ParseConfig(*configPath))
		ErrorHandler(util.ParseLoggingConfig(*logPath))
	}

	ErrorHandler(buildContext())


	app.Command("deposit", "deposit TRB stake", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			ErrorHandler(ops.Deposit(ctx))
		}
	})

	//		cmd.Before = ErrorWrapper(AddDBToCtx)
	return app
}

