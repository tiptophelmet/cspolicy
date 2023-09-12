// Sources for Content-Security-Policy directives.
package src

import (
	"regexp"
)

// SchemeVal represents the CSP scheme source.
// The scheme source allows content to be loaded from a specific URI scheme, such as 'http:' or 'https:'.
// See: [CSP.Sources.scheme].
//
// [CSP.Sources.scheme]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#scheme
type SchemeVal struct {
	sc string
}

// Returns the string representation of the scheme source.
func (v *SchemeVal) String() string {
	schemePattern := "^[a-zA-Z0-9+.-]+:$"

	regex, _ := regexp.Compile(schemePattern)
	if matched := regex.MatchString(v.sc); matched {
		return v.sc
	}

	return ""
}

// Creates and returns a new SchemeVal instance.
func Scheme(scheme string) *SchemeVal {
	return &SchemeVal{sc: scheme}
}
