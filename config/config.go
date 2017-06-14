package config

import "github.com/spf13/viper"

// Config ...
type Config struct {
	Filter *Filter
}

// NewConfig ...
func NewConfig() *Config {
	var f Filter
	err := viper.UnmarshalKey("filter", &f)
	if err != nil {
		panic(err)
	}
	return &Config{
		Filter: &f,
	}
}

type Filter struct {
	OutDir  []string
	OutFile []string
	InDir   []string
	InFile  []string
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}
