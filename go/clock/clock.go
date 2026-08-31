package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	// Fix minutes
	for m < 0 {
		m += 60
		h--
	}

	for m >= 60 {
		m -= 60
		h++
	}

	// Fix hours
	for h < 0 {
		h += 24
	}

	for h >= 24 {
		h -= 24
	}

	return Clock{
		hour:   h,
		minute: m,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
