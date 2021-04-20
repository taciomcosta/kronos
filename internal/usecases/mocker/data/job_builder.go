package data

import "github.com/taciomcosta/kronos/internal/entities"

func newJobBuilder() *JobBuilder {
	job, _ := entities.NewJob("name", "cmd", "* * * * *", true)
	builder := &JobBuilder{job}
	return builder
}

// JobBuilder is Tests Data Builder Pattern for entities.Assignment
type JobBuilder struct {
	job entities.Job
}

// Build builds a new entities.Assignment
func (b *JobBuilder) Build() entities.Job {
	return b.job
}

// WithExpression sets job expression
func (b *JobBuilder) WithExpression(expression string) *JobBuilder {
	job, _ := entities.NewJob(b.job.Name, b.job.Command, expression, b.job.Status)
	b.job = job
	return b
}

// WithDisabled sets job expression
func (b *JobBuilder) WithDisabled() *JobBuilder {
	b.job.Status = false
	return b
}
