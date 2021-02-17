package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/taciomcosta/kronos/internal/config"
)

var errKronosResponse = errors.New("Failed to obtain response from kronosd")
var url = "http://localhost" + config.GetString("host")
var client kronosClient

type kronosClient struct{}

func (c kronosClient) get(urlPath string, response interface{}) error {
	httpResponse, err := http.Get(url + urlPath)
	if err != nil {
		return errKronosResponse
	}
	err = readJSON(httpResponse.Body, response)
	return err
}

func (c kronosClient) delete(urlPath string, response interface{}) error {
	request, err := http.NewRequest("DELETE", url+urlPath, nil)
	if err != nil {
		return errKronosResponse
	}
	httpResponse, err := http.DefaultClient.Do(request)
	if err != nil {
		return errKronosResponse
	}
	err = readJSON(httpResponse.Body, response)
	return err
}

func (c kronosClient) post(urlPath string, request interface{}, response interface{}) error {
	body, err := newBody(request)
	if err != nil {
		return errKronosResponse
	}
	httpResponse, err := http.Post(url+urlPath, "application/json", body)
	if err != nil {
		return err
	}
	err = readJSON(httpResponse.Body, response)
	return err
}

func newBody(v interface{}) (io.Reader, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(buf), nil
}

func readJSON(r io.Reader, v interface{}) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(v)
}
