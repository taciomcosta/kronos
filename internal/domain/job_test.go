package domain

import (
	"testing"
)

func TestCreateJob(t *testing.T) {
	tests := []struct {
		request  CreateJobRequest
		response CreateJobResponse
	}{
		{
			request: CreateJobRequest{
				Name:    "list",
				Command: "ls",
				Tick:    "* * * * *",
			},
			response: CreateJobResponse{
				Msg:     "list created.",
				Success: true,
			},
		},
		{
			request: CreateJobRequest{
				Name:    "list",
				Command: "ls",
				Tick:    "1/2 * * * *",
			},
			response: CreateJobResponse{
				Msg:     "can't parse 1/2",
				Success: false,
			},
		},
	}

	Init(NewMockRepository())

	for _, tt := range tests {
		response := CreateJob(tt.request)
		if tt.response != response {
			t.Errorf("got %v, expected %v", response, tt.response)
		}
	}
}
