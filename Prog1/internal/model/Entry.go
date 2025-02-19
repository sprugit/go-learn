package model

import (
	"Prog1/internal/intfc"
	"fmt"
	"strconv"
	"strings"
)

type EntryFactory struct {
}
type Entry struct {
	Timestamp       *DateTime
	DurationInHours int
	SkipperName     string
	Cost            int
	RequestName     string
}

func (entry Entry) ToEntry() string {
	if entry.SkipperName == "" {
		return fmt.Sprintf("%s, not-assigned, %s",
			entry.Timestamp.DateString(),
			entry.RequestName)
	}
	return fmt.Sprintf("%s, %s, %d, %s, %d, %s",
		entry.Timestamp.DateString(),
		entry.Timestamp.TimeString(),
		entry.DurationInHours,
		entry.SkipperName,
		entry.Cost,
		entry.RequestName,
	)
}

func (EntryFactory) FromString(entry string) (e intfc.IMarshable, err error) {

	var args = strings.Split(entry, ", ")
	var datetime *DateTime
	datetime, err = DateTimeFromString(args[0], args[1])
	if err != nil {
		return
	}
	var duration int
	duration, err = strconv.Atoi(args[2])
	if err != nil {
		return
	}
	var cost int
	cost, err = strconv.Atoi(args[4])
	if err != nil {
		return
	}
	e = &Entry{
		Timestamp:       datetime,
		DurationInHours: duration,
		SkipperName:     args[3],
		Cost:            cost,
		RequestName:     args[5],
	}
	return
}

func (EntryFactory) GetTypeName() string {
	return "schedule"
}

func (this Entry) Equals(other Entry) bool {
	return this.Timestamp.Equals(other.Timestamp) &&
		this.DurationInHours == other.DurationInHours &&
		this.SkipperName == other.SkipperName &&
		this.Cost == other.Cost &&
		this.RequestName == other.RequestName
}
