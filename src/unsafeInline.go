package src

// UnsafeInlineVal represents the CSP 'unsafe-inline' source.
// The 'unsafe-inline' source allows the use of inline resources, such as inline `<script>` elements, `javascript:` URLs, inline event handlers, and inline `<style>` elements.
// See: [CSP.Sources.unsafe-inline].
//
// [CSP.Sources.unsafe-inline]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#unsafe-inline
type UnsafeInlineVal struct {
}

// Returns the string representation of the 'unsafe-inline' source.
func (v *UnsafeInlineVal) String() string {
	return "'unsafe-inline'"
}

// Creates and returns a new UnsafeInlineVal instance.
func UnsafeInline() *UnsafeInlineVal {
	return &UnsafeInlineVal{}
}
