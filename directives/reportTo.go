package directives

import (
	"fmt"
	"strings"
)

func ReportTo(endpointName string) string {
	endpointName = strings.TrimSpace(endpointName)
	if len(endpointName) == 0 {
		return ""
	}

	return fmt.Sprintf("report-to %s;", endpointName)
}
