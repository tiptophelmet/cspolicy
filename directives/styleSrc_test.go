package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyStyleSrc(t *testing.T) {
	gotVal := StyleSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestStyleSrc(t *testing.T) {
	gotVal := StyleSrc(
		src.Host("assets.example.com"),
		src.Host("cdn.example.com"),
	)

	wantVal := "style-src assets.example.com cdn.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
