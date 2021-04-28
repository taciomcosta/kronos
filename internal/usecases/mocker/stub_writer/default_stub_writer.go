package stubwriter

func newDefaultOutputs() map[string]interface{} {
	var outputs = make(map[string]interface{})
	d := &defaultStubWriter{}
	outputs["CreateJob"] = d.CreateJob()
	outputs["DeleteJob"] = d.DeleteJob()
	outputs["CreateExecution"] = d.CreateExecution()
	outputs["CreateNotifier"] = d.CreateNotifier()
	outputs["DeleteNotifier"] = d.DeleteNotifier()
	outputs["CreateAssignment"] = d.CreateAssignment()
	outputs["DeleteAssignment"] = d.DeleteAssignment()
	return outputs
}

// defaultStubWriter implements entities.Reader for tests purposes
type defaultStubWriter struct{}

// CreateJob creates a job.
func (mr *defaultStubWriter) CreateJob() []interface{} {
	return []interface{}{nil}
}

// DeleteJob deletes a job
func (mr *defaultStubWriter) DeleteJob() []interface{} {
	return []interface{}{nil}
}

// CreateExecution stubs a new Execution creation
func (mr *defaultStubWriter) CreateExecution() []interface{} {
	return []interface{}{nil}
}

// UpdateJob updates a job
func (mr *defaultStubWriter) UpdateJob() {}

// CreateNotifier creates a notifier
func (mr *defaultStubWriter) CreateNotifier() []interface{} {
	return []interface{}{nil}
}

// DeleteNotifier deletes a notifier
func (mr *defaultStubWriter) DeleteNotifier() []interface{} {
	return []interface{}{nil}
}

// CreateAssignment creates an assignment
func (mr *defaultStubWriter) CreateAssignment() []interface{} {
	return []interface{}{nil}
}

// DeleteAssignment deletes an assignment
func (mr *defaultStubWriter) DeleteAssignment() []interface{} {
	return []interface{}{nil}
}
