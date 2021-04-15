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
	handleExecutionsNotifications(execution, job)
	_ = writer.CreateExecution(&execution)
}

func handleExecutionsNotifications(execution entities.Execution, job entities.Job) {
	for _, assignment := range reader.FindAssignmentsByJob(job.Name) {
		if assignment.ShouldNotifyExecution(execution) {
			notifier, _ := reader.FindOneNotifier(assignment.Notifier)
			_ = notifierService.Send(execution.ErrorMessage(), notifier)
		}
	}
}
