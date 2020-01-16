package main

import (
	"flag"
	"fmt"
	"github.com/tellor-io/TellorMiner/cli"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/ops"
	"github.com/tellor-io/TellorMiner/util"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"
)

var mainLog = util.NewLogger("main", "Main")


func main() {

	//create os kill sig listener
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	//see what args are passed in
	app := cli.App()
	app.Run(os.Args)

	//set things up
	config.ParseConfig(globalFlags.ConfigPath)
	util.ParseLoggingConfig(globalFlags.LoggingConfigPath)

	//global config
	cfg := config.GetConfig()

	runtime.GOMAXPROCS(cfg.NumProcessors)

	var ds *ops.DataServerOps
	//var miner *ops.MinerOps
	var miner *ops.MiningMgr


	//everything left over from the earlier parsing
	cmdArgs := flag.Args()

	if len(cmdArgs) == 0 {
		cli.Help(ctx, nil)
		return
	}
	cmdName := cmdArgs[0]
	cmd, ok := cli.Commands[cmdName]
	if !ok {
		fmt.Fprintf(os.Stderr, "'%s' is not a valid command. For a list, try\n\t%s help\n", cmdName, os.Args[0])
		return
	}
	err = cmd.Options.Parse(cmdArgs[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid arguments to %s: %s\n", cmdName, err.Error())
		fmt.Fprintf(os.Stderr, "Usage:\n")
		cmd.Options.PrintDefaults()
		return
	}
	if cmd.RequiresDB {
		ctx, err = AddDBToCtx(ctx, globalFlags.DataServer, globalFlags.Miner)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to setup DB: %s\n", err.Error())
			return
		}
	}
	err = cmd.Cmd(ctx, cmd.Options.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute %s: %s\n", cmdName, err.Error())
		return
	}
	return

	exitChannels := make([]*chan os.Signal, 0)

	if globalFlags.Transfer {
		ops.Transfer(globalFlags.ToAddress, globalFlags.Amount, ctx)
	} else if globalFlags.Deposit {
		ops.Deposit(ctx)
	} else if globalFlags.Approve {
		ops.Approve(globalFlags.ToAddress, globalFlags.Amount, ctx)
	} else if globalFlags.Dispute {
		ops.Dispute(globalFlags.RequestId, globalFlags.Timestamp, globalFlags.MinerIndex, ctx)
	} else if globalFlags.RequestStakingWithdraw {
		ops.RequestStakingWithdraw(ctx)
	} else if globalFlags.WithdrawStake {
		ops.WithdrawStake(ctx)
	} else if globalFlags.Vote {
		ops.Vote(globalFlags.DisputeId, globalFlags.SupportsDispute, ctx)
	} else {
		if globalFlags.DataServer {
			ch := make(chan os.Signal)
			exitChannels = append(exitChannels, &ch)
			ds, err = ops.CreateDataServerOps(ctx, ch)
			if err != nil {
				log.Fatal(err)
			}
		}

		if globalFlags.Miner {
			ch := make(chan os.Signal)
			exitChannels = append(exitChannels, &ch)
			miner, err = ops.CreateMiningManager(ctx, ch, ops.NewSubmitter())
			if err != nil {
				log.Fatal(err)
			}
		}

		if ds != nil {
			//start the data server
			ds.Start(ctx)
		}

		if miner != nil {
			//start the miner after at least one cycle from the data server, if it's running
			if ds != nil {
				<-ds.Ready()
			}

			miner.Start(ctx)
		}

		//now we wait for kill sig
		<-c
		//and then notify exit channels
		for _, ch := range exitChannels {
			*ch <- os.Interrupt
		}
		cnt := 0
		start := time.Now()
		for {
			cnt++
			dsStopped := false
			minerStopped := false

			if ds != nil {
				dsStopped = !ds.Running
			} else {
				dsStopped = true
			}

			if miner != nil {
				minerStopped = !miner.Running
			} else {
				minerStopped = true
			}

			if !dsStopped && !minerStopped && cnt > 60 {
				mainLog.Warn("Taking longer than expected to operations. Waited %v so far", time.Now().Sub(start))
			} else if dsStopped && minerStopped {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	mainLog.Info("Main shutdown complete")
}
