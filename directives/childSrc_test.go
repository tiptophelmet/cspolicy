package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyChildSrc(t *testing.T) {
	gotVal := ChildSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestChildSrc(t *testing.T) {
	gotVal := ChildSrc(
		src.Host("cdn.example.com/assets"),
		src.Host("resources.example.com/artifacts"),
	)

	wantVal := "child-src cdn.example.com/assets resources.example.com/artifacts;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
