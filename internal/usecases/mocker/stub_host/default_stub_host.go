package stubhost

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

func newDefaultOutputs() map[string]interface{} {
	var outputs = make(map[string]interface{})
	d := &defaultStubHost{}
	outputs["RunJob"] = d.RunJob()
	return outputs
}

// defaultStubHost implements usecases.Host
type defaultStubHost struct{}

// RunJob runs a job in host
func (s *defaultStubHost) RunJob() entities.Execution {
	return entities.Execution{
		JobName:  "failing-job",
		Date:     "date",
		Status:   entities.FailedStatus,
		CPUTime:  1000,
		MemUsage: 1000,
	}
}
