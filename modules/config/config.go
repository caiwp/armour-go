package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/caiwp/armour-go/utils"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	cfg config.Config

	RootDir string

	App = struct {
		Env string `json:"env"`
	}{
		Env: "dev",
	}

	Log = struct {
		Path   string       `json:"path"`
		MaxAge int32        `json:"max_age"`
		Level  logrus.Level `json:"level"`
	}{
		Path:   "./logs",
		MaxAge: 7,
		Level:  logrus.DebugLevel,
	}
)

func Init(ctx *cli.Context) error {
	var err error
	if RootDir, err = utils.ExecPath(); err != nil {
		return err
	}
	cfgFile := ctx.String("conf")
	if cfgFile == "" {
		cfgFile = filepath.Join(RootDir, "conf", "app.json")
	}

	cfg = config.NewConfig()
	if err = cfg.Load(file.NewSource(file.WithPath(cfgFile))); err != nil {
		return err
	}

	if err = cfg.Get("app").Scan(&App); err != nil {
		return err
	}
	if err = cfg.Get("log").Scan(&Log); err != nil {
		return err
	}

	logrus.Info("Config init success.")

	if err = initLog(); err != nil {
		return err
	}

	// TODO init other

	return nil
}

func initLog() error {
	logrus.Infof("Init log with config: %+v", Log)

	var err error
	if !filepath.IsAbs(Log.Path) {
		if Log.Path, err = filepath.Abs(Log.Path); err != nil {
			return err
		}
	}
	if _, err := os.Stat(Log.Path); os.IsNotExist(err) {
		if err = os.MkdirAll(Log.Path, os.ModePerm); err != nil {
			return err
		}
	}
	writer, err := rotatelogs.New(
		filepath.Join(Log.Path, App.Env+".log.%Y%m%d"),
		rotatelogs.WithMaxAge(time.Duration(Log.MaxAge)*time.Hour*24),
		rotatelogs.WithLinkName(filepath.Join(Log.Path, App.Env+".log")),
	)
	if err != nil {
		return err
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(Log.Level)
	logrus.SetOutput(writer)

	logrus.Info("Init log success.")
	return nil
}

func Close(ctx *cli.Context) error {
	logrus.Info("Config close.")
	cfg.Close()
	return nil
}
