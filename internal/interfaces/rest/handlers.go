package rest

import (
	"encoding/json"
	"net/http"
)

func readJSONFromRequestBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}

func respond(w http.ResponseWriter, v interface{}, err error) {
	if err != nil {
		respondJSONBadRequest(w, err)
	} else {
		respondJSON(w, v)
	}
}

func respondJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-type", "application/json")
	bytes, err := json.Marshal(v)
	if err != nil {
		respondJSONBadRequest(w, err)
	} else {
		w.Write(bytes)
	}
}

// ErrorMessage represents a generic error message for http responses.
type ErrorMessage struct {
	Msg string `json:"msg"`
}

func respondJSONBadRequest(w http.ResponseWriter, err error) {
	errorMessage := ErrorMessage{Msg: err.Error()}
	w.WriteHeader(http.StatusBadRequest)
	respondJSON(w, errorMessage)
}
