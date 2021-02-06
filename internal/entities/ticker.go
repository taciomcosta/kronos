package entities

import (
	"strings"
	"time"
)

// Ticker represents a job tick
// Example: "* * * * *", "*/1 1-4 * * *"
type Ticker struct {
	minute     token
	hour       token
	dayOfMonth token
	month      token
	dayOfWeek  token
}

// NewTicker creates a new Ticker from a string
// like "* * * * *"
func NewTicker(expression string) (Ticker, error) {
	parts := strings.Fields(expression)
	minute, err := parseToken(parts[0], 0, 59)
	if err != nil {
		return Ticker{}, err
	}
	hour, err := parseToken(parts[1], 0, 23)
	if err != nil {
		return Ticker{}, err
	}
	dayOfMonth, err := parseToken(parts[2], 1, 31)
	if err != nil {
		return Ticker{}, err
	}
	month, err := parseToken(parts[3], 1, 12)
	if err != nil {
		return Ticker{}, err
	}
	dayOfWeek, err := parseToken(parts[4], 0, 6)
	if err != nil {
		return Ticker{}, err
	}
	return Ticker{minute, hour, dayOfMonth, month, dayOfWeek}, nil
}

// IsTimeSet tells if t is a time included in Ticker
func (ticker *Ticker) IsTimeSet(t time.Time) bool {
	return ticker.isMinuteSet(t) &&
		ticker.isHourSet(t) &&
		ticker.isMonthSet(t) &&
		ticker.isDayMonthOrDayWeekSet(t)
}

func (ticker *Ticker) isMinuteSet(t time.Time) bool {
	return ticker.minute.isSet(int(t.Minute()))
}

func (ticker *Ticker) isHourSet(t time.Time) bool {
	return ticker.hour.isSet(int(t.Hour()))
}

func (ticker *Ticker) isMonthSet(t time.Time) bool {
	return ticker.month.isSet(int(t.Month()))
}

func (ticker *Ticker) isDayMonthOrDayWeekSet(t time.Time) bool {
	return ticker.dayOfMonth.isSet(int(t.Day())) ||
		ticker.dayOfWeek.isSet(int(t.Weekday()))
}
