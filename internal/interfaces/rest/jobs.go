package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// CreateJob handles jobs creation request
func CreateJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var jobRequest uc.CreateJobRequest
	err := ReadJSON(r.Body, &jobRequest)
	if err != nil {
		respondError(w, err)
	}
	response, err := uc.CreateJob(jobRequest)
	respond(w, response, err)
}

// FindJobs handles finding all jobs request
func FindJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobs := uc.FindJobs()
	respond(w, jobs, nil)
}

// DeleteJob handles deleting a job
func DeleteJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	response, err := uc.DeleteJob(name)
	respond(w, response, err)
}

// DescribeJob handles describe a job
func DescribeJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	response, err := uc.DescribeJob(name)
	respond(w, response, err)
}
