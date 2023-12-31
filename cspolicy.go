// User-friendly API to build your Content-Security-Policy HTTP header.
package cspolicy

import (
	"strings"
)

// Builds Content-Security-Policy.
// 
// Accepts a slice of directives, formats them and returns a CSP header value.
func Build(directives ...string) string {
	if len(directives) == 0 {
		return ""
	}

	joined := strings.Join(directives, " ")
	
	return strings.TrimRight(joined, ";")
}
