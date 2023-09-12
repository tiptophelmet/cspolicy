// Directives for Content-Security-Policy headers.
package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyWorkerSrc(t *testing.T) {
	gotVal := WorkerSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestWorkerSrc(t *testing.T) {
	gotVal := WorkerSrc(
		src.Host("example.com/worker"),
		src.Host("collect.example.com/worker"),
	)

	wantVal := "worker-src example.com/worker collect.example.com/worker;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
