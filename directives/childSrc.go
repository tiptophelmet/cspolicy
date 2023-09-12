package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'child-src' directive.
// The 'child-src' directive restricts the URLs which can be loaded into worker, shared worker, or <iframe> contexts.
// This function takes in a variable number of sources and returns the formatted 'child-src' directive string.
// See: [CSP.Directives.child-src].
//
// [CSP.Directives.child-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/child-src
func ChildSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("child-src %s;", vals)
}
