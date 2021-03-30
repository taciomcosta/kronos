package usecases

// DescribeNotifierResponse represents response of DescribeNotifier
type DescribeNotifierResponse struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Metadata map[string]string `json:"metadata"`
}

// DescribeNotifier shows detailed information about a notifier
func DescribeNotifier(name string) (DescribeNotifierResponse, error) {
	return reader.DescribeNotifierResponse(name)
}
