package domain

import (
	"testing"
)

func TestCreateJob(t *testing.T) {
	tests := []struct {
		request     CreateJobRequest
		response    CreateJobResponse
		expectedErr bool
	}{
		{
			request: CreateJobRequest{
				Name:    "list",
				Command: "ls",
				Tick:    "* * * * *",
			},
			response: CreateJobResponse{
				Msg: "list created.",
			},
			expectedErr: false,
		},
		{
			request: CreateJobRequest{
				Name:    "list",
				Command: "ls",
				Tick:    "1/2 * * * *",
			},
			response: CreateJobResponse{
				Msg: "can't parse 1/2",
			},
			expectedErr: true,
		},
	}

	New(NewMockRepository())

	for _, tt := range tests {
		response, err := CreateJob(tt.request)
		if tt.response != response {
			t.Errorf("got %v, expected %v", response, tt.response)
		}

		if tt.expectedErr && err == nil {
			t.Errorf("expected error, got nil")
		}
	}
}
