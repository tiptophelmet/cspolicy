package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyFontSrc(t *testing.T) {
	gotVal := FontSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestFontSrc(t *testing.T) {
	gotVal := FontSrc(
		src.Host("assets.example.com"),
	)

	wantVal := "font-src assets.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
