package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyStyleSrcAttr(t *testing.T) {
	gotVal := StyleSrcAttr()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestStyleSrcAttr(t *testing.T) {
	gotVal := StyleSrcAttr(
		src.Host("cdn.example.com"),
		src.Host("front.example.com"),
	)

	wantVal := "style-src-attr cdn.example.com front.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
