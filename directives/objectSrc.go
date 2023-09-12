package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'object-src' directive.
// The 'object-src' directive specifies valid sources for the `<object>`, `<embed>`, and `<applet>` elements.
// See: [CSP.Directives.object-src].
//
// [CSP.Directives.object-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/object-src
func ObjectSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("object-src %s;", vals)
}
