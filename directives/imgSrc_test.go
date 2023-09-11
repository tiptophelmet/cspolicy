package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyImgSrc(t *testing.T) {
	gotVal := ImgSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestImgSrc(t *testing.T) {
	gotVal := ImgSrc(
		src.Self(),
		src.Scheme("data:"),
		src.Host("media.example.com"),
		src.Host("avatars.example.com"),
	)

	wantVal := "img-src 'self' data: media.example.com avatars.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
