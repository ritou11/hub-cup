package libhub

import (
  "fmt"
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

func (hc *hubCup) run(method string, path string, header *http.Header, res interface{}) (err error) {
  resp, err := hc.Client.Do(&http.Request{
    Method: method,
    URL: &url.URL{
      Scheme: "https",
      Host: "api.github.com",
      Path: path,
    },
    Header: *hc.AuthHeader,
  })
  respToJSON(resp, res)
  return
}

func (hc *hubCup) getMe() (string, error) {
  var me struct {
    Login string `json:"login"`
  }
  err := hc.run("GET", "/user", hc.AuthHeader, &me)
  if err != nil{
    logger.Debugf("Get user error!")
    return "", err
  }
  if err != nil || len(me.Login) == 0 {
    logger.Debugf("Get user error!")
    return "", errors.New("no-user")
  }
  return me.Login, err
}

func (hc *hubCup) getRepo(rep repo) (rif repoInfo, err error) {
  err = hc.run("GET", fmt.Sprintf("/repos/%s/%s", rep.User, rep.RepoName), hc.AuthHeader, &rif)
  logger.Debugf("%+v", rif)
  if err != nil{
    logger.Debugf("Get repo error!")
    return
  }
  return
}
