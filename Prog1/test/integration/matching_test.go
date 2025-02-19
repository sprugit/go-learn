package integration

import (
	"Prog1/internal/domain"
	"Prog1/internal/io"
	"Prog1/test/shared"
	"fmt"
	"testing"
)

func TestMatchSet1(t *testing.T) {

	pathToRequests := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test1", "requests17h00.txt",
	})
	pathToSchedule := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test1", "schedule17h00.txt",
	})
	pathToSkippers := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test1", "skippers17h00.txt",
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

	err = domain.ProcessFiles(scheduleFile, requestsFile, skippersFile)
	if err != nil {
		t.Fatal(err)
	}

	pathToExpectedSchedule := []string{
		"test", "integration", "data", "expected", "output", "test1", "schedule17h30.txt",
	}
	pathToExpectedSkippers := []string{
		"test", "integration", "data", "expected", "output", "test1", "skippers17h30.txt",
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

func TestMatchSet2(t *testing.T) {

	pathToRequests := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test2", "requests18h00.txt",
	})
	pathToSchedule := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test2", "schedule18h00.txt",
	})
	pathToSkippers := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test2", "skippers18h00.txt",
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

	err = domain.ProcessFiles(scheduleFile, requestsFile, skippersFile)
	if err == nil {
		t.Fatal(err)
	}
}

func TestMatchSet3(t *testing.T) {

	pathToRequests := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test3", "requests10h00.txt",
	})
	pathToSchedule := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test3", "schedule10h00.txt",
	})
	pathToSkippers := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test3", "skippers10h00.txt",
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

	err = domain.ProcessFiles(scheduleFile, requestsFile, skippersFile)
	if err != nil {
		t.Fatal(err)
	}

	pathToExpectedSchedule := []string{
		"test", "integration", "data", "expected", "output", "test3", "schedule10h30.txt",
	}
	pathToExpectedSkippers := []string{
		"test", "integration", "data", "expected", "output", "test3", "skippers10h30.txt",
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
		fmt.Println(skippersFile.FileToString())
		t.Fatal("Expected skippers doesn't match received skippers")
	}
}

func TestMatchSet4(t *testing.T) {

	pathToRequests := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test4", "requests13h00.txt",
	})
	pathToSchedule := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test4", "schedule13h00.txt",
	})
	pathToSkippers := shared.FindTestResource([]string{
		"test", "integration", "data", "expected", "input", "test4", "skippers13h00.txt",
	})

	_, err := io.ReadFile(pathToRequests)
	if err == nil {
		t.Fatal(err)
	}
	_, err = io.ReadFile(pathToSchedule)
	if err == nil {
		t.Fatal(err)
	}
	_, err = io.ReadFile(pathToSkippers)
	if err == nil {
		t.Fatal(err)
	}

}
