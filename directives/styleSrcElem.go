// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'style-src-elem' directive.
// The 'style-src-elem' directive specifies valid sources for styles loaded using the `<link>` and `<style>` elements.
// This directive is a fallback to `style-src` when not specified.
// See: [CSP.Directives.style-src-elem].
//
// [CSP.Directives.style-src-elem]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/style-src-elem
func StyleSrcElem(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("style-src-elem %s;", vals)
}
