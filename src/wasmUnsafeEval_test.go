// Sources for Content-Security-Policy directives.
package src

import "testing"

func TestWasmUnsafeEval(t *testing.T) {
	gotVal := WasmUnsafeEval().String()
	wantVal := "'wasm-unsafe-eval'"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
