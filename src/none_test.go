package src

import "testing"

func TestNone(t *testing.T) {
	gotVal := None().String()
	wantVal := "'none'"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
