package io

import (
	"Prog1/internal/domain"
	"Prog1/internal/io"
	"Prog1/internal/model"
	"Prog1/test/shared"
	"strings"
	"testing"
)

func TestReadCorrectRequestsFile(t *testing.T) {

	pathToTestFile := shared.FindTestResource(
		[]string{
			"test", "unit", "data", "sample", "correct", "requests17h00.txt",
		})

	var expected = model.Request{
		Name: "Vladislav Maraev",
		Languages: &model.Language{
			Languages: []string{"english", "russian"},
		},
		Category:          2,
		Specialty:         "price",
		CruisePeriodHours: 2,
	}

	reqs, err := io.ReadFile(pathToTestFile)
	if err != nil {
		t.Fatal(err)
	}

	reqObjs, err := domain.AsRequestList(reqs.Entries)
	if err != nil {
		t.Fatal(err)
	}
	if !reqObjs[0].Equals(expected) {
		t.Fatal("Parse Fault")
	}
}

func TestReadIncorrectRequestsFile(t *testing.T) {

	pathToTestFile := shared.FindTestResource(
		[]string{
			"test", "unit", "data", "sample", "incorrect", "requests17h00.txt",
		})

	_, err := io.ReadFile(pathToTestFile)
	if err == nil {
		t.Fatal(err)
	}
}

func TestReadCorrectSkippersFile(t *testing.T) {

	pathToTestFile := shared.FindTestResource(
		[]string{
			"test", "unit", "data", "sample", "correct", "skippers17h00.txt",
		})

	date := strings.Split("(08:11:2022, 12:30)", ",")
	datetime, _ := model.DateTimeFromString(date[0], date[1])

	var expected = model.Skipper{
		Name: "Ana Amaral",
		Language: &model.Language{
			Languages: []string{"english", "french", "portuguese"},
		},
		Category:  1,
		Cost:      40,
		Specialty: "price",
		MaxHours:  20,
		AcumHours: 12,
		Datetime:  datetime,
	}

	skps, err := io.ReadFile(pathToTestFile)
	if err != nil {
		t.Fatal(err)
	}

	skpsObjs, err := domain.AsSkippersList(skps.Entries)
	if err != nil {
		t.Fatal(err)
	}
	if !expected.Equals(&skpsObjs[0]) {
		t.Fatal("Parse Fault")
	}
}

func TestReadIncorrectSkippersFile(t *testing.T) {

	pathToTestFile := shared.FindTestResource(
		[]string{
			"test", "unit", "data", "sample", "incorrect", "skippers17h00.txt",
		})

	_, err := io.ReadFile(pathToTestFile)
	if err == nil {
		t.Fatal(err)
	}
}

func TestReadCorrectScheduleFile(t *testing.T) {

	pathToTestFile := shared.FindTestResource(
		[]string{
			"test", "unit", "data", "sample", "correct", "schedule17h00.txt",
		})

	date := strings.Split("(07:11:2022, 10:00)", ",")
	datetime, _ := model.DateTimeFromString(date[0], date[1])

	var expected = model.Entry{
		Timestamp:       datetime,
		DurationInHours: 2,
		SkipperName:     "Jorge Costa",
		Cost:            160,
		RequestName:     "Ehrard Hinrichs",
	}

	schds, err := io.ReadFile(pathToTestFile)
	if err != nil {
		t.Fatal(err)
	}

	schdObjs, err := domain.AsEntryList(schds.Entries)
	if err != nil {
		t.Fatal(err)
	}
	if !schdObjs[0].Equals(expected) {
		t.Fatal("Parse Fault")
	}
}

func TestReadIncorrectScheduleFile(t *testing.T) {

	pathToTestFile := shared.FindTestResource(
		[]string{
			"test", "unit", "data", "sample", "incorrect", "schedule17h00.txt",
		})

	_, err := io.ReadFile(pathToTestFile)
	if err == nil {
		t.Fatal(err)
	}
}
