package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/internal/domain"
)

func CreateJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var jobRequest domain.CreateJobRequest
	err := readJsonFromRequestBody(r, &jobRequest)
	if err != nil {
		respondError(w, err.Error())
	}
	response := domain.CreateJob(jobRequest)
	if response.Success {
		respondJson(w, response)
	} else {
		respondError(w, response.Msg)
	}
}

func FindJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobs := domain.FindJobs()
	respondJson(w, jobs)
}

func readJsonFromRequestBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}

func respondJson(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-type", "application/json")
	bytes, err := json.Marshal(v)
	if err != nil {
		respondError(w, "Error marshaling JSON")
		return
	}
	w.Write(bytes)
}

func respondError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	respondJson(w, msg)
}
