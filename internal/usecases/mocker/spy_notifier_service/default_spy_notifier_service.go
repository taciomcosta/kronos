package spynotifierservice

func newDefaultOutputs() map[string]interface{} {
	var outputs = make(map[string]interface{})
	d := &defaultSpyNotifierService{}
	outputs["Send"] = d.Send()
	return outputs
}

// defaultSpyNotifierService implements usecases.Host
type defaultSpyNotifierService struct{}

// Send sends message to an external notifier service
func (s *defaultSpyNotifierService) Send() []interface{} {
	return []interface{}{nil}
}
