package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'worker-src' directive.
// The 'worker-src' directive specifies valid sources for Worker, SharedWorker, or ServiceWorker scripts.
// This directive is a fallback to `script-src` when not specified.
// See: [CSP.Directives.worker-src].
//
// [CSP.Directives.worker-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/worker-src
func WorkerSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("worker-src %s;", vals)
}
