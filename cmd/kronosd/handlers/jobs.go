package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/internal/domain"
)

// CreateJob handles jobs creation request.
func CreateJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var jobRequest domain.CreateJobRequest
	err := readJSONFromRequestBody(r, &jobRequest)
	if err != nil {
		respondJSONBadRequest(w, err)
	}
	response, err := domain.CreateJob(jobRequest)
	respond(w, response, err)
}

// FindJobs handles finding all jobs request.
func FindJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobs := domain.FindJobs()
	respond(w, jobs, nil)
}
