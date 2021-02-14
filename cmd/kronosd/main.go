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
	writerReader := sqlite.NewWriterReader("kronos.db")
	host := os.NewHost()
	usecases.New(writerReader, writerReader, host)
	go usecases.ScheduleExistingJobs()

	log.Printf("%d job(s) loaded", usecases.FindJobs().Count)

	router := rest.NewRouter()

	service := ":8080"
	log.Fatal(http.ListenAndServe(service, router))
}
