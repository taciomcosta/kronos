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
		fmt.Println(err)
	}
	err = domain.CreateJob(jobRequest)
	if err != nil {
		fmt.Println(err)
	}
}

func readJsonFromRequestBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}
