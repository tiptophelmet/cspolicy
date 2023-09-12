package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'font-src' directive.
// The 'font-src' directive specifies valid sources for fonts loaded using @font-face.
// See: [CSP.Directives.font-src].
//
// [CSP.Directives.font-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/font-src
func FontSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("font-src %s;", vals)
}
