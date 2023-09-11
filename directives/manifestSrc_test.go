package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyManifestSrc(t *testing.T) {
	gotVal := ManifestSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestManifestSrc(t *testing.T) {
	gotVal := ManifestSrc(src.Self())

	wantVal := "manifest-src 'self';"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
