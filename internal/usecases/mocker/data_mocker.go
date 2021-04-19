package mocker

// mocker.Data().Assignment().WithErrorOnly().Build()

// DataMocker represents root mock.Data()
type DataMocker struct{}

// Assignment represents root for creating mock Assignment
func (d *DataMocker) Assignment() *AssignmentBuilder {
	return newAssignmentBuilder()
}

// Job represents root for creating mock Job
func (d *DataMocker) Job() *JobBuilder {
	return newJobBuilder()
}

// Notifier represents root for creating mock Job
func (d *DataMocker) Notifier() *NotifierBuilder {
	return newNotifierBuilder()
}
