package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// FindExecutions handles finding all executions request
func FindExecutions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := uc.FindExecutionsRequest{}
	executions := uc.FindExecutions(request)
	respond(w, executions, nil)
}
