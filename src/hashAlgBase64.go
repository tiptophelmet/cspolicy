package src

import (
	"fmt"
	"strings"

	"github.com/tiptophelmet/cspolicy/src/hashalg"
)

type HashAlgBase64Val struct {
	alg    hashalg.HashAlgorithm
	base64 string
}

func (v *HashAlgBase64Val) String() string {
	prefix := "cspolicy.src.HashAlgBase64"

	if len(v.base64) == 0 {
		fmt.Printf("[%s] provided base64 string is empty, skipping...\n", prefix)
		return ""
	}

	return fmt.Sprintf("'%s-%s'", v.alg.GetAlgorithmID(), v.base64)
}

func HashAlgBase64(alg hashalg.HashAlgorithm, base64 string) *HashAlgBase64Val {
	return &HashAlgBase64Val{alg, strings.TrimSpace(base64)}
}
