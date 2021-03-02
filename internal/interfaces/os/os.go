package os

import (
	"log"
	"os/exec"
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
	"github.com/taciomcosta/kronos/internal/usecases"
)

type defaultOS struct{}

// NewHost creates a new host using default OS lib
func NewHost() usecases.Host {
	return &defaultOS{}
}

// TickEverySecond creates a channel that emits current time on every second
func (o *defaultOS) TickEverySecond() <-chan time.Time {
	ticker := time.NewTicker(1 * time.Second)
	return ticker.C
}

// RunJob runs a job hosted by default OS lib
func (o *defaultOS) RunJob(job entities.Job) entities.Execution {
	err := o.runJob(job)
	return newExecution(job, err)
}

func (o *defaultOS) runJob(job entities.Job) error {
	log.Printf("Running job %s\n", job.Name)
	cmd := o.newCommandFromJob(job)
	return cmd.Run()
}

func (o *defaultOS) newCommandFromJob(job entities.Job) *exec.Cmd {
	cmd := exec.Command(job.Command)
	return cmd
}

func newExecution(job entities.Job, err error) entities.Execution {
	var execution entities.Execution
	execution.JobName = job.Name
	execution.Status = newExecutionStatus(err)
	execution.Date = time.Now().UTC().Format(time.RFC822)
	execution.MemUsage = 1.0
	execution.CPUUsage = 1.0
	execution.NetIn = 1
	execution.NetOut = 1
	return execution
}

func newExecutionStatus(err error) string {
	if err != nil {
		return "Failed"
	}
	return "Succeeded"
}
