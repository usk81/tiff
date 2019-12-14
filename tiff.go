package tiff

import (
	"time"
)

// Duration is used to save time diff
type Duration struct {
	From   time.Time
	To     time.Time
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
	Minus  bool
}

// New calcrates time diff and sets Duration
func New(from, to time.Time) (result *Duration) {
	a := from
	b := to
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	var minus bool
	if a.After(b) {
		a, b = b, a
		minus = true
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	d := Duration{
		From:   from,
		To:     to,
		Year:   int(y2 - y1),
		Month:  int(M2 - M1),
		Day:    int(d2 - d1),
		Hour:   int(h2 - h1),
		Minute: int(m2 - m1),
		Second: int(s2 - s1),
		Minus:  minus,
	}

	// Normalize negative values
	if d.Second < 0 {
		d.Second += 60
		d.Minute--
	}
	if d.Minute < 0 {
		d.Minute += 60
		d.Hour--
	}
	if d.Hour < 0 {
		d.Hour += 24
		d.Day--
	}
	if d.Day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		d.Day += 32 - t.Day()
		d.Month--
	}
	if d.Month < 0 {
		d.Month += 12
		d.Year--
	}
	return &d
}

// Term returns the difference between two dates
func (d *Duration) Term() (years, months, days, hours, minutes, seconds int) {
	return d.Year, d.Month, d.Day, d.Hour, d.Minute, d.Second
}

// Years returns the number of years elapsed from `from date` to `to date`
func (d *Duration) Years() int {
	if d.Minus {
		return 0 - d.Year
	}
	return d.Year
}

// Months returns the number of months elapsed from `from date` to `to date`
func (d *Duration) Months() int {
	r := yearToMonth(d.Year) + d.Month
	if d.Minus {
		return 0 - r
	}
	return r
}

// Days returns the number of days elapsed from `from date` to `to date`
func (d *Duration) Days() int {
	return d.Seconds() / (60 * 60 * 24)
}

// Hours returns the number of hours elapsed from `from date` to `to date`
func (d *Duration) Hours() int {
	return d.Seconds() / (60 * 60)
}

// Minutes returns the number of minutes elapsed from `from date` to `to date`
func (d *Duration) Minutes() int {
	return d.Seconds() / 60
}

// Seconds returns the number of seconds elapsed from `from date` to `to date`
func (d *Duration) Seconds() int {
	return int(d.To.Unix() - d.From.Unix())
}

func yearToMonth(y int) int {
	if y == 0 {
		return 0
	}
	return y * 12
}
