package main

import (
	"fmt"
	"os"

	"github.com/caiwp/armour-go/cmd"
	"github.com/caiwp/armour-go/modules/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	BuildTime = "unset"
	Commit    = "unset"
	Release   = "unset"
)

func main() {
	app := cli.NewApp()
	app.Name = "Armour"
	app.Usage = "TODO"
	app.Version = fmt.Sprintf("release:%s commit:%s build:%s", Release, Commit, BuildTime)
	app.Before = config.Init
	app.After = config.Close
	app.Commands = []cli.Command{
		cmd.Demo,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Errorf("Run app failed with %s: %v", os.Args, err)
	}
}
