package src

// StrictDynamicVal represents the CSP 'strict-dynamic' source.
// The 'strict-dynamic' source indicates that the trust explicitly given to a script present in the markup, by accompanying it with a nonce or a hash, shall be propagated to all the scripts loaded by that root script.
// See: [CSP.Sources.strict-dynamic].
//
// [CSP.Sources.strict-dynamic]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#strict-dynamic
type StrictDynamicVal struct {
}

// Returns the string representation of the 'strict-dynamic' source.
func (v *StrictDynamicVal) String() string {
	return "'strict-dynamic'"
}

// Creates and returns a new StrictDynamicVal instance.
func StrictDynamic() *StrictDynamicVal {
	return &StrictDynamicVal{}
}
