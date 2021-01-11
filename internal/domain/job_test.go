package domain

import (
	"testing"
)

func TestCreateJob(t *testing.T) {
	tests := []struct {
		request   CreateJobRequest
		expectErr bool
	}{
		{
			request: CreateJobRequest{
				Name:    "list",
				Command: "ls",
				Tick:    "* * * * *",
			},
			expectErr: false,
		},
		{
			request: CreateJobRequest{
				Name:    "list",
				Command: "ls",
				Tick:    "1/2 * * * *",
			},
			expectErr: true,
		},
	}

	Init(NewMockRepository())

	for _, tt := range tests {
		err := CreateJob(tt.request)
		if (tt.expectErr && err == nil) || (!tt.expectErr && err != nil) {
			t.Errorf("got %v", err)
		}
	}
}
