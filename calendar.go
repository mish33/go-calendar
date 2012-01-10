package calendar

import "time"

const (
	dow  = "\000\003\002\005\000\003\005\001\004\006\002\004"
	days = "\037\034\037\036\037\036\037\037\036\037\036\037"
)

// Weekday returns day of week number.
//
// Use Sakamoto's method (valid for dates in Oct 15, 1582 - Dec 31, 9999)
// (see http://en.wikipedia.org/wiki/Calculating_the_day_of_the_week )
//
// Args:
//	year (e.g., 1988)
//	month (1-12)
//	day of the month (1-31)
// Returns:
//	0 = Sunday, 1 = Monday, ...
//
func Weekday(year, month, day int) int {
	if month < 3 {
		year--
	}
	return (year + year/4 - year/100 + year/400 + int(dow[month-1]) + day) % 7
}

// IsLeap returns true if year is a leap year.
func IsLeap(year int64) bool {
	return year%400 == 0 || year%100 != 0 && year%4 == 0
}

// NextDay advance Time structure to the next day.
func NextDay(t *time.Time) {
	t.Day++
	if recalcDay(t) {
		t.Day = 1
		NextMonth(t)
	}
	if t.Weekday == 0 { // May be Weekday was not set, recalculate.
		recalcDow(t)
	} else { // Assume Weekday was set correctly.
		t.Weekday = (t.Weekday+1) % 7
	}
}

func recalcDow(t *time.Time) {
	if t.Year > 1582 && t.Year < 10000 {
		t.Weekday = Weekday(int(t.Year), t.Month, t.Day)
	}
}

// NextMonth advance Time structure one month ahead.
func NextMonth(t *time.Time) {
	t.Month++
	if t.Month > 12 {
		t.Month = 1
		t.Year++
	}
	recalcDay(t)
	recalcDow(t)
}

func recalcDay(t *time.Time) bool {
	ndays := int(days[t.Month-1])
	if IsLeap(t.Year) && t.Month == 2 {
		ndays++
	}
	if t.Day > ndays {
		t.Day = ndays
		return true
	}
	return false
}

// PrevMonth adjust Time structure one month behind.
func PrevMonth(t *time.Time) {
	t.Month--
	if t.Month == 0 {
		t.Month = 12
		t.Year--
	}
	recalcDay(t)
	recalcDow(t)
}

// PrevDay adjust Time structure one day behind.
func PrevDay(t *time.Time) {
	t.Day--
	if t.Day == 0 {
		t.Month--
		if t.Month == 0 {
			t.Month = 12
			t.Year--
		}
		t.Day = 31
		recalcDay(t)
	}
	if t.Weekday == 0 { // May be Weekday was not set, recalculate.
		recalcDow(t)
	} else { // Assume Weekday was set correctly.
		t.Weekday = (t.Weekday-1) % 7
	}
}
