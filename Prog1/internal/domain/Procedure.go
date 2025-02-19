package domain

import (
	"Prog1/internal/io"
	"Prog1/internal/model"
	"fmt"
	"log"
)

func DropCompletedAssingments(header *model.Header, schedule []model.Entry) []model.Entry {
	log.Println("Procedure: Dropping completed trips")
	var newEntries = make([]model.Entry, 0)
	for index, entry := range schedule {
		// Explicit dereference
		var temp = entry.Timestamp.Copy() //entry.Timestamp is starting time
		temp.IncrementHours(entry.DurationInHours)
		if header.Timestamp.LessThan(temp) {
			newEntries = append(newEntries, entry)
		} else {
			log.Printf("Procedure: dropping entry #%d - completed", index+1)
		}
	}
	return newEntries
}

func ProcessFiles(scheduleFile *io.EntryFile,
	requestsFile *io.EntryFile,
	skippersFile *io.EntryFile) (err error) {

	if !(requestsFile.Header.Timestamp.Equals(scheduleFile.Header.Timestamp) &&
		scheduleFile.Header.Timestamp.Equals(skippersFile.Header.Timestamp) &&
		skippersFile.Header.Timestamp.Equals(requestsFile.Header.Timestamp)) {
		return fmt.Errorf("Input file error: time inconsistency between files.")
	}

	var (
		requests []model.Request
		skippers []model.Skipper
		schedule []model.Entry
	)

	// Update timestamps here since they will be used by the inner logic of assignment
	scheduleFile.Header.IncrementTimestamp()
	skippersFile.Header.IncrementTimestamp()

	requests, err = AsRequestList(requestsFile.Entries)
	if err != nil {
		return
	}

	log.Printf("Loaded %s onto memory as request objects", requestsFile.Filename)

	skippers, err = AsSkippersList(skippersFile.Entries)
	if err != nil {
		return
	}
	log.Printf("Loaded %s onto memory as skipper objects", skippersFile.Filename)
	schedule, err = AsEntryList(scheduleFile.Entries)
	if err != nil {
		return
	}
	log.Printf("Loaded %s onto memory as entry objects", scheduleFile.Filename)

	schedule = DropCompletedAssingments(scheduleFile.Header, schedule)

	err = ProcessRequests(&schedule, &requests, &skippers, scheduleFile.Header.Timestamp)
	if err != nil {
		return
	}

	skippersFile.UpdateListing(AsStringList(skippers))
	scheduleFile.UpdateListing(AsStringList(schedule))

	return
}
