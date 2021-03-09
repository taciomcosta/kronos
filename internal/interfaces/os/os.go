package os

import (
	"log"
	"os"
	"os/exec"
	"syscall"
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
	processState, err := runJob(job)
	return newExecution(job, processState, err)
}

func runJob(job entities.Job) (*os.ProcessState, error) {
	log.Printf("Running job %s\n", job.Name)
	cmd := exec.Command(job.Command)
	err := cmd.Run()
	return cmd.ProcessState, err
}

func newExecution(job entities.Job, processState *os.ProcessState, err error) entities.Execution {
	if err != nil {
		return failedExecution(job, processState)
	}
	return succeededExecution(job, processState)
}

func failedExecution(job entities.Job, processState *os.ProcessState) entities.Execution {
	var execution entities.Execution
	execution.JobName = job.Name
	execution.Status = entities.FailedStatus
	execution.Date = time.Now().UTC().Format(time.RFC822)
	return execution
}

func succeededExecution(job entities.Job, processState *os.ProcessState) entities.Execution {
	var execution entities.Execution
	execution.JobName = job.Name
	execution.Status = entities.SucceededStatus
	execution.Date = time.Now().UTC().Format(time.RFC822)
	sysusage := processState.SysUsage()
	usage := sysusage.(*syscall.Rusage)
	execution.MemUsage = int(usage.Maxrss)
	execution.CPUTime = int(usage.Utime.Usec + usage.Stime.Usec)
	return execution
}
