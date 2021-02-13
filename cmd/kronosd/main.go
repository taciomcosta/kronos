package main

import (
	"log"
	"net/http"

	"github.com/taciomcosta/kronos/internal/interfaces/os"
	"github.com/taciomcosta/kronos/internal/interfaces/rest"
	"github.com/taciomcosta/kronos/internal/interfaces/sqlite"
	"github.com/taciomcosta/kronos/internal/usecases"
)

func main() {
	repository := sqlite.NewRepository("kronos.db")
	host := os.NewHost()
	usecases.New(repository, host)
	go usecases.ScheduleExistingJobs()

	log.Printf("%d job(s) loaded", usecases.CountJobs())

	router := rest.NewRouter()

	service := ":8080"
	log.Fatal(http.ListenAndServe(service, router))
}
