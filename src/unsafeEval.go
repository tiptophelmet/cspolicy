// Sources for Content-Security-Policy directives.
package src

// UnsafeEvalVal represents the CSP 'unsafe-eval' source.
// The 'unsafe-eval' source allows the use of eval() and similar methods for creating code from strings.
// See: [CSP.Sources.unsafe-eval].
//
// [CSP.Sources.unsafe-eval]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#unsafe-eval
type UnsafeEvalVal struct {
}

// Returns the string representation of the 'unsafe-eval' source.
func (v *UnsafeEvalVal) String() string {
	return "'unsafe-eval'"
}

// Creates and returns a new UnsafeEvalVal instance.
func UnsafeEval() *UnsafeEvalVal {
	return &UnsafeEvalVal{}
}
