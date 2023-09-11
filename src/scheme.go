package src

import (
	"regexp"
)

type SchemeVal struct {
	sc string
}

func (v *SchemeVal) String() string {
	schemePattern := "^[a-zA-Z0-9+.-]+:$"

	regex, _ := regexp.Compile(schemePattern)
	if matched := regex.MatchString(v.sc); matched {
		return v.sc
	}

	return ""
}

func Scheme(scheme string) *SchemeVal {
	return &SchemeVal{sc: scheme}
}
