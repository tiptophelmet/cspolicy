package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyScriptSrcElem(t *testing.T) {
	gotVal := ScriptSrcElem()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestScriptSrcElem(t *testing.T) {
	gotVal := ScriptSrcElem(
		src.Host("cdnjs.example.com"),
		src.Host("front.example.com"),
	)

	wantVal := "script-src-elem cdnjs.example.com front.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
