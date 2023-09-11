package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyDefaultSrc(t *testing.T) {
	gotVal := DefaultSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestDefaultSrc(t *testing.T) {
	gotVal := DefaultSrc(
		src.None(),
	)

	wantVal := "default-src 'none';"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
