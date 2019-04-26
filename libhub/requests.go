package libhub

import (
  "time"
  "errors"
  "net/http"
  "net/url"

  "github.com/juju/loggo"
)

type hubCup struct {
  Token string
  AuthHeader *http.Header
  Client *http.Client
}

var logger = loggo.GetLogger("libhub")

func New(token string) *hubCup {
  cup := new(hubCup)
  cup.Client = &http.Client{
    Timeout: 10 * time.Second,
  }
  cup.Token = token
  cup.AuthHeader = &http.Header{"Authorization": []string{"token " + token}}
  return cup
}

func (hc *hubCup) run(method string, path string, header *http.Header) (resp *http.Response, err error) {
  resp, err = hc.Client.Do(&http.Request{
    Method: method,
    URL: &url.URL{
      Scheme: "https",
      Host: "api.github.com",
      Path: path,
    },
    Header: *hc.AuthHeader,
  })
  return
}

func (hc *hubCup) GetMe() (string, error) {
  parsed := make(map[string]string)
  r, err := hc.run("GET", "/user", hc.AuthHeader)
  if err != nil {
    logger.Debugf("Get user error!")
    return "", err
  }
  respToJSON(r, &parsed)
  if len(parsed["login"]) == 0 {
    return "", errors.New("Cannot get user!")
  }
  return parsed["login"], err
}
