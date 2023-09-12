package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyBaseURI(t *testing.T) {
	gotVal := BaseURI()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestBaseURI(t *testing.T) {
	gotVal := BaseURI(
		src.Self(),
		src.Host("*.example.com"),
	)

	wantVal := "base-uri 'self' *.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
