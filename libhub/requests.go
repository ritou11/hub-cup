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

func New(token string) (cup hubCup) {
  cup.Token = token
  cup.AuthHeader = req.Header{
    "Host": "https://api.github.com",
    "Authorization": "token " + token,
  }
  req.SetTimeout(5 * time.Second)
  return
}

func (hc hubCup) getMe() (string, error) {
  r, err := req.Do("get", "/user")
  if err != nil {
    logger.Debugf("Get user error!")
    return "", err
  }
  logger.Infof("%+v", r)
  return "1", err
}
