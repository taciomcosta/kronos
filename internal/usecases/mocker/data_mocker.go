package mocker

// mocker.Data().Assignment().WithErrorOnly().Build()

// DataMocker represents root mock.Data()
type DataMocker struct{}

// Assignment represents root for creating mock Assignment
func (d *DataMocker) Assignment() *AssignmentBuilder {
	return newAssignmentBuilder()
}
