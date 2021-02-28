package rest

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var maxLast = 10

// FindExecutions handles finding all executions request
func FindExecutions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := parseFindExecutionsRequest(r)
	executions := uc.FindExecutions(request)
	respond(w, executions, nil)
}

func parseFindExecutionsRequest(r *http.Request) uc.FindExecutionsRequest {
	query := r.URL.Query()
	jobName := query.Get("jobName")
	last, _ := strconv.Atoi(query.Get("last"))
	if last > maxLast || last == 0 {
		last = maxLast
	}
	return uc.FindExecutionsRequest{JobName: jobName, Last: last}
}
