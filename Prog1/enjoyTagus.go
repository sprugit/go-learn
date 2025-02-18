package main

import (
	"Prog1/internal/domain"
	"Prog1/internal/io"
	"log"
	"os"
)

func UpdateSchedule(
	skippersFilePath string,
	scheduleFilePath string,
	requestsFilePath string) (
	skippers *io.EntryFile,
	schedule *io.EntryFile,
	requests *io.EntryFile,
	err error) {

	skippers, err = io.ReadFile(skippersFilePath)
	if err != nil {
		return
	}
	schedule, err = io.ReadFile(scheduleFilePath)
	if err != nil {
		return
	}
	requests, err = io.ReadFile(requestsFilePath)
	if err != nil {
		return
	}

	log.Printf("Successfully loaded files %s %s %s into memory.",
		skippersFilePath, scheduleFilePath, requestsFilePath)

	err = domain.ProcessFiles(schedule, requests, skippers)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("Successfully processed files.")
	
	return
}

func IfErrrorExit(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {

	args := os.Args[1:]

	skippersFilePath := args[0]
	scheduleFilePath := args[1]
	requestsFilePath := args[2]

	var (
		_        *io.EntryFile /*Requests go unused in this scope*/
		schedule *io.EntryFile
		skippers *io.EntryFile
		err      error
	)

	skippers, schedule, _, err =
		UpdateSchedule(skippersFilePath, scheduleFilePath, requestsFilePath)
	IfErrrorExit(err)

	err = schedule.WriteFile()
	IfErrrorExit(err)

	err = skippers.WriteFile()
	IfErrrorExit(err)

	log.Printf("Successfully wrote files.")
	log.Printf("Goodbye :)")

	os.Exit(0)
}
