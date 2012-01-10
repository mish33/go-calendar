package calendar

import (
	"testing"
	"time"
)

type dowTest struct {
	y, m, d, dow int
}

var dowTests = []dowTest {
	dowTest{2012, 1,  9, 1},
	dowTest{2000, 2, 29, 2},
}

func TestDow(test *testing.T) {
	for _, t := range dowTests {
		wd := Weekday(t.y, t.m, t.d)
		if wd != t.dow {
			test.Errorf("DoW(%d, %d, %d)=%d, want %d", t.y, t.m, t.d, wd, t.dow)
		}
	}
}


func TestDowConst(test *testing.T) {
	wiki := []byte{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	if dow != string(wiki) {
		test.Error("dow const != ", wiki)
	}
}


func TestDayConst(test *testing.T) {
	idays := []byte{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if days != string(idays) {
		test.Error("days const != ", idays)
	}
}


type dayTest struct {
	y1, m1, d1, y2, m2, d2 int
}

var nextdayTests = []dayTest {
	dayTest{2012,  1,  9, 2012, 1, 10},
	dayTest{2000,  2, 29, 2000, 3,  1},
	dayTest{2001, 12, 31, 2002, 1,  1},
	dayTest{2002,  2, 28, 2002, 3,  1},
	dayTest{2003,  6, 30, 2003, 7,  1},
	dayTest{2004,  2, 28, 2004, 2, 29},
	dayTest{1900,  2, 28, 1900, 3,  1},
}

func TestNextday(test *testing.T) {
	tm := new(time.Time)
	for _, t := range nextdayTests {
		tm.Year, tm.Month, tm.Day = int64(t.y1), t.m1, t.d1
		NextDay(tm)
		if tm.Year != int64(t.y2) || tm.Month != t.m2 || tm.Day != t.d2 {
			test.Errorf("NextDay(%d/%d/%d)=%d/%d/%d, want %d/%d/%d", t.y1, t.m1, t.d1, tm.Year, tm.Month, tm.Day, t.y2, t.m2, t.d2)
		}
	}
}

func TestPrevday(test *testing.T) {
	tm := new(time.Time)
	for _, t := range nextdayTests {
		tm.Year, tm.Month, tm.Day = int64(t.y2), t.m2, t.d2
		PrevDay(tm)
		if tm.Year != int64(t.y1) || tm.Month != t.m1 || tm.Day != t.d1 {
			test.Errorf("PrevDay(%d/%d/%d)=%d/%d/%d, want %d/%d/%d", t.y2, t.m2, t.d2, tm.Year, tm.Month, tm.Day, t.y1, t.m1, t.d1)
		}
	}
}


var nextmonthTests = []dayTest {
	dayTest{2012,  1,  9, 2012, 2,  9},
	dayTest{2000,  2, 29, 2000, 3, 29},
	dayTest{2001, 12, 31, 2002, 1, 31},
	dayTest{2003,  5, 31, 2003, 6, 30},
}

func TestNextmonth(test *testing.T) {
	tm := new(time.Time)
	for _, t := range nextmonthTests {
		tm.Year, tm.Month, tm.Day = int64(t.y1), t.m1, t.d1
		NextMonth(tm)
		if tm.Year != int64(t.y2) || tm.Month != t.m2 || tm.Day != t.d2 {
			test.Errorf("NextMonth(%d/%d/%d)=%d/%d/%d, want %d/%d/%d", t.y1, t.m1, t.d1, tm.Year, tm.Month, tm.Day, t.y2, t.m2, t.d2)
		}
	}
}

var prevmonthTests = []dayTest {
	dayTest{2012,  1,  9, 2012, 2,  9},
	dayTest{2001, 12, 31, 2002, 1, 31},
	dayTest{2003,  4, 30, 2003, 5, 31},
	dayTest{2000,  2, 29, 2000, 3, 31},
	dayTest{1900,  2, 28, 1900, 3, 31},
}

func TestPrevmonth(test *testing.T) {
	tm := new(time.Time)
	for _, t := range prevmonthTests {
		tm.Year, tm.Month, tm.Day = int64(t.y2), t.m2, t.d2
		PrevMonth(tm)
		if tm.Year != int64(t.y1) || tm.Month != t.m1 || tm.Day != t.d1 {
			test.Errorf("PrevMonth(%d/%d/%d)=%d/%d/%d, want %d/%d/%d", t.y2, t.m2, t.d2, tm.Year, tm.Month, tm.Day, t.y1, t.m1, t.d1)
		}
	}
}

