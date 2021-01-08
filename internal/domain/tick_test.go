package domain

import (
	"errors"
	"testing"
)

//type testTick struct {
//description string
//tick        string
//expected    []time.Time
//}

//var tests []testTick = []testTick{
//{
//description: "every minute",
//tick:        "* * * * *",
//expected: []time.Time{
//time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
//time.Date(2021, 1, 1, 0, 1, 0, 0, time.UTC),
//time.Date(2021, 1, 1, 0, 2, 0, 0, time.UTC),
//time.Date(2021, 1, 1, 0, 3, 0, 0, time.UTC),
//time.Date(2021, 1, 1, 0, 4, 0, 0, time.UTC),
//},
//},
////{
////description: "every day at midnight",
////tick:        "0 0 * * *",
////expected: []time.Time{
////time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
////time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
////time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
////time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC),
////time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
////},
////},
//}

//func TestParseTick(t *testing.T) {
//currentTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

//for _, tt := range tests {
//ticks := ParseTick(currentTime, tt.tick, 5)

//if len(tt.expected) != len(ticks) {
//t.Fatalf("different lengths: %s", tt.description)
//}

//for i := range tt.expected {
//if !tt.expected[i].Equal(ticks[i]) {
//t.Fatalf("values don't match: %s", tt.description)
//}
//}
//}
//}

func TestParseToken(t *testing.T) {
	tests := []struct {
		in    string
		inMin int
		inMax int
		out   []int
		err   error
	}{
		{"*", 0, 5, []int{0, 1, 2, 3, 4, 5}, nil},
		{"1", 0, 59, []int{1}, nil},
		{"60", 0, 59, []int{}, errors.New("60 out of range 0-59")},
		{"n", 0, 59, []int{}, errors.New("can't parse n")},
		{"1,2,3", 0, 59, []int{1, 2, 3}, nil},
		{"1,*,3", 1, 5, []int{1, 2, 3, 4, 5}, nil},
		{"1-4", 1, 5, []int{1, 2, 3, 4}, nil},
		{"1-", 1, 5, []int{}, errors.New("missing value in range 1-")},
		{"1-10", 1, 5, []int{}, errors.New("1-10 out of range 1-5")},
		{"1-5", 2, 5, []int{}, errors.New("1-5 out of range 2-5")},
		{"*/2", 1, 10, []int{1, 3, 5, 7, 9}, nil},
		{"1-5/2", 1, 10, []int{1, 3, 5}, nil},
		{"1-10/2,6", 1, 10, []int{1, 3, 5, 6, 7}, nil},
		{"1-20/2,6", 1, 10, []int{}, errors.New("1-20 out of range 1-10")},
		{"1-10/2,11", 1, 10, []int{}, errors.New("11 out of range 1-10")},
		{"10/2", 1, 10, []int{}, errors.New("can't parse 10/2")},
	}

	for _, tt := range tests {
		got, err := parseToken(tt.in, tt.inMin, tt.inMax)

		if tt.err != nil && tt.err.Error() != err.Error() {
			t.Fatalf("expected %v, got %v\n", tt.err, err)
		}

		for i := range tt.out {
			if got[i] != tt.out[i] {
				t.Fatalf("got %d, expected %d\n", got[i], tt.out[i])
			}
		}
	}
}
