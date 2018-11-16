package cmd

import (
	"github.com/caiwp/armour-go/modules/demo"
	"github.com/urfave/cli"
)

var Demo = cli.Command{
	Name:   "demo",
	Action: demo.Run,
	//Before: config.Init,
	//After:  config.Close,
}
