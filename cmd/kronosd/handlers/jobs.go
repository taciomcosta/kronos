package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/internal/domain"
)

func CreateJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var jobRequest domain.CreateJobRequest
	err := readJsonFromRequestBody(r, &jobRequest)
	if err != nil {
		respondJsonBadRequest(w, err)
	}
	response, err := domain.CreateJob(jobRequest)
	respond(w, response, err)
}

func FindJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobs := domain.FindJobs()
	respond(w, jobs, nil)
}
