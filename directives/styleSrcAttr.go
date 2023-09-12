// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'style-src-attr' directive.
// The 'style-src-attr' directive specifies valid sources for inline styles applied to individual DOM elements.
// This directive is a fallback to `style-src` when not specified.
// See: [CSP.Directives.style-src-attr].
//
// [CSP.Directives.style-src-attr]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/style-src-attr
func StyleSrcAttr(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("style-src-attr %s;", vals)
}
