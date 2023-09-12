// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'connect-src' directive.
// The 'connect-src' directive restricts the URLs which can be loaded using script interfaces.
// This function takes in a variable number of sources and returns the formatted 'connect-src' directive string.
// See: [CSP.Directives.connect-src].
//
// [CSP.Directives.connect-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/connect-src
func ConnectSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("connect-src %s;", vals)
}
