package usecases

// FindExecutionsRequest represent FindExecutions request
type FindExecutionsRequest struct {
	JobName string `json:"job_name"`
	Last    int    `json:"last"`
}

// FindExecutionsResponse represents FindExecutions response
type FindExecutionsResponse struct {
	Executions []ExecutionDTO `json:"executions"`
}

// ExecutionDTO represents a execution returned by FindExecutionsResponse
type ExecutionDTO struct {
	JobName  string  `json:"job_name"`
	Date     string  `json:"date"`
	Status   string  `json:"status"`
	MemUsage float32 `json:"mem_usage"`
	CPUUsage float32 `json:"cpu_usage"`
	NetIn    float32 `json:"net_in"`
	NetOut   float32 `json:"net_out"`
}

// FindExecutions finds last N executions of all jobs or a specific job
func FindExecutions(request FindExecutionsRequest) FindExecutionsResponse {
	return reader.FindExecutionsResponse()
}
