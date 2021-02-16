package usecases

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

// DeleteJobResponse represents response of delete job usecase
type DeleteJobResponse struct {
	Msg string `json:"msg"`
}

// DeleteJob deletes a job given its name
func DeleteJob(name string) (DeleteJobResponse, error) {
	job, err := reader.FindOneJob(name)
	if err != nil {
		return DeleteJobResponse{}, err
	}
	_ = writer.DeleteJob(name)
	unscheduleJob(job)
	return DeleteJobResponse{Msg: name + " deleted."}, nil
}

func unscheduleJob(job entities.Job) {
	var index int
	for i := range jobs {
		if jobs[i].Name == job.Name {
			index = i
			break
		}
	}
	jobs = append(jobs[:index], jobs[index+1:]...)
}
