package model

import (
	"Prog1/internal/model"
	"testing"
)

func TestIncrementSeconds(t *testing.T) {
	var date = model.DateTime{
		Year:    2024,
		Month:   1,
		Day:     1,
		Hours:   00,
		Minutes: 00,
		Seconds: 45,
	}

	date.IncrementSeconds(30)
	var result = date.String()
	var expected = "2024:01:01-00:01:15"
	if result != expected {
		t.Fatalf("Incorrect result: got %s expected %s", result, expected)
	}
}

func TestMinuteUpdate(t *testing.T) {
	var date = model.DateTime{
		Year:    2024,
		Month:   1,
		Day:     1,
		Hours:   00,
		Minutes: 00,
		Seconds: 45,
	}

	date.IncrementSeconds(90)
	var result = date.String()
	var expected = "2024:01:01-00:02:15"
	if result != expected {
		t.Fatalf("Incorrect result: got %s expected %s", result, expected)
	}
}

func TestHourUpdate(t *testing.T) {
	var date = model.DateTime{
		Year:    2024,
		Month:   1,
		Day:     1,
		Hours:   00,
		Minutes: 02,
		Seconds: 15,
	}

	date.IncrementMinutes(61)
	var result = date.String()
	var expected = "2024:01:01-01:03:15"
	if result != expected {
		t.Fatalf("Incorrect result: got %s expected %s", result, expected)
	}
}

func TestDayUpdate(t *testing.T) {
	var date = model.DateTime{
		Year:    2024,
		Month:   1,
		Day:     1,
		Hours:   01,
		Minutes: 03,
		Seconds: 15,
	}

	date.IncrementHours(25)
	var result = date.String()
	var expected = "2024:01:02-02:03:15"
	if result != expected {
		t.Fatalf("Incorrect result: got %s expected %s", result, expected)
	}
}

func TestConversion(t *testing.T) {
	var date = model.DateTime{
		Year:    2024,
		Month:   12,
		Day:     31,
		Hours:   23,
		Minutes: 59,
		Seconds: 59,
	}

	var result = date.ToSeconds()
	var expected = 541231235959
	if result != expected {
		t.Fatalf("Incorrect result: got %d expected %d", result, expected)
	}
}

func TestComparator(t *testing.T) {
	var date = model.DateTime{
		Year:    2024,
		Month:   1,
		Day:     1,
		Hours:   00,
		Minutes: 00,
		Seconds: 45,
	}
	var date2 = model.DateTime{
		Year:    2024,
		Month:   1,
		Day:     1,
		Hours:   00,
		Minutes: 00,
		Seconds: 45,
	}

	var expected = 0
	var result = date.CompareTo(&date2)
	if result != expected {
		t.Fatalf("Incorrect result: got %d expected %d", result, expected)
	}

	date.IncrementSeconds(60)
	result = date.CompareTo(&date2)
	expected = 1
	if result != expected {
		t.Fatalf("Incorrect result: got %d expected %d", result, expected)
	}

	date2.IncrementMinutes(2)
	result = date.CompareTo(&date2)
	expected = -1
	if result != expected {
		t.Fatalf("Incorrect result: got %d expected %d", result, expected)
	}

}
