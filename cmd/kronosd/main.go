package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/cmd/kronosd/handlers"
	"github.com/taciomcosta/kronos/internal/data"
	"github.com/taciomcosta/kronos/internal/domain"
)

func main() {
	data.New()
	domain.New(data.NewRepository())

	router := httprouter.New()
	router.POST("/jobs", handlers.CreateJob)

	service := ":8080"
	log.Fatal(http.ListenAndServe(service, router))
}
