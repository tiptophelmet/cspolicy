package src

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src/hashalg"
)

type HashAlgBase64Val struct {
	alg    hashalg.HashAlgorithm
	base64 string
}

func (v *HashAlgBase64Val) String() string {
	return fmt.Sprintf("'%s-%s'", v.alg, v.base64)
}

func HashAlgBase64(alg hashalg.HashAlgorithm, base64 string) *HashAlgBase64Val {
	return &HashAlgBase64Val{alg, base64}
}
