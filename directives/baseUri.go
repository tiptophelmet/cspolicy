// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'base-uri' directive.
// The 'base-uri' directive restricts the URLs which can be used in a document's `<base>` element.
// This function takes in a variable number of sources and returns the formatted 'base-uri' directive string.
// See: [CSP.Directives.base-uri].
//
// [CSP.Directives.base-uri]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/base-uri
func BaseURI(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("base-uri %s;", vals)
}
