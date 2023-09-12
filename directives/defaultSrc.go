// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'default-src' directive.
// The 'default-src' directive sets a default source list for a number of directives. 
// If those directives do not have an explicit source list, they will fall back to using this default list.
// See: [CSP.Directives.default-src].
//
// [CSP.Directives.default-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/default-src
func DefaultSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("default-src %s;", vals)
}
