package client

import (
  "net/http"
  "crypto/tls"
  "io/ioutil"
  "encoding/base64"
)

type Client struct {
  ApiId string
  ApiKey string
}

func New(api_id string, api_key string) (client *Client) {
  return &Client{
    ApiId: api_id,
    ApiKey: api_key,
  }
}

func (client *Client) auth() string {
  return base64.StdEncoding.EncodeToString([]byte(client.ApiId + ":" + client.ApiKey))
}

func (client *Client) UserJSON() string {
  configure_url := "https://configure.userify.com/api/userify/configure"

  // Ignore bad or self-signed keys
  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }

  api := &http.Client{Transport: tr}

  req, _ := http.NewRequest("POST", configure_url, nil)

  // We want JSON, and we need to provide our auth creds
  req.Header.Add("Accept", "text/plain */json")
  req.Header.Add("Authorization", "Basic " + client.auth())

  resp, _ := api.Do(req)
  contents, _ := ioutil.ReadAll(resp.Body)

  return string(contents)
}

