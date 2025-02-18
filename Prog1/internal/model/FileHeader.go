package model

import (
	"fmt"
	"strings"
)

type Header struct {
	Company   string
	Timestamp *DateTime
	FileType  string
}

const HeaderLines = 7
const CompanyLine = 1
const DayLine = 3
const TimeLine = 5
const FileTypeLine = 6
const HeaderArgs = 4

func (header *Header) ToString() string {

	return fmt.Sprintf("Company:\n%s\nDay:\n%s\nTime:\n%s\n%s:",
		header.Company,
		header.Timestamp.DateString(),
		header.Timestamp.TimeString(),
		header.FileType)
}

func HeaderFromString(header []string) (*Header, error) {
	var datetime, err = DateTimeFromString(header[1], header[2])
	if err != nil {
		return nil, err
	}
	return &Header{
		Company:   header[0],
		Timestamp: datetime,
		FileType:  strings.TrimRight(header[3], ":"),
	}, nil
}

func (header *Header) IncrementTimestamp() {
	header.Timestamp.IncrementMinutes(30)
}

func (header *Header) GetFileType() string {
	return strings.ToLower(strings.TrimRight(header.FileType, ":"))
}
