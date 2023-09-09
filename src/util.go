package src

import (
	"strings"
)

func Join(sources ...SourceVal) string {
	var joined string
	for _, src := range sources {
		joined += src.String() + " "
	}

	return strings.TrimSpace(joined)
}
