package src

import "testing"

func TestUnsafeHashes(t *testing.T) {
	gotVal := UnsafeHashes().String()
	wantVal := "'unsafe-hashes'"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
