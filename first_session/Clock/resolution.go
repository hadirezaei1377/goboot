package clock

import "fmt"

type Clock struct {
	Hours   int
	Minutes int
}

const (
	TotalMinutesPerHour = 60
	TotalMinutesPerDay  = 1440
)

func New(h, m int) Clock {
	totalMinutes := h*TotalMinutesPerHour + m

	totalMinutes %= TotalMinutesPerDay

	if totalMinutes < 0 {
		totalMinutes += TotalMinutesPerDay // for h < 0 in test cases
	}

	return Clock{
		Hours:   totalMinutes / TotalMinutesPerHour,
		Minutes: totalMinutes % TotalMinutesPerHour,
	}
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hours, c.Minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes) // left padded zero  ---> 23 = 23 but 3 = 03,  not more than 2 characteres

}
