// Directives for Content-Security-Policy headers.
package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src/sandbox"
)

func TestEmptySandbox(t *testing.T) {
	gotVal := Sandbox("   ")
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestSandbox(t *testing.T) {
	gotVal := Sandbox(sandbox.AllowDownloads)

	wantVal := "sandbox allow-downloads;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
