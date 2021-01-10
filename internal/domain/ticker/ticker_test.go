package ticker

import (
	"errors"
	"testing"
	"time"
)

func TestNewTicker(t *testing.T) {
	tests := []struct {
		expression string
		err        error
	}{
		{expression: "* * * * *", err: nil},
		{expression: "*  *  *  *  *", err: nil},
		{expression: "n * * * *", err: errors.New("can't parse n")},
		{expression: "* n * * *", err: errors.New("can't parse n")},
		{expression: "* * n * *", err: errors.New("can't parse n")},
		{expression: "* * * n *", err: errors.New("can't parse n")},
		{expression: "* * * * n", err: errors.New("can't parse n")},
	}

	for _, tt := range tests {
		_, err := NewTicker(tt.expression)
		if tt.err != nil && tt.err.Error() != err.Error() {
			t.Errorf("got %v, expected %v", err, tt.err)
		}
		if err != nil && tt.err == nil {
			t.Errorf("got %v, expected %v", err, tt.err)
		}
	}
}

func TestIsTimeSet(t *testing.T) {
	tests := []struct {
		expression string
		times      []time.Time
		isSet      bool
	}{
		{
			expression: "* * * * *",
			times: []time.Time{
				time.Date(2021, 1, 1, 1, 1, 1, 0, time.UTC),
				time.Date(2021, 1, 1, 2, 1, 1, 0, time.UTC),
			},
			isSet: true,
		},
		{
			expression: "*/2 * * * *",
			times: []time.Time{
				time.Date(2021, 1, 1, 1, 1, 1, 0, time.UTC),
				time.Date(2021, 1, 1, 1, 3, 1, 0, time.UTC),
				time.Date(2021, 1, 1, 1, 7, 1, 0, time.UTC),
			},
			isSet: false,
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
			isSet: true,
		},
	}

	for _, tt := range tests {
		ticker, err := NewTicker(tt.expression)
		if err != nil {
			t.Fatal(err)
		}

		for _, time := range tt.times {
			if tt.isSet != ticker.IsTimeSet(time) {
				t.Errorf("%v should be set to %v", time, tt.isSet)
			}
		}
	}
}
