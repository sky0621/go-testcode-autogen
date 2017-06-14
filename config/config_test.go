package config

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	err := ReadConfig("config.toml")
	if err != nil {
		panic(err)
	}
	cfg := NewConfig()
	fmt.Printf("%#v\n", cfg.Filter)
}
