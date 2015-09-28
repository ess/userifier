package main

import (
  "fmt"
  "flag"
  "github.com/ess/userifier/user"
  "github.com/ess/userifier/client"
  "github.com/ess/userifier/config"
)

func main() {
  var config_file = flag.String(
    "config",
    "/opt/userify/config.toml",
    "config file to use")
  flag.Parse()

  config := config.New(*config_file)
  client := client.New(config.ApiId, config.ApiKey)
  users := user.GetUsers(client.UserJSON())

  for _, user := range users {
    fmt.Println(user.String())
  }
}
