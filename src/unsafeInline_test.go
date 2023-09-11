package src

import "testing"

func TestUnsafeInline(t *testing.T) {
	gotVal := UnsafeInline().String()
	wantVal := "'unsafe-inline'"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
