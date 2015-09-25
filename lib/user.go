package userifier

import (
  "github.com/jeffail/gabs"
)

type User struct {
  UserName string
  RealName string
  SshPublicKey string
  PreferredShell string
  Perm string
}

func (u User) String() string {
  return u.UserName + " (real name: " + u.RealName + ", preferred shell: " + u.PreferredShell + ")"
}

func GetUsers(json string) []User {
  parsed, _ := gabs.ParseJSON([]byte(json))

  user_blobs, _ := parsed.Search("users").ChildrenMap()

  var users []User

  for user_name, _ := range user_blobs {
    user_data := parsed.Path("users." + user_name).Data().(map[string]interface{})

    users = append(users, User{user_name, user_data["name"].(string), user_data["ssh_public_key"].(string), user_data["preferred_shell"].(string), user_data["perm"].(string)})
  }

  return users
}
