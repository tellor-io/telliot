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
	flags := flag.NewFlagSet("deposit", flag.ContinueOnError)
	return Command{
		Options: flags,
		Cmd: func(ctx context.Context) error {
			return ops.Deposit(ctx)
		},
	}
}


func Help(ctx context.Context) error {
	fmt.Fprintf(os.Stderr, "TellorMiner commands\n")
	for k := range Commands {
		fmt.Fprintf(os.Stderr, "\t%s\n", k)
	}
	return nil
}
