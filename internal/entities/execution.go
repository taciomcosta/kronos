package entities

// Execution represents execution of a job
type Execution struct {
	JobName  string
	Date     string
	Status   string
	MemUsage float32
	CPUUsage float32
	NetIn    float32
	NetOut   float32
}
