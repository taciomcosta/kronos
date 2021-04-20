package mocker

import (
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

func newFindJobsResponseBuilder() *FindJobsResponseBuilder {
	response := uc.FindJobsResponse{
		Jobs: []uc.JobDTO{
			{
				Name:    "name",
				Command: "cmd",
				Tick:    "* * * * *",
				Status:  true,
			},
		},
		Count: 1,
	}
	builder := &FindJobsResponseBuilder{response}
	return builder
}

// FindJobsResponseBuilder is Tests Data Builder Pattern for uc.FindJobsResponse
type FindJobsResponseBuilder struct {
	response uc.FindJobsResponse
}

// Build builds a new entities.Assignment
func (b *FindJobsResponseBuilder) Build() uc.FindJobsResponse {
	return b.response
}

// WithExpression sets job expression
func (b *FindJobsResponseBuilder) WithExpression(expression string) *FindJobsResponseBuilder {
	b.response = uc.FindJobsResponse{
		Jobs: []uc.JobDTO{
			{
				Name:    "name",
				Command: "cmd",
				Tick:    expression,
				Status:  true,
			},
		},
		Count: 1,
	}
	return b
}
