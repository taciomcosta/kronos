package main

import (
	"log"
	"net/http"

	"github.com/taciomcosta/kronos/internal/config"
	"github.com/taciomcosta/kronos/internal/interfaces/os"
	"github.com/taciomcosta/kronos/internal/interfaces/rest"
	"github.com/taciomcosta/kronos/internal/interfaces/services"
	"github.com/taciomcosta/kronos/internal/interfaces/sqlite"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

func main() {
	initializeConfig()
	uc.New(buildDependencies())
	initializeScheduler()
	initializeAPI()
}

func initializeConfig() {
	err := config.EnableDefaultMode()
	if err != nil {
		log.Println("No configuration file detected, using default values")
	}
}

func buildDependencies() uc.Dependencies {
	writerReader := sqlite.NewCacheableWriterReader(config.GetString("db"))
	host := os.NewHost()
	notifierService := services.NewNotifierService()
	return uc.Dependencies{
		Writer:          writerReader,
		Reader:          writerReader,
		Host:            host,
		NotifierService: notifierService,
	}
}

func initializeScheduler() {
	go uc.ScheduleExistingJobs()
	log.Printf("%d job(s) loaded", uc.FindJobs().Count)
}

func initializeAPI() {
	router := rest.NewRouter()
	log.Fatal(http.ListenAndServe(config.GetString("host"), router))
}
