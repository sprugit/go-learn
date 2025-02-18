package domain

import "Prog1/internal/io"

func ProcessFiles(scheduleFile *io.EntryFile,
	requestsFile *io.EntryFile,
	skippersFile *io.EntryFile) (err error) {
	err = ProcessRequests(scheduleFile, requestsFile, skippersFile)
	if err != nil {
		return
	}
	scheduleFile.Header.IncrementTimestamp()
	skippersFile.Header.IncrementTimestamp()
	return
}
