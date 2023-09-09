package src

type WasmUnsafeEvalVal struct {
}

func (v *WasmUnsafeEvalVal) String() string {
	return "'wasm-unsafe-eval'"
}

func WasmUnsafeEval() *WasmUnsafeEvalVal {
	return &WasmUnsafeEvalVal{}
}
