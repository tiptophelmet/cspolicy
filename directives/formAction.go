// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'form-action' directive.
// The 'form-action' directive restricts the URLs which can be used as the target of a form submissions from a given context.
// See: [CSP.Directives.form-action].
//
// [CSP.Directives.form-action]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/form-action
func FormAction(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("form-action %s;", vals)
}
