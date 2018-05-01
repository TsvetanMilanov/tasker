package common

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/TsvetanMilanov/tasker/src/common/cconstants"
)

var (
	byteSliceKind = reflect.SliceOf(reflect.TypeOf(byte(1))).Kind()
)

// HTTPClient is client which helps with HTTP requests.
type HTTPClient struct {
}

// PostJSON makes POST request with JSON content type.
func (c *HTTPClient) PostJSON(url string, body interface{}, headers map[string]string, out interface{}) error {
	return c.doRequest(url, http.MethodPost, headers, body, out)
}

// GetJSON makes GET request with JSON content type.
func (c *HTTPClient) GetJSON(url string, headers map[string]string, out interface{}) error {
	return c.doRequest(url, http.MethodGet, headers, nil, out)
}

func (c *HTTPClient) doRequest(url, method string, headers map[string]string, body, out interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		var b []byte
		var err error
		// Marshal the body only if it is not []byte.
		if reflect.TypeOf(body).Kind() != byteSliceKind {
			b, err = json.Marshal(body)
			if err != nil {
				return err
			}
		}

		bodyReader = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", cconstants.ContentTypeApplicationJSON)
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Return []byte when the out is *[]byte.
	outElem := reflect.ValueOf(out).Elem()
	if outElem.Type().Kind() == byteSliceKind {
		outElem.SetBytes(resBodyBytes)
		return nil
	}

	return json.Unmarshal(resBodyBytes, out)
}
