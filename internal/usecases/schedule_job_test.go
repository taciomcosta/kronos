package usecases_test

import (
	"testing"
	"time"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

var testScheduleExistingJobs = []struct {
	expression string
	times      []time.Time
}{
	{
		expression: "* * * * *",
		times: []time.Time{
			time.Date(2021, 1, 1, 1, 1, 0, 0, time.UTC),
			time.Date(2021, 1, 1, 2, 1, 0, 0, time.UTC),
		},
	},
	{
		expression: "*/2 * * * *",
		times: []time.Time{
			time.Date(2021, 1, 1, 1, 2, 0, 0, time.UTC),
			time.Date(2021, 1, 1, 1, 4, 0, 0, time.UTC),
			time.Date(2021, 1, 1, 1, 6, 0, 0, time.UTC),
		},
	},
	{
		expression: "0 0 4 * 3",
		times: []time.Time{
			time.Date(2021, 1, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 1, 20, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 1, 27, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 2, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 2, 4, 0, 0, 0, 0, time.UTC),
		},
	},
	{
		expression: "every 5 minutes",
		times: []time.Time{
			time.Date(2021, 1, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 1, 13, 0, 5, 0, 0, time.UTC),
			time.Date(2021, 1, 13, 0, 10, 0, 0, time.UTC),
			time.Date(2021, 2, 13, 0, 15, 0, 0, time.UTC),
			time.Date(2021, 2, 13, 0, 20, 0, 0, time.UTC),
		},
	},
}

func TestScheduleExistingJobs(t *testing.T) {
	for _, tt := range testScheduleExistingJobs {
		for _, now := range tt.times {
			givenExpressionAssertJobIsCalledOnTime(t, tt.expression, now)
		}
	}
}

func givenExpressionAssertJobIsCalledOnTime(t *testing.T, expr string, now time.Time) {
	host := mocks.NewSpyHost()
	dependencies := uc.Dependencies{
		mocker.Stub().Writer().Build(),
		mocker.Stub().Reader().
			Set("FindJobs").Return(mocker.Data().Job().WithExpression(expr).Build()).
			Build(),
		host,
		mocks.SpyNotifierService(),
	}
	uc.New(dependencies)
	host.NotifyCurrentTimeIs(now)
	uc.ScheduleExistingJobs()
	if !host.DidJobRun() {
		t.Fatalf("job was not called in time %v", now)
	}
}

func TestScheduleDisabledJob(t *testing.T) {
	host := mocks.NewSpyHost()
	dependencies := uc.Dependencies{
		mocker.Stub().Writer().Build(),
		mocker.Stub().Reader().
			Set("FindJobs").Return(mocker.Data().Job().WithDisabled().Build()).
			Build(),
		host,
		mocks.SpyNotifierService(),
	}
	uc.New(dependencies)
	host.NotifyCurrentTimeIs(time.Date(2021, 2, 13, 0, 20, 0, 0, time.UTC))
	uc.ScheduleExistingJobs()
	if host.DidJobRun() {
		t.Fatalf("disabled job was called")
	}
}

func TestScheduleNotify(t *testing.T) {
	host := mocks.StubFailingHost()
	spyNotifierService := mocks.SpyNotifierService()
	dependencies := uc.Dependencies{
		mocker.Stub().Writer().Build(),
		mocker.Stub().Reader().Build(),
		host,
		spyNotifierService,
	}
	uc.New(dependencies)
	host.NotifyCurrentTimeIs(time.Date(2021, 2, 13, 0, 20, 0, 0, time.UTC))
	uc.ScheduleExistingJobs()
	if !spyNotifierService.SendWasCalled() {
		t.Fatalf("notifier was not called on job execution")
	}
}

func TestScheduleNotifyOnError(t *testing.T) {
	host := mocks.StubFailingHost()
	spyNotifierService := mocks.SpyNotifierService()
	writer := mocker.Stub().Writer().Build()
	reader := mocker.
		Stub().Reader().
		Set("FindAssignmentsByJob").Return(mocker.Data().Assignment().WithErrorOnly().Build()).
		Build()
	dependencies := uc.Dependencies{writer, reader, host, spyNotifierService}
	uc.New(dependencies)
	host.NotifyCurrentTimeIs(time.Date(2021, 2, 13, 0, 20, 0, 0, time.UTC))
	uc.ScheduleExistingJobs()
	if !spyNotifierService.SendWasCalled() {
		t.Fatalf("notifier was not called on job execution error")
	}
}

func TestScheduleDoesNotNotifyOnSucceed(t *testing.T) {
	host := mocks.NewSpyHost()
	spyNotifierService := mocks.SpyNotifierService()
	writer := mocker.Stub().Writer().Build()
	reader := mocker.
		Stub().Reader().
		Set("FindAssignmentsByJob").Return(mocker.Data().Assignment().WithErrorOnly().Build()).
		Build()
	dependencies := uc.Dependencies{writer, reader, host, spyNotifierService}
	uc.New(dependencies)
	host.NotifyCurrentTimeIs(time.Date(2021, 2, 13, 0, 20, 0, 0, time.UTC))
	uc.ScheduleExistingJobs()
	if spyNotifierService.SendWasCalled() {
		t.Fatalf("notifier was called on job execution error")
	}
}
