package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var jobs []entities.Job

// CreateJob creates a new job into database
func (c *CacheableWriterReader) CreateJob(job *entities.Job) error {
	err := c.wr.CreateJob(job)
	if err != nil {
		return err
	}
	jobs = append(jobs, *job)
	return nil
}

// DeleteJob deletes a job
func (c *CacheableWriterReader) DeleteJob(name string) error {
	err := c.wr.DeleteJob(name)
	if err != nil {
		return err
	}
	var index int
	for i := range jobs {
		if jobs[i].Name == name {
			index = i
			break
		}
	}
	jobs = append(jobs[:index], jobs[index+1:]...)
	return nil
}

// FindJobs finds all jobs.
func (c *CacheableWriterReader) FindJobs() []entities.Job {
	if len(jobs) == 0 {
		jobs = c.wr.FindJobs()
	}
	return jobs
}

// FindJobsResponse returns all jobs in FindJobsResponse format
func (c *CacheableWriterReader) FindJobsResponse() uc.FindJobsResponse {
	return c.wr.FindJobsResponse()
}

// FindOneJob finds all jobs.
func (c *CacheableWriterReader) FindOneJob(name string) (entities.Job, error) {
	for _, job := range jobs {
		if job.Name == name {
			return job, nil
		}

	}
	return entities.Job{}, errResourceNotFound
}

// DescribeJobResponse finds job in DeDescribeJobResponse format
func (c *CacheableWriterReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	return c.wr.DescribeJobResponse(name)
}
