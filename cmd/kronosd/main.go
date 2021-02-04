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
	log.Printf("%d jobs loaded", domain.CountJobs())

	router := httprouter.New()
	router.POST("/jobs", handlers.CreateJob)
	router.GET("/jobs", handlers.FindJobs)

	service := ":8080"
	log.Fatal(http.ListenAndServe(service, router))
}
