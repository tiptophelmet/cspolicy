// Sources for Content-Security-Policy directives.
package src

import (
	"fmt"
	"net/url"
	"strings"
)

// HostVal represents a CSP host source.
// The host source allows for a specific domain to be whitelisted as a valid source for content.
// See: [CSP.Sources.host-source].
//
// [CSP.Sources.host-source]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#host-source
type HostVal struct {
	source string
}

// Returns the string representation of the host source.
func (v *HostVal) String() string {
	if !v.isValid() {
		return ""
	}

	return v.source
}

func (v *HostVal) isValid() bool {
	prefix := "cspolicy.src.HostSource"

	parsedURL, err := url.Parse(v.source)
	if err != nil {
		format := "[%s] '%s' is invalid, skipping... (url.Parse error: '%v')\n"
		fmt.Printf(format, prefix, v.source, err)

		return false
	}

	var padScheme string
	if parsedURL.Scheme == "" {
		padScheme = "noscheme://"
	}

	_, err = url.ParseRequestURI(padScheme + v.source)
	if err != nil {
		preparedErr := strings.ReplaceAll(err.Error(), padScheme, "")

		format := "[%s] '%s' is invalid, skipping... (url.ParseRequestURI error: '%v')\n"
		fmt.Printf(format, prefix, v.source, preparedErr)

		return false
	}

	return err == nil
}

// Instantiates a struct for <host-source> source.
func Host(source string) *HostVal {
	return &HostVal{source}
}
