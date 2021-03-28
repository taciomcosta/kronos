package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var jobsInMemory []entities.Job

// CreateJob creates a new job into database
func (c *CacheableWriterReader) CreateJob(job *entities.Job) error {
	err := c.parent.CreateJob(job)
	if err != nil {
		return err
	}
	jobsInMemory = append(jobsInMemory, *job)
	return nil
}

// UpdateJob updates a job
func (c *CacheableWriterReader) UpdateJob(job *entities.Job) {
	c.parent.UpdateJob(job)
	for _, j := range c.FindJobs() {
		if j.Name == job.Name {
			j.Status = job.Status
		}
	}
}

// DeleteJob deletes a job
func (c *CacheableWriterReader) DeleteJob(name string) error {
	err := c.parent.DeleteJob(name)
	if err != nil {
		return err
	}
	var index int
	for i := range c.FindJobs() {
		if jobsInMemory[i].Name == name {
			index = i
			break
		}
	}
	jobsInMemory = append(jobsInMemory[:index], jobsInMemory[index+1:]...)
	return nil
}

// FindJobs finds all jobs.
func (c *CacheableWriterReader) FindJobs() []entities.Job {
	if len(jobsInMemory) == 0 {
		jobsInMemory = c.parent.FindJobs()
	}
	return jobsInMemory
}

// FindJobsResponse returns all jobs in FindJobsResponse format
func (c *CacheableWriterReader) FindJobsResponse() uc.FindJobsResponse {
	return c.parent.FindJobsResponse()
}

// FindOneJob finds all jobs.
func (c *CacheableWriterReader) FindOneJob(name string) (entities.Job, error) {
	for _, job := range c.FindJobs() {
		if job.Name == name {
			return job, nil
		}

	}
	return entities.Job{}, errResourceNotFound
}

// DescribeJobResponse finds job in DeDescribeJobResponse format
func (c *CacheableWriterReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	return c.parent.DescribeJobResponse(name)
}
