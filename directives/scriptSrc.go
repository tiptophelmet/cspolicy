package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'script-src' directive.
// The 'script-src' directive specifies valid sources for JavaScript.
// This includes not only URLs loaded directly into `<script>` elements, but also things like inline script event handlers (onclick) and XSLT stylesheets which can trigger script execution.
// See: [CSP.Directives.script-src].
//
// [CSP.Directives.script-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/script-src
func ScriptSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("script-src %s;", vals)
}
