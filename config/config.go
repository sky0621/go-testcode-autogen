package config

import "github.com/spf13/viper"

// Config ...
type Config struct {
	Filter *FilterConfig
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Filter: NewFilterConfig(),
	}
}

type FilterConfig struct {
	OutDir  []string
	OutFile []string
	InDir   []string
	InFile  []string
}

func NewFilterConfig() *FilterConfig {
	return &FilterConfig{
		OutDir:  viper.GetStringSlice("filter.outdir"),
		OutFile: viper.GetStringSlice("filter.outfile"),
		InDir:   viper.GetStringSlice("filter.indir"),
		InFile:  viper.GetStringSlice("filter.infile"),
	}
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}
