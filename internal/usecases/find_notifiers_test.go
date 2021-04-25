package usecases_test

import (
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
)

var testsFindNotifiersResponse = []struct {
	expect uc.FindNotifiersResponse
}{
	{
		expect: uc.FindNotifiersResponse{
			Count:     1,
			Notifiers: []uc.NotifierDTO{{Name: "myslack", Type: "slack"}},
		},
	},
}

func TestFindNotifiers(t *testing.T) {
	for _, tt := range testsFindNotifiersResponse {
		dependencies := uc.Dependencies{
			mocker.Dependencies().Writer().Build(),
			mocker.Dependencies().Reader().Build(),
			mocker.Dependencies().Host().Build(),
			mocker.Dependencies().NotifierService().Build(),
		}
		uc.New(dependencies)
		got := uc.FindNotifiers()
		assertFindNotifiersResponse(t, got, tt.expect)
	}
}

func assertFindNotifiersResponse(t *testing.T, got, want uc.FindNotifiersResponse) {
	if got.Count != want.Count {
		t.Fatalf("got count %v, want count %v", got.Count, want.Count)
	}
	for i := 0; i < len(want.Notifiers); i++ {
		if want.Notifiers[i] != got.Notifiers[i] {
			t.Fatalf("got %v, want %v", got.Notifiers[i], want.Notifiers[i])
		}
	}
}
