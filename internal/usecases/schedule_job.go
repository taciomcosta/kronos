package usecases

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

// ScheduleExistingJobs schedules jobs on startup
func ScheduleExistingJobs() {
	tickForever()
}

func tickForever() {
	for now := range host.TickEverySecond() {
		if now.Second() == 0 {
			runAllJobs(now)
		}
	}
}

func runAllJobs(t time.Time) {
	for _, job := range reader.FindJobs() {
		if job.ShouldRun(t) {
			go runOneJob(job)
		}
	}
}

func runOneJob(job entities.Job) {
	execution := host.RunJob(job)
	if execution.Status == entities.FailedStatus {
		// TODO: use assigned notifier
		notifier := entities.Notifier{}
		_ = notifierService.Send(execution.ErrorMessage(), notifier)
	}
	_ = writer.CreateExecution(&execution)
}
