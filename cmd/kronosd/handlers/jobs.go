package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/internal/domain"
)

func CreateJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var jobRequest domain.CreateJobRequest
	err := readJsonFromRequestBody(r, &jobRequest)
	if err != nil {
		respondJsonBadRequest(w, err.Error())
	}
	response := domain.CreateJob(jobRequest)
	if response.Success {
		respondJson(w, response)
	} else {
		respondJsonBadRequest(w, response)
	}
}

func FindJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobs := domain.FindJobs()
	fmt.Println(jobs)
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
		respondJsonBadRequest(w, "Error marshaling JSON")
		return
	}
	w.Write(bytes)
}

func respondJsonBadRequest(w http.ResponseWriter, v interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	respondJson(w, v)
}
