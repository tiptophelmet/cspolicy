package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyFrameSrc(t *testing.T) {
	gotVal := FrameSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestFrameSrc(t *testing.T) {
	gotVal := FrameSrc(
		src.Host("notes.example.com"),
		src.Host("viewbox.example.com"),
	)

	wantVal := "frame-src notes.example.com viewbox.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
