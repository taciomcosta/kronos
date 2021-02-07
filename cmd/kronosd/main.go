package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/internal/interfaces/rest"
	"github.com/taciomcosta/kronos/internal/interfaces/sqlite"
	"github.com/taciomcosta/kronos/internal/usecases"
)

func main() {
	repository := sqlite.NewRepository("kronos.db")
	usecases.New(repository)
	log.Printf("%d job(s) loaded", usecases.CountJobs())

	router := httprouter.New()
	router.POST("/jobs", rest.CreateJob)
	router.GET("/jobs", rest.FindJobs)

	service := ":8080"
	log.Fatal(http.ListenAndServe(service, router))
}
