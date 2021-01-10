package main

import (
	"fmt"
	"log"
	"os"

	"github.com/taciomcosta/kronos/internal/domain"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s name executable-path tick", os.Args[0])
		return
	}

	name := os.Args[1]
	command := os.Args[2]
	tick := os.Args[3]

	job, err := domain.NewJob(name, command, tick)
	if err != nil {
		log.Fatalf("Error on creating job: %v", err)
	}

	runner := domain.NewJobRunner()
	runner.AddJob(job)

	runner.Start()
}
