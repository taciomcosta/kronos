package data

import "github.com/taciomcosta/kronos/internal/entities"

func newExecutionBuilder() *ExecutionBuilder {
	execution := entities.Execution{
		JobName:  "spy-host",
		Date:     "date",
		Status:   entities.SucceededStatus,
		CPUTime:  1000,
		MemUsage: 1000,
	}
	builder := &ExecutionBuilder{execution}
	return builder
}

// ExecutionBuilder is Tests Data Builder Pattern for entities.Execution
type ExecutionBuilder struct {
	execution entities.Execution
}

// Build builds a new entities.Execution
func (b *ExecutionBuilder) Build() entities.Execution {
	return b.execution
}

// WithFailure sets a failing execution
func (b *ExecutionBuilder) WithFailure() *ExecutionBuilder {
	b.execution.Status = entities.FailedStatus
	return b
}
