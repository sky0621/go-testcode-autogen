package main

import (
	"flag"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	ag "github.com/sky0621/go-testcode-autogen"
	"github.com/sky0621/go-testcode-autogen/config"
)

func main() {
	target := flag.String("target", "../example/sampleproject", "Parse Target")
	cfg := flag.String("config", "../config/config.toml", "Config File")
	flag.Parse()

	config.ReadConfig(*cfg)

	err := filepath.Walk(*target, ag.Apply)
	if err != nil {
		logrus.Error(err)
	}
}
