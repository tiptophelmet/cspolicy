package directives

import (
	"fmt"
	"strings"
)

// Constructs the CSP 'report-to' directive.
// The 'report-to' directive specifies a group (defined in the `Report-To` header) to which to report attempts to violate the Content Security Policy.
// See: [CSP.Directives.report-to].
//
// [CSP.Directives.report-to]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/report-to
func ReportTo(endpointName string) string {
	endpointName = strings.TrimSpace(endpointName)
	if len(endpointName) == 0 {
		return ""
	}

	return fmt.Sprintf("report-to %s;", endpointName)
}
