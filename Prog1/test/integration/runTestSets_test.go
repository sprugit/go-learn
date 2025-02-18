package integration

import (
	"Prog1/internal/domain"
	"Prog1/internal/io"
	"Prog1/test/shared"
	"fmt"
	"testing"
)

func TestSet1(t *testing.T) {

	pathToRequests := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test1", "requests17h00.txt",
	})
	pathToSchedule := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test1", "schedule17h00.txt",
	})
	pathToSkippers := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test1", "skippers17h00.txt",
	})

	requestsFile, err := io.ReadFile(pathToRequests)
	if err != nil {
		t.Fatal(err)
	}
	scheduleFile, err := io.ReadFile(pathToSchedule)
	if err != nil {
		t.Fatal(err)
	}
	skippersFile, err := io.ReadFile(pathToSkippers)
	if err != nil {
		t.Fatal(err)
	}

	domain.ProcessFiles(scheduleFile, requestsFile, skippersFile)

	pathToExpectedSchedule := []string{
		"test", "integration", "data", "output", "test1", "schedule17h30.txt",
	}
	pathToExpectedSkippers := []string{
		"test", "integration", "data", "output", "test1", "skippers17h30.txt",
	}

	check1, err := shared.AssertMatchFileContents(pathToExpectedSchedule, scheduleFile.FileToString())
	if err != nil {
		t.Fatal(err)
	}
	if !check1 {
		fmt.Println(scheduleFile.FileToString())
		t.Fatal("Expected schedule doesn't match received schedule")
	}

	check2, err := shared.AssertMatchFileContents(pathToExpectedSkippers, skippersFile.FileToString())
	if err != nil {
		t.Fatal(err)
	}
	if !check2 {
		t.Fatal("Expected skippers doesn't match received skippers")
	}
}

func TestSet2(t *testing.T) {

	pathToRequests := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test2", "requests18h00.txt",
	})
	pathToSchedule := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test2", "schedule18h00.txt",
	})
	pathToSkippers := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test2", "skippers18h00.txt",
	})

	requestsFile, err := io.ReadFile(pathToRequests)
	if err != nil {
		t.Fatal(err)
	}
	scheduleFile, err := io.ReadFile(pathToSchedule)
	if err != nil {
		t.Fatal(err)
	}
	skippersFile, err := io.ReadFile(pathToSkippers)
	if err != nil {
		t.Fatal(err)
	}

	domain.ProcessFiles(scheduleFile, requestsFile, skippersFile)

	pathToExpectedSchedule := []string{
		"test", "integration", "data", "output", "test2", "schedule18h30.txt",
	}
	pathToExpectedSkippers := []string{
		"test", "integration", "data", "output", "test2", "skippers18h30.txt",
	}

	check1, err := shared.AssertMatchFileContents(pathToExpectedSchedule, scheduleFile.FileToString())
	if err != nil {
		t.Fatal(err)
	}
	if !check1 {
		t.Fatal("Expected schedule doesn't match received schedule")
	}

	check2, err := shared.AssertMatchFileContents(pathToExpectedSkippers, skippersFile.FileToString())
	if err != nil {
		t.Fatal(err)
	}
	if !check2 {
		t.Fatal("Expected skippers doesn't match received skippers")
	}
}

func TestSet3(t *testing.T) {

	pathToRequests := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test3", "requests10h00.txt",
	})
	pathToSchedule := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test3", "schedule10h00.txt",
	})
	pathToSkippers := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test3", "skippers10h00.txt",
	})

	requestsFile, err := io.ReadFile(pathToRequests)
	if err != nil {
		t.Fatal(err)
	}
	scheduleFile, err := io.ReadFile(pathToSchedule)
	if err != nil {
		t.Fatal(err)
	}
	skippersFile, err := io.ReadFile(pathToSkippers)
	if err != nil {
		t.Fatal(err)
	}

	domain.ProcessFiles(scheduleFile, requestsFile, skippersFile)

	pathToExpectedSchedule := []string{
		"test", "integration", "data", "output", "test3", "schedule10h30.txt",
	}
	pathToExpectedSkippers := []string{
		"test", "integration", "data", "output", "test3", "skippers10h30.txt",
	}

	check1, err := shared.AssertMatchFileContents(pathToExpectedSchedule, scheduleFile.FileToString())
	if err != nil {
		t.Fatal(err)
	}
	if !check1 {
		t.Fatal("Expected schedule doesn't match received schedule")
	}

	check2, err := shared.AssertMatchFileContents(pathToExpectedSkippers, skippersFile.FileToString())
	if err != nil {
		t.Fatal(err)
	}
	if !check2 {
		t.Fatal("Expected skippers doesn't match received skippers")
	}
}

func TestSet4(t *testing.T) {

	pathToRequests := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test4", "requests13h00.txt",
	})
	pathToSchedule := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test4", "schedule13h00.txt",
	})
	pathToSkippers := shared.FindTestResource([]string{
		"test", "integration", "data", "input", "test4", "skippers13h00.txt",
	})

	requestsFile, err := io.ReadFile(pathToRequests)
	if err != nil {
		t.Fatal(err)
	}
	scheduleFile, err := io.ReadFile(pathToSchedule)
	if err != nil {
		t.Fatal(err)
	}
	skippersFile, err := io.ReadFile(pathToSkippers)
	if err != nil {
		t.Fatal(err)
	}

	domain.ProcessFiles(scheduleFile, requestsFile, skippersFile)

	pathToExpectedSchedule := []string{
		"test", "integration", "data", "output", "test4", "schedule13h30.txt",
	}
	pathToExpectedSkippers := []string{
		"test", "integration", "data", "output", "test4", "skippers13h30.txt",
	}

	check1, err := shared.AssertMatchFileContents(pathToExpectedSchedule, scheduleFile.FileToString())
	if err != nil {
		t.Fatal(err)
	}
	if !check1 {
		t.Fatal("Expected schedule doesn't match received schedule")
	}

	check2, err := shared.AssertMatchFileContents(pathToExpectedSkippers, skippersFile.FileToString())
	if err != nil {
		t.Fatal(err)
	}
	if !check2 {
		t.Fatal("Expected skippers doesn't match received skippers")
	}
}
