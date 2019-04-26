package libhub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func respToJSON(r *http.Response, v interface{}) error {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func show(r repo) string {
	return fmt.Sprintf("https://github.com/%s/%s/tree/%s", r.User, r.RepoName, r.Branch)
}
