package src

import (
	"fmt"
	"strings"

	"github.com/tiptophelmet/cspolicy/src/hashalg"
)

// '<hash-algorithm>-<base64-value>' CSP source struct.
// See: [CSP.Sources.hash-algorithm-base64-value].
//
// [CSP.Sources.hash-algorithm-base64-value]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#hash-algorithm-base64-value
type HashAlgBase64Val struct {
	alg    hashalg.HashAlgorithm
	base64 string
}

// [src.SourceVal] interface implementation.
func (v *HashAlgBase64Val) String() string {
	prefix := "cspolicy.src.HashAlgBase64"

	if len(v.base64) == 0 {
		fmt.Printf("[%s] provided base64 string is empty, skipping...\n", prefix)
		return ""
	}

	return fmt.Sprintf("'%s-%s'", v.alg.GetAlgorithmID(), v.base64)
}

// Instantiates a struct for '<hash-algorithm>-<base64-value>' source.
func HashAlgBase64(alg hashalg.HashAlgorithm, base64 string) *HashAlgBase64Val {
	return &HashAlgBase64Val{alg, strings.TrimSpace(base64)}
}
