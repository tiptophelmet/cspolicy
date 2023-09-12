// Sources for Content-Security-Policy directives.
package src

import "testing"

func TestJoin(t *testing.T) {
	gotVal := Join(
		Self(),
		Host("*.example.com"),
		Host("https://*.example.com:12/path/to/file.js"),
	)

	wantVal := "'self' *.example.com https://*.example.com:12/path/to/file.js"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
