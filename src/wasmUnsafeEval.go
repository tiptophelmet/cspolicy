package src

// WasmUnsafeEvalVal represents the CSP 'wasm-unsafe-eval' source.
// The 'wasm-unsafe-eval' source allows the use of WebAssembly text format (WAT) as a source for WebAssembly modules.
// See: [CSP.Sources.wasm-unsafe-eval].
//
// [CSP.Sources.wasm-unsafe-eval]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#wasm-unsafe-eval
type WasmUnsafeEvalVal struct {
}

// Returns the string representation of the 'wasm-unsafe-eval' source.
func (v *WasmUnsafeEvalVal) String() string {
	return "'wasm-unsafe-eval'"
}

// Creates and returns a new WasmUnsafeEvalVal instance.
func WasmUnsafeEval() *WasmUnsafeEvalVal {
	return &WasmUnsafeEvalVal{}
}
