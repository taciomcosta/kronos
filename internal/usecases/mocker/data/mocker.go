package data

// Mocker represents root mock.Data()
type Mocker struct{}

// Assignment represents root for creating mock Assignment
func (d *Mocker) Assignment() *AssignmentBuilder {
	return newAssignmentBuilder()
}

// Job represents root for creating mock Job
func (d *Mocker) Job() *JobBuilder {
	return newJobBuilder()
}

// Notifier represents root for creating mock Job
func (d *Mocker) Notifier() *NotifierBuilder {
	return newNotifierBuilder()
}

// FindJobsResponse represents root for creating mock FindJobsResponse
func (d *Mocker) FindJobsResponse() *FindJobsResponseBuilder {
	return newFindJobsResponseBuilder()
}

// Execution represents root for creating mock Execution
func (d *Mocker) Execution() *ExecutionBuilder {
	return newExecutionBuilder()
}
