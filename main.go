package main

import (
  "fmt"
  "flag"
  "github.com/BurntSushi/toml"
  "os"
  "github.com/ess/userifier/user"
  "github.com/ess/userifier/client"
)

type userifierConfig struct {
  ApiId  string `toml:"api_id"`
  ApiKey string `toml:"api_key"`
}

func read_config(config_file string) userifierConfig {
  var config userifierConfig

  if _, err := toml.DecodeFile(config_file, &config); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  return config
}

func main() {
  var config_file = flag.String("config", "/opt/userify/config.toml", "config file to use")
  flag.Parse()

  config := read_config(*config_file)


  client := client.New(config.ApiId, config.ApiKey)

  users := user.GetUsers(client.UserJSON())

  for _, user := range users {
    fmt.Println(user.String())
  }
}
