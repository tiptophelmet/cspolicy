// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'script-src-attr' directive.
// The 'script-src-attr' directive specifies valid sources for inline scripts in event handlers and the `javascript:` pseudo-protocol.
// This directive is a fallback to `script-src` when not specified.
// See: [CSP.Directives.script-src-attr].
//
// [CSP.Directives.script-src-attr]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/script-src-attr
func ScriptSrcAttr(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("script-src-attr %s;", vals)
}
