package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyStyleSrcElem(t *testing.T) {
	gotVal := StyleSrcElem()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestStyleSrcElem(t *testing.T) {
	gotVal := StyleSrcElem(
		src.Host("cdn.example.com"),
		src.Host("front.example.com"),
	)

	wantVal := "style-src-elem cdn.example.com front.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
