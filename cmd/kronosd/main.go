package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/cmd/kronosd/handlers"
)

func main() {
	router := httprouter.New()
	router.POST("/jobs", handlers.CreateJob)

	service := ":8080"
	log.Fatal(http.ListenAndServe(service, router))
}
