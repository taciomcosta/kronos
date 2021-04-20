package data

import "github.com/taciomcosta/kronos/internal/entities"

func newAssignmentBuilder() *AssignmentBuilder {
	assignment := entities.Assignment{
		Job:         "name",
		Notifier:    "myslack",
		OnErrorOnly: false,
	}
	builder := &AssignmentBuilder{assignment}
	return builder
}

// AssignmentBuilder is Tests Data Builder Pattern for entities.Assignment
type AssignmentBuilder struct {
	assignment entities.Assignment
}

// Build builds a new entities.Assignment
func (b *AssignmentBuilder) Build() entities.Assignment {
	return b.assignment
}

// WithErrorOnly sets Assignment.OnErrorOnly to true
func (b *AssignmentBuilder) WithErrorOnly() *AssignmentBuilder {
	b.assignment.OnErrorOnly = true
	return b
}
