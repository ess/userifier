package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	File   string
	ApiId  string `toml:"api_id"`
	ApiKey string `toml:"api_key"`
}

func New(config_file string) (config *Config, err error) {
	config = &Config{File: config_file}
	err = config.readConfigFile()
	return config, err
}

func (config *Config) readConfigFile() (err error) {
	if _, err := toml.DecodeFile(config.File, &config); err != nil {
		return err
	}
	return nil
}
