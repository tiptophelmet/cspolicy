package src

import "github.com/tiptophelmet/cspolicy/src/scheme"

type SchemeVal struct {
	sc scheme.Scheme
}

func (v *SchemeVal) String() string {
	return v.sc.GetSchemeID()
}

func Scheme(sc scheme.Scheme) *SchemeVal {
	return &SchemeVal{sc}
}
