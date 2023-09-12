// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'script-src-elem' directive.
// The 'script-src-elem' directive specifies valid sources for scripts loaded using the `<script>` element.
// This directive is a fallback to `script-src` when not specified.
// See: [CSP.Directives.script-src-elem].
//
// [CSP.Directives.script-src-elem]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/script-src-elem
func ScriptSrcElem(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("script-src-elem %s;", vals)
}
