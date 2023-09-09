package cspolicy

import (
	"strings"
)

func Build(directives ...string) string {
	if len(directives) == 0 {
		return ""
	}

	joined := strings.Join(directives, " ")
	
	return strings.TrimRight(joined, ";")
}
