package src

import (
	"strings"
)

// Joins multiple CSP sources into a single string.
// This function aids in creating a consolidated CSP string from multiple sources.
func Join(sources ...SourceVal) string {
	var joined string
	for _, src := range sources {
		joined += src.String() + " "
	}

	return strings.TrimSpace(joined)
}
