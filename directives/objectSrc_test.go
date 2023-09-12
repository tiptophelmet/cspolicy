package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyObjectSrc(t *testing.T) {
	gotVal := ObjectSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestObjectSrc(t *testing.T) {
	gotVal := ObjectSrc(
		src.Host("embed.example.com"),
	)

	wantVal := "object-src embed.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
