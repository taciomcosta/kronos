package spyhost

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

func newDefaultOutputs() map[string]interface{} {
	var outputs = make(map[string]interface{})
	d := &defaultSpyHost{}
	outputs["RunJob"] = d.RunJob()
	return outputs
}

// defaultSpyHost implements usecases.Host
type defaultSpyHost struct{}

// RunJob runs a job in host
func (s *defaultSpyHost) RunJob() []interface{} {
	execution := entities.Execution{
		JobName:  "spy-host",
		Date:     "date",
		Status:   "Succeeded",
		CPUTime:  1000,
		MemUsage: 1000,
	}
	return []interface{}{execution}
}
