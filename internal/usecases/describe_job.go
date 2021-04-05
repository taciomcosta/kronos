package usecases

// DescribeJobResponse represents response of DescribeJob
type DescribeJobResponse struct {
	Name                string   `json:"name"`
	Command             string   `json:"command"`
	Tick                string   `json:"tick"`
	LastExecution       string   `json:"last_execution"`
	Status              bool     `json:"status"`
	ExecutionsSucceeded int      `json:"executions_succeeded"`
	ExecutionsFailed    int      `json:"executions_failed"`
	AverageCPU          int      `json:"average_cpu"`
	AverageMem          int      `json:"average_mem"`
	AssignedNotifiers   []string `json:"assigned_notifiers"`
}

//DescribeJob shows detailed information about a job
func DescribeJob(name string) (DescribeJobResponse, error) {
	return reader.DescribeJobResponse(name)
}
