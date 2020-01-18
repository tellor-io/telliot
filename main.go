package main

import (
	"fmt"
	"github.com/tellor-io/TellorMiner/cli"
	"os"
	"os/signal"
)

func main() {

	//create os kill sig listener
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	//see what args are passed in
	app := cli.App()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app.Run failed: %v\n", err)
	}
	//
	//var ds *ops.DataServerOps
	////var miner *ops.MinerOps
	//var miner *ops.MiningMgr
	//
	//
	//if cmd.RequiresDB {
	//	ctx, err = AddDBToCtx(ctx, globalFlags.DataServer, globalFlags.Miner)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "failed to setup DB: %s\n", err.Error())
	//		return
	//	}
	//}
	//err = cmd.Cmd(ctx, cmd.Options.Args())
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "failed to execute %s: %s\n", cmdName, err.Error())
	//	return
	//}
	//return
	//
	//exitChannels := make([]*chan os.Signal, 0)
	//
	//	if globalFlags.DataServer {
	//		ch := make(chan os.Signal)
	//		exitChannels = append(exitChannels, &ch)
	//		ds, err = ops.CreateDataServerOps(ctx, ch)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//
	//	if globalFlags.Miner {
	//		ch := make(chan os.Signal)
	//		exitChannels = append(exitChannels, &ch)
	//		miner, err = ops.CreateMiningManager(ctx, ch, ops.NewSubmitter())
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//
	//	if ds != nil {
	//		//start the data server
	//		ds.Start(ctx)
	//	}
	//
	//	if miner != nil {
	//		//start the miner after at least one cycle from the data server, if it's running
	//		if ds != nil {
	//			<-ds.Ready()
	//		}
	//
	//		miner.Start(ctx)
	//	}
	//
	//	//now we wait for kill sig
	//	<-c
	//	//and then notify exit channels
	//	for _, ch := range exitChannels {
	//		*ch <- os.Interrupt
	//	}
	//	cnt := 0
	//	start := time.Now()
	//	for {
	//		cnt++
	//		dsStopped := false
	//		minerStopped := false
	//
	//		if ds != nil {
	//			dsStopped = !ds.Running
	//		} else {
	//			dsStopped = true
	//		}
	//
	//		if miner != nil {
	//			minerStopped = !miner.Running
	//		} else {
	//			minerStopped = true
	//		}
	//
	//		if !dsStopped && !minerStopped && cnt > 60 {
	//			mainLog.Warn("Taking longer than expected to operations. Waited %v so far", time.Now().Sub(start))
	//		} else if dsStopped && minerStopped {
	//			break
	//		}
	//		time.Sleep(500 * time.Millisecond)
	//	}
	//}
	//mainLog.Info("Main shutdown complete")
}
