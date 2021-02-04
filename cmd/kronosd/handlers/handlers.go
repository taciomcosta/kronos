package handlers

import (
	"encoding/json"
	"net/http"
)

func readJsonFromRequestBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}

func respond(w http.ResponseWriter, v interface{}, err error) {
	if err != nil {
		respondJsonBadRequest(w, err)
	} else {
		respondJson(w, v)
	}
}

func respondJson(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-type", "application/json")
	bytes, err := json.Marshal(v)
	if err != nil {
		respondJsonBadRequest(w, err)
	} else {
		w.Write(bytes)
	}
}

type ErrorMessage struct {
	Msg string `json:"msg"`
}

func respondJsonBadRequest(w http.ResponseWriter, err error) {
	errorMessage := ErrorMessage{Msg: err.Error()}
	w.WriteHeader(http.StatusBadRequest)
	respondJson(w, errorMessage)
}
