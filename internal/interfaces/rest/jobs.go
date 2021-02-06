package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/internal/usecases"
)

// CreateJob handles jobs creation request.
func CreateJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var jobRequest usecases.CreateJobRequest
	err := readJSONFromRequestBody(r, &jobRequest)
	if err != nil {
		respondJSONBadRequest(w, err)
	}
	response, err := usecases.CreateJob(jobRequest)
	respond(w, response, err)
}

// FindJobs handles finding all jobs request.
func FindJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobs := usecases.FindJobs()
	respond(w, jobs, nil)
}
