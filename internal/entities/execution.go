package entities

// Execution represents execution of a job
type Execution struct {
	JobName  string
	Date     string
	Status   string
	MemUsage float64
	CPUUsage float64
	NetIn    int
	NetOut   int
}
