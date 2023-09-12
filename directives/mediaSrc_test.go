package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyMediaSrc(t *testing.T) {
	gotVal := MediaSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestMediaSrc(t *testing.T) {
	gotVal := MediaSrc(
		src.Host("media.example.com"),
		src.Host("assets.example.com"),
	)

	wantVal := "media-src media.example.com assets.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
