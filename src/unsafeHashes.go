// Sources for Content-Security-Policy directives.
package src

// UnsafeHashesVal represents the CSP 'unsafe-hashes' source.
// The 'unsafe-hashes' source allows inline event handlers and javascript: URIs to be executed, as long as their contents are correctly hashed.
// See: [CSP.Sources.unsafe-hashes].
//
// [CSP.Sources.unsafe-hashes]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#unsafe-hashes
type UnsafeHashesVal struct {
}

// Returns the string representation of the 'unsafe-hashes' source.
func (v *UnsafeHashesVal) String() string {
	return "'unsafe-hashes'"
}

// Creates and returns a new UnsafeHashesVal instance.
func UnsafeHashes() *UnsafeHashesVal {
	return &UnsafeHashesVal{}
}
