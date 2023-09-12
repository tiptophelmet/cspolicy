// Directives for Content-Security-Policy headers.
package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyScriptSrcAttr(t *testing.T) {
	gotVal := ScriptSrcAttr()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestScriptSrcAttr(t *testing.T) {
	gotVal := ScriptSrcAttr(
		src.Host("cdnjs.example.com"),
		src.Host("front.example.com"),
	)

	wantVal := "script-src-attr cdnjs.example.com front.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
