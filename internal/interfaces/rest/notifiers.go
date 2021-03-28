package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// CreateNotifier handles notifier creation request
func CreateNotifier(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var notifierRequest uc.CreateNotifierRequest
	err := ReadJSON(r.Body, &notifierRequest)
	if err != nil {
		respondError(w, err)
	}
	response, err := uc.CreateNotifier(notifierRequest)
	respond(w, response, err)
}

// FindNotifiers handles finding all notifiers request
func FindNotifiers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	notifiers := uc.FindNotifiers()
	respond(w, notifiers, nil)
}

// DeleteNotifier handles deleting a job
func DeleteNotifier(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	response, err := uc.DeleteNotifier(name)
	respond(w, response, err)
}
