// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'frame-src' directive.
// The 'frame-src' directive specifies valid sources for nested browsing contexts loading using elements such as `<frame>` and `<iframe>`.
// See: [CSP.Directives.frame-src].
//
// [CSP.Directives.frame-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/frame-src
func FrameSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("frame-src %s;", vals)
}
