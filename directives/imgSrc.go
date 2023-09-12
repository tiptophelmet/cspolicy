package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

// Constructs the CSP 'img-src' directive.
// The 'img-src' directive specifies valid sources of images and favicons.
// See: [CSP.Directives.img-src].
//
// [CSP.Directives.img-src]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/img-src
func ImgSrc(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("img-src %s;", vals)
}
