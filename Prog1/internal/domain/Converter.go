package domain

import (
	"Prog1/internal/intfc"
	"Prog1/internal/model"
	"fmt"
)

func AsRequestList(list []string) (reqs []model.Request, err error) {

	var (
		temp intfc.IMarshable
		fac  = model.RequestFactory{}
	)

	reqs = make([]model.Request, 0, len(list))
	for index, line := range list {
		temp, err = fac.FromString(line)
		if err != nil {
			return
		}
		if val, ok := temp.(*model.Request); ok {
			reqs = append(reqs, *val)
		} else {
			return nil, fmt.Errorf("Invalid Request line at line %d", index)
		}
	}
	return
}

func AsSkippersList(list []string) (skps []model.Skipper, err error) {

	var (
		temp intfc.IMarshable
		fac  = model.SkipperFactory{}
	)
	skps = make([]model.Skipper, 0, len(list))
	for index, line := range list {
		temp, err = fac.FromString(line)
		if err != nil {
			return
		}
		if val, ok := temp.(*model.Skipper); ok {
			skps = append(skps, *val)
		} else {
			return nil, fmt.Errorf("Invalid Skipper line at line %d", index)
		}
	}
	return
}

func AsEntryList(list []string) (schd []model.Entry, err error) {

	var (
		temp intfc.IMarshable
		fac  = model.EntryFactory{}
	)

	schd = make([]model.Entry, 0, len(list))
	for index, line := range list {
		temp, err = fac.FromString(line)
		if err != nil {
			return
		}
		if val, ok := temp.(*model.Entry); ok {
			schd = append(schd, *val)
		} else {
			return nil, fmt.Errorf("Invalid Entry line at line %d", index)
		}
	}
	return
}

func AsStringList[T intfc.IMarshable](list []T) []string {
	var lst = make([]string, 0, len(list))
	for _, marsh := range list {
		lst = append(lst, marsh.ToEntry())
	}
	return lst
}
