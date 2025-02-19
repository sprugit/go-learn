package domain

import (
	"Prog1/internal/model"
)

/*
Weighted sort. Order: (Most Relevant to Least)
  - Must speak the language
  - Must share specialty
  - Must share category
  - 1st: Timestamp (Earliest Preferable)
  - 2nd: Cost (Cheaper preferable)
  - 3rd: AccumHours (Least Accumulated)
  - 4th: Name (First Letter Alphabet order)
*/
var SortSkippers = func(a model.Skipper, b model.Skipper) int {
	var comparison = a.Datetime.CompareTo(b.Datetime)
	if comparison != 0 {
		return comparison
	}
	comparison = a.Cost - b.Cost
	if comparison != 0 {
		return comparison
	}
	comparison = a.AcumHours - b.AcumHours
	if comparison != 0 {
		return comparison
	}
	return int(a.Name[0]) - int(b.Name[0])
}

/*
Weighted sort. Order:
  - 1st: Isn't assigned
  - 2nd: Timestamp
  - 3rd: SkipperName
*/
var SortSchedule = func(a model.Entry, b model.Entry) int {
	if a.SkipperName == "" {
		if b.SkipperName != "" {
			return -1
		}
	} else {
		if b.SkipperName == "" {
			return 1
		}
	}
	var comparison = a.Timestamp.CompareTo(b.Timestamp)
	if comparison != 0 {
		return comparison
	}
	comparison = int(a.SkipperName[0]) - int(b.SkipperName[0])
	return comparison
}

var SortSkippersAlphabet = func(a model.Skipper, b model.Skipper) int {
	return int(a.Name[0]) - int(b.Name[0])
}
