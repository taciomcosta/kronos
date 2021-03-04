package usecases

// FindExecutionsRequest represent FindExecutions request
type FindExecutionsRequest struct {
	JobName string `json:"job_name"`
	Page    int    `json:"page"`
}

// FindExecutionsResponse represents FindExecutions response
type FindExecutionsResponse struct {
	Executions []ExecutionDTO `json:"executions"`
}

// ExecutionDTO represents a execution returned by FindExecutionsResponse
type ExecutionDTO struct {
	JobName  string `json:"job_name"`
	Date     string `json:"date"`
	Status   string `json:"status"`
	MemUsage int    `json:"mem_usage"`
	CPUTime  int    `json:"cpu_time"`
}

// FindExecutions finds last N executions of all jobs or a specific job
func FindExecutions(request FindExecutionsRequest) FindExecutionsResponse {
	return reader.FindExecutionsResponse(request)
}
