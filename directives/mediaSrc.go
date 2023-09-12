package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'media-src' directive.
// The 'media-src' directive specifies valid sources for loading media using the `<audio>` and `<video>` elements.
// See: [CSP.Directives.media-src].
//
// [CSP.Directives.media-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/media-src
func MediaSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("media-src %s;", vals)
}
