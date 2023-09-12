// Sources for Content-Security-Policy directives.
package src

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/tiptophelmet/cspolicy/src/hashalg"
)

var validBase64Str string

func TestEmptyBase64(t *testing.T) {
	hashAlgBase64 := HashAlgBase64(hashalg.Sha256(), "   ")
	gotVal := hashAlgBase64.String()

	if gotVal != "" {
		t.Errorf("got: %s, want: empty string", gotVal)
	}
}

func TestHashAlgBase64(t *testing.T) {
	scriptContent := `
		function showMessage() {
			alert('Hello world!');
		}

		showMessage();
	`

	hash := sha256.Sum256([]byte(scriptContent))
	validBase64Str = base64.StdEncoding.EncodeToString(hash[:])

	hashAlgBase64 := HashAlgBase64(hashalg.Sha256(), validBase64Str)

	gotVal := hashAlgBase64.String()
	wantVal := fmt.Sprintf("'%s-%s'", hashalg.Sha256().GetAlgorithmID(), validBase64Str)

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
