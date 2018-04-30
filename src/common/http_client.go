package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/TsvetanMilanov/tasker/src/common/cconstants"
)

// HTTPClient is client which helps with HTTP requests.
type HTTPClient struct {
}

// PostJSON makes POST request with JSON content type.
func (c *HTTPClient) PostJSON(url string, body interface{}, headers map[string]string, out interface{}) error {
	var b []byte
	var err error
	// Marshal the body only if it is not []byte.
	if reflect.TypeOf(body).Kind() != reflect.SliceOf(reflect.TypeOf(byte(1))).Kind() {
		b, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	bReader := bytes.NewBuffer(b)
	res, err := http.Post(url, cconstants.ContentTypeApplicationJSON, bReader)
	if err != nil {
		return err
	}

	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(resBodyBytes, out)
}
