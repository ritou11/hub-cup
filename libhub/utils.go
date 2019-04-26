package libhub

import (
  "io/ioutil"
  "net/http"
  "encoding/json"
)

func respToJSON(r *http.Response, v interface{}) error {
  defer r.Body.Close()
  data, err := ioutil.ReadAll(r.Body)
  if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
