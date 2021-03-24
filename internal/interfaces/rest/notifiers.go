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
