// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'style-src' directive.
// The 'style-src' directive specifies valid sources for stylesheets.
// This includes not only URLs loaded into `<link>` elements, but also inline `<style>` elements and styles applied via DOM APIs.
// See: [CSP.Directives.style-src].
//
// [CSP.Directives.style-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/style-src
func StyleSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("style-src %s;", vals)
}
