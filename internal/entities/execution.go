package entities

// Execution represents execution of a job
type Execution struct {
	JobName  string
	Date     string
	Status   string
	MemUsage int
	CPUTime  int
}

const (
	// SucceededStatus is set when execution exits with no errors
	SucceededStatus = "Succeeded"
	// FailedStatus is set when execution exits with an error
	FailedStatus = "Failed"
)
