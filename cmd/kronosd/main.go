package main

import (
	"log"
	"net/http"

	"github.com/taciomcosta/kronos/internal/config"
	"github.com/taciomcosta/kronos/internal/interfaces/os"
	"github.com/taciomcosta/kronos/internal/interfaces/rest"
	"github.com/taciomcosta/kronos/internal/interfaces/sqlite"
	"github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

func main() {
	err := config.EnableDefaultMode()
	if err != nil {
		log.Println("No configuration file detected, using default values")
	}

	writerReader := sqlite.NewCacheableWriterReader(config.GetString("db"))
	host := os.NewHost()
	notifierService := mocks.NewSpyNotifierService()
	usecases.New(writerReader, writerReader, host, notifierService)

	go usecases.ScheduleExistingJobs()
	log.Printf("%d job(s) loaded", usecases.FindJobs().Count)

	router := rest.NewRouter()
	log.Fatal(http.ListenAndServe(config.GetString("host"), router))
}
