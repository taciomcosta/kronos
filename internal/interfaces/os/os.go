package os

import (
	"log"
	"os/exec"
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

type defaultOS struct{}

// NewHost creates a new host using default OS lib
func NewHost() entities.Host {
	return &defaultOS{}
}

// GetDettachedStream gets empty streams
func (o *defaultOS) GetDettachedStream() entities.Stream {
	return entities.Stream{}
}

// TickEverySecond creates a channel that emits current time on every second
func (o *defaultOS) TickEverySecond() <-chan time.Time {
	ticker := time.NewTicker(1 * time.Second)
	return ticker.C
}

// RunJob runs a job hosted by default OS lib
func (o *defaultOS) RunJob(job entities.Job) {
	go o.runInBackground(job)
}

func (o *defaultOS) runInBackground(job entities.Job) {
	log.Printf("Running job %s\n", job.Name)
	cmd := o.newCommandFromJob(job)
	err := cmd.Run()
	if err != nil {
		log.Printf("Error on running %s: %v\n", job.Name, err)
	}
}

func (o *defaultOS) newCommandFromJob(job entities.Job) *exec.Cmd {
	cmd := exec.Command(job.Command)
	// TODO: implement attach/detach
	//cmd.Stdin = os.Stdin
	//cmd.Stderr = os.Stderr
	//cmd.Stdout = os.Stdout
	return cmd
}
