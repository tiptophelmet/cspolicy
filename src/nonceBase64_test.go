// Sources for Content-Security-Policy directives.
package src

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestEmptyNonceBase64(t *testing.T) {
	nonceBase64 := NonceBase64("    ")
	gotVal := nonceBase64.String()

	if gotVal != "" {
		t.Errorf("got: %s, want: empty string", gotVal)
	}
}

func TestNonceBase64(t *testing.T) {
	nonce := make([]byte, 16)
	rand.Read(nonce)

	base64 := base64.StdEncoding.EncodeToString(nonce)
	nonceBase64 := NonceBase64(base64)

	gotVal := nonceBase64.String()
	wantVal := fmt.Sprintf("'nonce-%s'", base64)

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
