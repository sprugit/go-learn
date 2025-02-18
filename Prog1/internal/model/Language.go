package model

import (
	"fmt"
	"slices"
	"strings"
)

type Language struct {
	Languages []string
}

func (l *Language) ToEntryFormat() string {

	return fmt.Sprintf("(%s)", strings.Join(l.Languages, "; "))
}

func (l1 *Language) HasMatch(l2 Language) bool {
	for _, concrete_lang1 := range l1.Languages {
		for _, concrete_lang2 := range l2.Languages {
			if concrete_lang1 == concrete_lang2 {
				return true
			}
		}
	}
	return false
}

func LanguageFromString(langlist string) (lang *Language, err error) {
	langlist = strings.Trim(langlist, "()")
	lang = &Language{
		Languages: strings.Split(langlist, "; "),
	}
	return
}

func (this Language) Equals(other Language) bool {
	if len(this.Languages) != len(other.Languages) {
		return false
	}
	for _, langStr := range this.Languages {
		if !slices.Contains(other.Languages, langStr) {
			return false
		}
	}
	return true
}
