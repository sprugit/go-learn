package io

import (
	"Prog1/internal/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Note about interface types: if the struct implementing the interface
has its methods defined with (s *Struct) then, when using the interface
in generics, you need to pass the generic as *Struct for the method contract
to be Upheld, since (s Struct) doesn't uphold the contract
*/

const FILE_EXT = ".txt"
const TIMESTAMP_LEN = 5

type EntryFile struct {
	path     string
	Filename string
	Header   *model.Header
	Entries  []string
}

func retrieveFileTypeFromName(filename string) (fileType string) {
	var pathBits = strings.Split(filename, string(os.PathSeparator))
	fileType = pathBits[len(pathBits)-1][:(TIMESTAMP_LEN+len(FILE_EXT))-1]
	return
}

func retrieveTimestampFromName(filename string) (fileTimestamp []string) {
	var leftTrimIndex, rightTrimIndex int
	leftTrimIndex = len(filename) - (TIMESTAMP_LEN + len(FILE_EXT))
	rightTrimIndex = len(filename) - len(FILE_EXT)
	fileTimestamp = strings.Split(filename[leftTrimIndex:rightTrimIndex], "h")
	return
}

func ValidateHeader(filename string, header *model.Header) (err error) {

	time := retrieveTimestampFromName(filename)
	filetype := retrieveFileTypeFromName(filename)
	var hours, minutes int
	hours, err = strconv.Atoi(time[0])
	if err != nil {
		return
	}
	minutes, err = strconv.Atoi(time[1])
	if err != nil {
		return
	}
	if !(header.Timestamp.Hours == hours && header.Timestamp.Minutes == minutes) {
		err = fmt.Errorf("Input file error: time inconsistency between file name and header in file %s.", filename)
	}
	if strings.ToLower(header.FileType) != filetype {
		err = fmt.Errorf("Input file error: scope inconsistency between file name and header in file %s.", filename)
	}
	return
}

func ReadFile(path string) (file *EntryFile, err error) {

	var (
		header  *model.Header
		entries []string
	)

	f, err := os.Open(path)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)
	headerStr := make([]string, 0)
	for i := 0; i < model.HeaderLines; i++ {
		scanner.Scan()
		if i == model.CompanyLine || i == model.DayLine || i == model.TimeLine ||
			i == model.FileTypeLine {
			headerStr = append(headerStr, scanner.Text())
		}
	}
	header, err = model.HeaderFromString(headerStr)
	if err != nil {
		return
	}
	err = ValidateHeader(f.Name(), header)
	if err != nil {
		return
	}
	entries = make([]string, 0, 10) // type, int, capacity
	for scanner.Scan() {
		entries = append(entries, scanner.Text())
	}
	file = &EntryFile{
		path:     path,
		Filename: f.Name(),
		Header:   header,
		Entries:  entries,
	}
	return
}

func (entryFile *EntryFile) UpdateListing(listing []string) {
	entryFile.Entries = listing
}

func (entryFile *EntryFile) FileToString() (contents string) {
	contents = fmt.Sprintf("%s\n%s",
		entryFile.Header.ToString(),
		strings.Join(entryFile.Entries, "\n"))
	return
}

func (entryFile *EntryFile) WriteFileToPath(path string) (err error) {
	var (
		file *os.File
	)
	file, err = os.Create(path)
	if err != nil {
		return
	}
	var writer = bufio.NewWriter(file)
	_, err = writer.WriteString(entryFile.FileToString())
	return
}

func (entryFile *EntryFile) WriteFile() (err error) {
	var (
		file *os.File
	)

	var path = entryFile.path
	var pathArgs = strings.Split(path, string(os.PathSeparator))
	path = strings.Join(
		pathArgs[:cap(pathArgs)],
		string(os.PathSeparator))
	var time = strings.Replace(entryFile.Header.Timestamp.TimeString(), ":", "h", 1)
	path = path + string(os.PathSeparator) + entryFile.Header.GetFileType() + time + FILE_EXT
	file, err = os.Create(path)
	if err != nil {
		return
	}
	var writer = bufio.NewWriter(file)
	_, err = writer.WriteString(entryFile.FileToString())
	return err
}
