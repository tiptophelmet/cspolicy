// Directives for Content-Security-Policy headers.
package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'manifest-src' directive.
// The 'manifest-src' directive specifies which manifest can be applied to the resource.
// See: [CSP.Directives.manifest-src].
//
// [CSP.Directives.manifest-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/manifest-src
func ManifestSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("manifest-src %s;", vals)
}
