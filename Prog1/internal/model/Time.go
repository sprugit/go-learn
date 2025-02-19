package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var monthDays = map[int]int{
	0:  31,
	1:  28,
	2:  31,
	3:  30,
	4:  31,
	5:  30,
	6:  31,
	7:  31,
	8:  30,
	9:  31,
	10: 30,
	11: 31,
}

type DateTime struct {
	Year, Month, Day, Hours, Minutes, Seconds int
}

func (t *DateTime) IncrementSeconds(seconds int) {
	var temp int
	t.Seconds += seconds
	if t.Seconds/60 > 0 {
		temp, t.Seconds = baseRemainder(t.Seconds, 60)
		t.Minutes += temp
	}
	for t.Minutes/60 > 0 {
		temp, t.Minutes = baseRemainder(t.Minutes, 60)
		t.Hours += temp
	}
	for t.Hours/24 > 0 {
		temp, t.Hours = baseRemainder(t.Hours, 24)
		t.Day += temp
	}
	for t.Day/monthDays[t.Month-1] > 0 {
		temp, t.Day = baseRemainder(t.Day, monthDays[t.Month-1])
		t.Month += temp
	}
	for t.Month/12 > 0 {
		temp, t.Month = baseRemainder(t.Month, 12)
		t.Year += 1
	}
}

func (t *DateTime) IncrementMinutes(minutes int) {
	t.IncrementSeconds(minutes * 60)
}

func (t *DateTime) IncrementHours(hours int) {
	t.IncrementMinutes(hours * 60)
}

func (t *DateTime) IncrementDays(days int) {
	t.IncrementHours(days * 24)
}

// Very simplified
func (t *DateTime) ToSeconds() int {
	return ((t.Year - 1970) * 10000000000) +
		(t.Month * 100000000) +
		(t.Day * 1000000) +
		(t.Hours * 10000) +
		(t.Minutes * 100) +
		t.Seconds
}

func (t *DateTime) DateString() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Day, t.Month, t.Year)
}

func (t *DateTime) TimeString() string {
	return fmt.Sprintf("%02d:%02d", t.Hours, t.Minutes)
}

func (t *DateTime) String() string {
	return fmt.Sprintf("%02d:%02d:%02d-%02d:%02d:%02d",
		t.Year, t.Month, t.Day, t.Hours, t.Minutes, t.Seconds)
}

func (t *DateTime) CompareTo(other *DateTime) int {
	var delta = t.ToSeconds() - other.ToSeconds()
	if delta > 0 {
		return 1
	}
	if delta < 0 {
		return -1
	}
	return 0
}

func (t *DateTime) Equals(other *DateTime) bool {
	var result = t.CompareTo(other)
	return result == 0
}

func (t *DateTime) BiggerThan(other *DateTime) bool {
	var result = t.CompareTo(other)
	return result > 0
}

func (t *DateTime) LessThan(other *DateTime) bool {
	var result = t.CompareTo(other)
	return result < 0
}

func baseRemainder(number int, base int) (quotient int, remainder int) {
	quotient = number / base
	remainder = number % base
	return
}

func DateTimeFromString(date string, time string) (*DateTime, error) {
	date = strings.TrimLeft(date, "(")  // might show up in certain cases
	time = strings.TrimRight(time, ")") // might show up in certain cases
	var dateArgs = strings.Split(date, ":")
	var timeArgs = strings.Split(time, ":")
	for index, val := range dateArgs {
		dateArgs[index] = strings.Trim(val, " ")
	}
	for index, val := range timeArgs {
		timeArgs[index] = strings.Trim(val, " ")
	}
	var dd, erry = strconv.Atoi(dateArgs[0])
	var mm, errm = strconv.Atoi(dateArgs[1])
	var yy, errd = strconv.Atoi(dateArgs[2])
	var hh, errh = strconv.Atoi(timeArgs[0])
	var mi, errmi = strconv.Atoi(timeArgs[1])
	if erry != nil || errm != nil || errd != nil || errh != nil || errmi != nil {
		return nil, errors.New("Invalid date format")
	}
	return &DateTime{
		Year:    yy,
		Month:   mm,
		Day:     dd,
		Hours:   hh,
		Minutes: mi,
		Seconds: 0,
	}, nil
}

func (t DateTime) IsBeforeClosingTime(seconds int) bool {
	t.IncrementSeconds(seconds)
	return (t.Hours < 20) || (t.Hours == 20 && t.Minutes == 00)
}

func MinuteAsSeconds(minutes int) int {
	return minutes * 60
}

func HourAsSeconds(hours int) int {
	return hours * MinuteAsSeconds(60)
}

func (t *DateTime) RoundUp() {
	if t.Minutes > 0 && t.Minutes < 30 {
		t.Minutes = 30
	} else {
		if t.Minutes > 30 {
			t.Minutes = 00
			t.IncrementHours(1)
		}
	}
}

func (t *DateTime) Copy() *DateTime {
	return &DateTime{
		Year:    t.Year,
		Month:   t.Month,
		Day:     t.Day,
		Hours:   t.Hours,
		Minutes: t.Minutes,
		Seconds: t.Seconds,
	}
}

func (t *DateTime) SetToNextDay() {
	t.IncrementDays(1)
	t.Hours = 8
	t.Minutes = 0
	t.Seconds = 0
}
