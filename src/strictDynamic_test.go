package src

import "testing"

func TestStrictDynamic(t *testing.T) {
	gotVal := StrictDynamic().String()
	wantVal := "'strict-dynamic'"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
