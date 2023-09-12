// Sources for Content-Security-Policy directives.
package src

import (
	"fmt"
	"strings"
)

// Represents a CSP nonce source in base64 format.
// The nonce source allows for inline scripts or styles to be executed, as long as they have the specified nonce attribute.
// See: [CSP.Sources.nonce-base64].
//
// [CSP.Sources.nonce-base64]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#nonce-base64-value
type NonceBase64Val struct {
	base64 string
}

// Returns the string representation of the nonce source in base64 format.
func (v *NonceBase64Val) String() string {
	prefix := "cspolicy.src.NonceBase64"

	if len(v.base64) == 0 {
		fmt.Printf("[%s] provided base64 string is empty, skipping...\n", prefix)
		return ""
	}

	return fmt.Sprintf("'nonce-%s'", v.base64)
}

// Creates and returns a new NonceBase64Val instance.
func NonceBase64(base64 string) *NonceBase64Val {
	return &NonceBase64Val{strings.TrimSpace(base64)}
}
