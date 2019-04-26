package libhub

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/juju/loggo"
)

type hubCup struct {
	Token      string
	AuthHeader *http.Header
	Client     *http.Client
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

func (hc *hubCup) run(method string, path string, res interface{}, data ...string) (err error) {
	var body string
	if len(data) > 0 {
		body = data[0]
	}
	resp, err := hc.Client.Do(&http.Request{
		Method: method,
		URL: &url.URL{
			Scheme: "https",
			Host:   "api.github.com",
			Path:   path,
		},
		Header: *hc.AuthHeader,
		Body:   ioutil.NopCloser(strings.NewReader(body)),
	})
	respToJSON(resp, res)
	return
}

func (hc *hubCup) getMe() (string, error) {
	var me struct {
		Login string `json:"login"`
	}
	err := hc.run("GET", "/user", &me)
	if err != nil {
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
	err = hc.run("GET", fmt.Sprintf("/repos/%s/%s", rep.User, rep.RepoName), &rif)
	logger.Debugf("%+v", rif)
	if err != nil {
		logger.Debugf("Get repo error!")
		return
	}
	return
}

func (hc *hubCup) getRefs(rep repo) (sha string, err error) {
	var shainfo struct {
		Object struct {
			Sha string `json: "object.sha"`
		} `json: "object"`
	}
	err = hc.run("GET",
		fmt.Sprintf("/repos/%s/%s/git/refs/heads/%s",
			rep.User, rep.RepoName, rep.Branch),
		&shainfo)
	sha = shainfo.Object.Sha
	return
}

func (hc *hubCup) setRefs(rep repo, sha string, force bool) (err error) {
	var v string
	err = hc.run("PATCH",
		fmt.Sprintf("/repos/%s/%s/git/refs/heads/%s",
			rep.User, rep.RepoName, rep.Branch), &v,
		fmt.Sprintf(`{"sha":"%s","force":%s}`, sha, strconv.FormatBool(force)),
	)
	return nil
}
