package demo

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Run(ctx *cli.Context) error {
	logrus.Info("Demo run")
	return nil
}
