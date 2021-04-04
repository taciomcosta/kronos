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
