package main

import (
	"math/rand"
	"os"
	"sort"
	"syscall"
	"time"

	"github.com/nknorg/nkn/cmd/nknc/asset"
	"github.com/nknorg/nkn/cmd/nknc/common"
	"github.com/nknorg/nkn/cmd/nknc/debug"
	"github.com/nknorg/nkn/cmd/nknc/id"
	"github.com/nknorg/nkn/cmd/nknc/info"
	"github.com/nknorg/nkn/cmd/nknc/name"
	"github.com/nknorg/nkn/cmd/nknc/pruning"
	"github.com/nknorg/nkn/cmd/nknc/pubsub"
	"github.com/nknorg/nkn/cmd/nknc/service"
	"github.com/nknorg/nkn/cmd/nknc/wallet"
	"github.com/nknorg/nkn/util/log"
	"github.com/urfave/cli"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Panic: %+v", r)
		}
	}()

	app := cli.NewApp()
	app.Name = "nknc"
	app.Version = common.Version
	app.HelpName = "nknc"
	app.Usage = "command line tool for blockchain"
	app.UsageText = "nknc [global options] command [command options] [args]"
	app.HideHelp = false
	app.HideVersion = false
	//global options
	app.Flags = []cli.Flag{
		common.NewIpFlag(),
		common.NewPortFlag(),
	}
	//commands
	app.Commands = []cli.Command{
		*debug.NewCommand(),
		*info.NewCommand(),
		*wallet.NewCommand(),
		*asset.NewCommand(),
		*name.NewCommand(),
		*pubsub.NewCommand(),
		*id.NewCommand(),
		*pruning.NewCommand(),
		*service.NewCommand(),
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))

	if err := app.Run(os.Args); err != nil {
		switch err.(type) {
		case syscall.Errno:
			os.Exit(int(err.(syscall.Errno)))
		default:
			os.Exit(1)
		}
	}
}
