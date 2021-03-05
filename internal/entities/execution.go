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
	SucceededStatus = "Succeeded"
	FailedStatus    = "Failed"
)
