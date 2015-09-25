package main

import (
  "fmt"
  "net/http"
  "crypto/tls"
  "io/ioutil"
  "encoding/base64"
  "flag"
  "github.com/BurntSushi/toml"
  "os"
  "github.com/ess/userifier/lib"
)

type userifierConfig struct {
  ApiId  string `toml:"api_id"`
  ApiKey string `toml:"api_key"`
}

func auth(id string, key string) string {
  return base64.StdEncoding.EncodeToString([]byte(id + ":" + key))
}

func read_config(config_file string) userifierConfig {
  var config userifierConfig

  if _, err := toml.DecodeFile(config_file, &config); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  return config
}

func userify_config(id string, key string) string {
  configure_url := "https://configure.userify.com/api/userify/configure"

  // Ignore bad or self-signed keys
  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }

  client := &http.Client{Transport: tr}

  req, _ := http.NewRequest("POST", configure_url, nil)

  // We want JSON, and we need to provide our auth creds
  req.Header.Add("Accept", "text/plain */json")
  req.Header.Add("Authorization", "Basic " + auth(id, key))

  resp, _ := client.Do(req)
  contents, _ := ioutil.ReadAll(resp.Body)

  return string(contents)
}

func get_users(json string) []User {
  fmt.Println(json)
  parsed, _ := gabs.ParseJSON([]byte(json))

  user_blobs, _ := parsed.Search("users").ChildrenMap()

  var users []User

  for user_name, _ := range user_blobs {
    user_data := parsed.Path("users." + user_name).Data().(map[string]interface{})

    users = append(users, User{user_name, user_data["name"].(string), user_data["ssh_public_key"].(string), user_data["preferred_shell"].(string), user_data["perm"].(string)})
  }

  return users
}

func main() {
  var config_file = flag.String("config", "/opt/userify/config.toml", "config file to use")
  flag.Parse()

  config := read_config(*config_file)


  api_id := config.ApiId
  api_key := config.ApiKey

  users := get_users(userify_config(api_id, api_key))

  for _, user := range users {
    fmt.Println(user)
  }
}
