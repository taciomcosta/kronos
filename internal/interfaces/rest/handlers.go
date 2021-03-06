package rest

import (
	"encoding/json"
	"io"
	"net/http"
)

// ReadJSON reads entire JSON from reader and decodes it
func ReadJSON(r io.Reader, v interface{}) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(v)
}

func respond(w http.ResponseWriter, v interface{}, err error) {
	if err != nil {
		respondError(w, err)
	} else {
		respondSuccess(w, v)
	}
}

func respondError(w http.ResponseWriter, err error) {
	errorMessage := ErrorMessage{Msg: err.Error()}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	respondSuccess(w, errorMessage)
}

func respondSuccess(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-type", "application/json")
	bytes, _ := json.Marshal(v)
	_, _ = w.Write(bytes)
}

// ErrorMessage represents a generic error message for http responses.
type ErrorMessage struct {
	Msg string `json:"msg"`
}
