package domain

import (
	"Prog1/internal/io"
	"Prog1/internal/model"
	"fmt"
	"log"
	"slices"
)

const CruiseLength = 30

/*
Misread the requirements: Unused code
*/ /*
func dropCompletedAssingments(schedule *io.EntryFile[*model.Entry]) {
	var newEntries = make([]model.Entry, 10)
	for _, entry := range schedule.Entries {
		// Explicit dereference
		if entry.Timestamp.BiggerThan(schedule.Header.Timestamp) {
			newEntries = append(newEntries, *entry)
		}
	}
	schedule.Entries = &newEntries
}*/

func assignSkipper(schedule *[]model.Entry,
	request *model.Request,
	skipper *model.Skipper,
	timestamp *model.DateTime) {
	var RequestName, SkipperName string
	var SkipperCost int
	RequestName = request.Name
	if skipper != nil {
		SkipperName = skipper.Name
		SkipperCost = skipper.Cost
	}
	var newEntry = model.Entry{
		Timestamp:   timestamp,
		Duration:    request.CruisePeriod,
		SkipperName: SkipperName,
		Cost:        SkipperCost,
		RequestName: RequestName,
	}
	*schedule = append(*schedule, newEntry)
}

func ProcessRequests(scheduleFile *io.EntryFile,
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

	for indexReq, request := range requests {
		log.Printf("Processing request #%d for client %s", indexReq, request.Name)
		slices.SortFunc(skippers, SortSkippers)
		var foundSkipper = false
		var timestamp model.DateTime
		timestamp = *scheduleFile.Header.Timestamp
		for i := 0; i < len(skippers) && !foundSkipper; i++ {
			var skipper = skippers[i]
			log.Printf("Attempting to match client %s with Skipper %s", request.Name, skipper.Name)
			var cLang, cSpec, cCat, cPastTime, cAccum bool
			cLang = request.Languages.HasMatch(*skipper.Language)
			cSpec = request.Specialty == skipper.Specialty
			cCat = request.Category == skipper.Category
			if cLang && cSpec && cCat {
				log.Printf("%s matches %s's needs in terms of language, specialty and category", skipper.Name, request.Name)
				var RequestCruiseLengthInMinutes = CruiseLength * request.CruisePeriod
				cPastTime = skipper.Datetime.
					IsBeforeClosingTime(
						model.HourAsSeconds(RequestCruiseLengthInMinutes))
				cAccum = (skipper.AcumHours + request.CruisePeriod) < skipper.MaxHours
				if cPastTime && cAccum {
					log.Printf("Assigning %s to %s", skipper.Name, request.Name)
					foundSkipper = true
					skipper.Datetime.IncrementMinutes(RequestCruiseLengthInMinutes)
					timestamp = *skipper.Datetime // copy timestamp
					timestamp.RoundUp()
					assignSkipper(&schedule, &request, &skipper, &timestamp)
				}
				if !cPastTime {
					log.Printf("Couldn't assign %s since the new booking would be past 8 ", skipper.Name)
				}
				if !cAccum {
					log.Printf("Couldn't assign %s since accumulated hours would go past skipper's limit", skipper.Name)
				}
			}
		}
		if !foundSkipper {
			log.Printf("No skipper found for client %s, leaving unassigned.", request.Name)
			assignSkipper(&schedule, &request, nil, &timestamp)
		}
	}

	// Sort out outputlists
	slices.SortFunc(skippers, SortSkippers)
	slices.SortFunc(schedule, SortSchedule)

	// Update the files
	// requestsFile.UpdateListing(AsStringList(requests))
	skippersFile.UpdateListing(AsStringList(skippers))
	scheduleFile.UpdateListing(AsStringList(schedule))

	return nil
}
