package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// CreateAssignment handles assignements creation request
func CreateAssignment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var assignmentRequest uc.AssignNotifierToJobRequest
	err := ReadJSON(r.Body, &assignmentRequest)
	if err != nil {
		respondError(w, err)
	}
	response, err := uc.AssignNotifierToJob(assignmentRequest)
	respond(w, response, err)
}

// DeleteAssignment handles assignements creation request
func DeleteAssignment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var unassignRequest uc.UnassignNotifierFromJobRequest
	err := ReadJSON(r.Body, &unassignRequest)
	if err != nil {
		respondError(w, err)
	}
	response, err := uc.UnassignNotifierFromJob(unassignRequest)
	respond(w, response, err)
}
