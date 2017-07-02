package main

import (
	"flag"
	"path/filepath"

	ag "github.com/sky0621/go-testcode-autogen"
	"github.com/sky0621/go-testcode-autogen/config"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("START")

	target := flag.String("target", "../_example/sampleproject", "Parse Target")
	cfg := flag.String("config", "../config/config.toml", "Config File")
	flag.Parse()

	config.ReadConfig(*cfg)
	ag.Cfg = config.NewConfig()

	err = filepath.Walk(*target, ag.Apply)
	if err != nil {
		logger.Error("ABEND", zap.Error(err))
	}
	logger.Info("END")
}
