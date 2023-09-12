// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/directives/constraint"
	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'frame-ancestors' directive.
// The 'frame-ancestors' directive specifies valid parents that may embed a page using `<frame>`, `<iframe>`, `<object>`, `<embed>`, or `<applet>`.
// See: [CSP.Directives.frame-ancestors].
//
// [CSP.Directives.frame-ancestors]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/frame-ancestors
func FrameAncestors(c *constraint.FrameAncestorsConstraint) string {
	if len(c.Sources()) == 0 {
		return ""
	}

	vals := src.Join(c.Sources()...)
	return fmt.Sprintf("frame-ancestors %s;", vals)
}
