package src

// SelfVal represents the CSP 'self' source.
// The 'self' source allows content to be loaded from the same origin as the document.
// See: [CSP.Sources.self].
//
// [CSP.Sources.self]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#self
type SelfVal struct {
}

// Returns the string representation of the 'self' source.
func (v *SelfVal) String() string {
	return "'self'"
}

// Creates and returns a new SelfVal instance.
func Self() *SelfVal {
	return &SelfVal{}
}
