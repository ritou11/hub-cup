package libhub

import (
  "time"

  "github.com/juju/loggo"
  "github.com/imroc/req"
)

type hubCup struct {
  Token string
  AuthHeader req.Header
}

var logger = loggo.GetLogger("libhub")

func New(token string) hubCup {
  var cup hubCup
  hubCup.Token = token
  hub.AuthHeader = req.Header{
    "Host": "https://api.github.com",
    "Authorization": "token " + token,
  }
  req.SetTimeout(5 * time.Second)
}

func (hc hubCup) getMe() (string, error) {
  r, err := req.Do("get", "/user")
  if err != nil {
    logger.Debug("Get user error!")
    return "", err
  }
  log.Printf("%+v", r)
  return "1", err
}
