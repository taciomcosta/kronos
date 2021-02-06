package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/cmd/kronosd/handlers"
	"github.com/taciomcosta/kronos/internal/interfaces/sqlite"
	"github.com/taciomcosta/kronos/internal/usecases"
)

func main() {
	usecases.New(sqlite.NewRepository())
	log.Printf("%d job(s) loaded", usecases.CountJobs())

	router := httprouter.New()
	router.POST("/jobs", handlers.CreateJob)
	router.GET("/jobs", handlers.FindJobs)

	service := ":8080"
	log.Fatal(http.ListenAndServe(service, router))
}
