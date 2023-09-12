// Sources for Content-Security-Policy directives.
package src

import "testing"

func TestUnsafeEval(t *testing.T) {
	gotVal := UnsafeEval().String()
	wantVal := "'unsafe-eval'"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
