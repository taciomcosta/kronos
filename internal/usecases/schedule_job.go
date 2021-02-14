package usecases

import (
	"time"
)

// ScheduleExistingJobs schedules jobs on startup
func ScheduleExistingJobs() {
	jobs = writer.FindJobs()
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
	for _, job := range jobs {
		if job.IsTimeSet(t) {
			host.RunJob(job)
		}
	}
}
