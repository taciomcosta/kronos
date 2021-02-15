package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func get(urlPath string, response interface{}) error {
	httpResponse, err := http.Get(url + urlPath)
	if err != nil {
		return errors.New("Failed to obtain response from kronosd")
	}
	err = readJSON(httpResponse.Body, response)
	return err
}

func post(urlPath string, request interface{}, response interface{}) error {
	body, err := newBody(request)
	if err != nil {
		return errors.New("Failed to obtain response from kronosd")
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
