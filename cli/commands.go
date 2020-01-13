package cli

import (
	"context"
	"flag"
	"fmt"
	"github.com/tellor-io/TellorMiner/ops"
	"os"
)

type Command struct {
	Cmd     func(context.Context) error
	Options *flag.FlagSet
	RequiresDB bool
}

var Commands = map[string]Command{}

func init() {
	Commands["help"] = Command{
		Cmd:     Help,
		Options: flag.NewFlagSet("", flag.ContinueOnError),
	}
	Commands["deposit"] = depositCmd()
}

func depositCmd() Command {
	flags := new(flag.FlagSet)
	_ = flags.Uint64("bar", 10, "a foo bar")
	return Command{
		Options: flags,
		Cmd: ops.Deposit,
	}
}

func withdrawCmd() Command {
	flags := flag.NewFlagSet("withdraw", flag.ContinueOnError)
	return Command{
		Options: flags,
		Cmd: ops.WithdrawStake,
	}
}

//func HelpCmd() Command {
//
//}

func Help(ctx context.Context) error {
	fmt.Fprintf(os.Stderr, "TellorMiner commands\n")
	for k := range Commands {
		fmt.Fprintf(os.Stderr, "\t%s\n", k)
	}
	return nil
}
