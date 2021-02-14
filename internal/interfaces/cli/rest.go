package cli

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func post(urlPath string, request interface{}, response interface{}) error {
	body, err := newBody(request)
	if err != nil {
		return err
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
