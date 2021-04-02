package entities

import "fmt"

// Execution represents execution of a job
type Execution struct {
	JobName  string
	Date     string
	Status   string
	MemUsage int
	CPUTime  int
}

// ErrorMessage formats error messsage of an execution
func (e Execution) ErrorMessage() string {
	return fmt.Sprintf("%s: %s execution failed", e.Date, e.JobName)
}

const (
	// SucceededStatus is set when execution exits with no errors
	SucceededStatus = "Succeeded"
	// FailedStatus is set when execution exits with an error
	FailedStatus = "Failed"
)
