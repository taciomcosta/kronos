package rest

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// FindExecutions handles finding all executions request
func FindExecutions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := parseFindExecutionsRequest(r)
	executions := uc.FindExecutions(request)
	respond(w, executions, nil)
}

func parseFindExecutionsRequest(r *http.Request) uc.FindExecutionsRequest {
	query := r.URL.Query()
	jobName := query.Get("jobName")
	page, _ := strconv.Atoi(query.Get("page"))
	page--
	if page < 0 {
		page = 0
	}
	return uc.FindExecutionsRequest{JobName: jobName, Page: page}
}
