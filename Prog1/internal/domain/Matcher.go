package domain

import (
	"Prog1/internal/model"
	"log"
	"slices"
)

const CruiseLength = 30

func createScheduleEntry(schedule *[]model.Entry,
	request *model.Request,
	skipper *model.Skipper,
	headerStamp *model.DateTime) {

	var (
		tripStart, tripEnd       model.DateTime
		RequestName, SkipperName string
		SkipperCost              int
	)

	RequestName = request.Name
	tripStart = *headerStamp.Copy()

	if skipper != nil {

		SkipperName = skipper.Name
		SkipperCost = skipper.Cost

		if tripStart.LessThan(skipper.Datetime) {
			log.Printf("Skipper is available after header timestamp")
			tripStart = *skipper.Datetime.Copy()
		}

		cPastTime := tripStart.IsBeforeClosingTime(
			model.HourAsSeconds(request.CruisePeriodHours))

		if !cPastTime {
			log.Printf("Cruise duration would result in working overtime - setting cruise to next day")
			tripStart.SetToNextDay()
		}

		tripEnd = *tripStart.Copy()
		tripEnd.IncrementHours(request.CruisePeriodHours)
		tripEnd.RoundUp()

		skipper.Datetime = &tripEnd
		skipper.AcumHours += request.CruisePeriodHours
	}
	var newEntry = model.Entry{
		Timestamp:       &tripStart,
		DurationInHours: request.CruisePeriodHours,
		SkipperName:     SkipperName,
		Cost:            SkipperCost * request.CruisePeriodHours,
		RequestName:     RequestName,
	}
	*schedule = append(*schedule, newEntry) // append returns new Slice header
}

func ProcessRequests(schedule *[]model.Entry,
	requests *[]model.Request,
	skippers *[]model.Skipper,
	headerStamp *model.DateTime) (err error) {

	log.Printf("Processing requests for %s-%s.", headerStamp.DateString(), headerStamp.TimeString())

	for indexReq, request := range *requests {
		log.Printf("Processing request #%d for client %s", indexReq, request.Name)
		slices.SortFunc(*skippers, SortSkippers)

		var foundSkipper = false

		for i := 0; i < len(*skippers) && !foundSkipper; i++ {

			var skipper *model.Skipper
			skipper = &(*skippers)[i]
			log.Printf("Attempting to match client %s with Skipper %s", request.Name, skipper.Name)

			var cLang, cSpec, cCat bool
			cLang = request.Languages.HasMatch(*skipper.Language)
			cSpec = request.Specialty == skipper.Specialty
			cCat = request.Category == skipper.Category
			if cLang && cSpec && cCat {

				log.Printf("%s matches %s's needs in terms of language, specialty and category", skipper.Name, request.Name)

				foundSkipper = (skipper.AcumHours + request.CruisePeriodHours) < skipper.MaxHours // Assign skipper if not post 8pm
				if foundSkipper {
					log.Printf("Assigning %s to %s", skipper.Name, request.Name)
					createScheduleEntry(schedule, &request, skipper, headerStamp)
				} else {
					log.Printf("Couldn't assign %s since accumulated hours would go past skipper's limit", skipper.Name)
				}
			}
		}
		if !foundSkipper {
			log.Printf("No skipper found for client %s, leaving unassigned.", request.Name)
			createScheduleEntry(schedule, &request, nil, headerStamp)
		}
	}

	// Sort out outputlists
	slices.SortFunc(*skippers, SortSkippersAlphabet)
	slices.SortFunc(*schedule, SortSchedule)

	return nil
}
