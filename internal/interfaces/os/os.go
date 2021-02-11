package os

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/taciomcosta/kronos/internal/entities"
)

type defaultOS struct{}

// NewHost creates a new host using default OS lib
func NewHost() entities.Host {
	return &defaultOS{}
}

// RunJob runs a job hosted by default OS lib
func (o *defaultOS) RunJob(job *entities.Job) error {
	fmt.Printf("Running job %s\n", job.Name)
	cmd := exec.Command(job.Command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// GetDettachedStream gets empty streams
func (o *defaultOS) GetDettachedStream() entities.Stream {
	return entities.Stream{}
}
