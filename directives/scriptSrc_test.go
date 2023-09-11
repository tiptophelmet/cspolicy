package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyScriptSrc(t *testing.T) {
	gotVal := ScriptSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestScriptSrc(t *testing.T) {
	gotVal := ScriptSrc(
		src.Host("assets.example.com"),
		src.Host("cdn.example.com"),
	)

	wantVal := "script-src assets.example.com cdn.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
