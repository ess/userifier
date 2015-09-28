package config

import (
  "os"
  "fmt"
  "github.com/BurntSushi/toml"
)

type Config struct {
  File string
  ApiId  string `toml:"api_id"`
  ApiKey string `toml:"api_key"`
}

func New(config_file string) (config *Config) {
  config = &Config{File: config_file}
  config.readConfigFile()
  return config
}

func (config *Config) readConfigFile() {
  if _, err := toml.DecodeFile(config.File, &config); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
