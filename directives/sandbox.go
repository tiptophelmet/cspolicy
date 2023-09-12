// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"
	"strings"

	"github.com/tiptophelmet/cspolicy/src/sandbox"
)

// Constructs the CSP 'sandbox' directive.
// The 'sandbox' directive enables a sandbox for the requested resource similar to the `<iframe>` sandbox attribute.
// It applies restrictions to a page's actions including preventing pop-ups, preventing the execution of plugins and scripts, and enforcing a same-origin policy.
// See: [CSP.Directives.sandbox].
//
// [CSP.Directives.sandbox]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/sandbox
func Sandbox(val sandbox.SandboxVal) string {
	sandboxVal := strings.TrimSpace(string(val))
	if len(sandboxVal) == 0 {
		return ""
	}

	return fmt.Sprintf("sandbox %s;", sandboxVal)
}
