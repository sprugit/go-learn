package model

import (
	"Prog1/internal/intfc"
	"fmt"
	"strconv"
	"strings"
)

type SkipperFactory struct {
}
type Skipper struct {
	Name      string
	Language  *Language
	Category  int
	Cost      int
	Specialty string
	MaxHours  int
	AcumHours int
	Datetime  *DateTime
}

func (skipper Skipper) ToEntry() string {
	return fmt.Sprintf("%s, (%s), %d*, %d, %s, %d, %d, (%s, %s)",
		skipper.Name,
		skipper.Language.ToEntryFormat(),
		skipper.Category,
		skipper.Cost,
		skipper.Specialty,
		skipper.MaxHours,
		skipper.AcumHours,
		skipper.Datetime.DateString(),
		skipper.Datetime.TimeString(),
	)
}

func (SkipperFactory) FromString(skipperline string) (skipper intfc.IMarshable, err error) {
	var skipperArgs = strings.Split(skipperline, ", ")
	var category, cost, maxhours, acumhours int
	var language *Language
	var datetime *DateTime
	language, err = LanguageFromString(skipperArgs[1])
	if err != nil {
		return
	}
	category, err = strconv.Atoi(strings.TrimRight(skipperArgs[2], "*"))
	if err != nil {
		return
	}
	cost, err = strconv.Atoi(skipperArgs[3])
	if err != nil {
		return
	}
	maxhours, err = strconv.Atoi(skipperArgs[5])
	if err != nil {
		return
	}
	acumhours, err = strconv.Atoi(skipperArgs[6])
	if err != nil {
		return
	}
	datetime, err = DateTimeFromString(skipperArgs[7], skipperArgs[8])
	if err != nil {
		return
	}
	skipper = &Skipper{
		Name:      skipperArgs[0],
		Language:  language,
		Category:  category,
		Cost:      cost,
		Specialty: skipperArgs[4],
		MaxHours:  maxhours,
		AcumHours: acumhours,
		Datetime:  datetime,
	}
	return
}

func (SkipperFactory) GetTypeName() string {
	return "skippers"
}

func (this *Skipper) Equals(other *Skipper) bool {
	return this.Name == other.Name &&
		this.Language.Equals(*other.Language) &&
		this.Category == other.Category &&
		this.Cost == other.Cost &&
		this.Specialty == other.Specialty &&
		this.MaxHours == other.MaxHours &&
		this.AcumHours == other.AcumHours &&
		this.Datetime.Equals(other.Datetime)
}
