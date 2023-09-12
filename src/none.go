// Sources for Content-Security-Policy directives.
package src


// NoneVal represents the CSP 'none' source.
// The 'none' source ensures that no content can be loaded from any source.
// See: [CSP.Sources.none].
//
// [CSP.Sources.none]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#none
type NoneVal struct {
}

// Returns the string representation of the 'none' source.
func (v *NoneVal) String() string {
	return "'none'"
}

// Creates and returns a new NoneVal instance.
func None() *NoneVal {
	return &NoneVal{}
}
