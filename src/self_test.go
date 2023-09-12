// Sources for Content-Security-Policy directives.
package src

import "testing"

func TestSelf(t *testing.T) {
	gotVal := Self().String()
	wantVal := "'self'"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
