package usecases

import (
	"time"
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
		if job.IsTimeSet(t) {
			host.RunJob(job)
		}
	}
}
