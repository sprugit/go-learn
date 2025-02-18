package model

import (
	"Prog1/internal/intfc"
	"fmt"
	"strconv"
	"strings"
)

type RequestFactory struct {
}
type Request struct {
	Name         string
	Languages    *Language
	Category     int
	Specialty    string
	CruisePeriod int
}

// Dummy method to  mock static factory
func (RequestFactory) FromString(requestline string) (req intfc.IMarshable, err error) {
	var requestArgs = strings.Split(requestline, ", ")
	var language *Language
	language, err = LanguageFromString(requestArgs[1])
	if err != nil {
		return
	}
	var category, period int
	category, err = strconv.Atoi(strings.TrimRight(requestArgs[2], "*"))
	if err != nil {
		return
	}
	period, err = strconv.Atoi(requestArgs[4])
	if err != nil {
		return
	}
	req = &Request{
		Name:         requestArgs[0],
		Languages:    language,
		Category:     category,
		Specialty:    requestArgs[3],
		CruisePeriod: period,
	}
	return
}

// Read only method (r Request)
func (RequestFactory) GetTypeName() string {
	return "requests"
}

// Read only method
func (r Request) ToEntry() string {
	return fmt.Sprintf("%s, (%s), %d, %s, %d",
		r.Name,
		r.Languages.ToEntryFormat(),
		r.Category,
		r.Specialty,
		r.CruisePeriod)
}

func (this Request) Equals(other Request) bool {
	return this.Name == other.Name &&
		this.Languages.Equals(*other.Languages) &&
		this.Category == other.Category &&
		this.Specialty == other.Specialty &&
		this.CruisePeriod == other.CruisePeriod
}
