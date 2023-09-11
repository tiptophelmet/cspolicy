package src

import (
	"fmt"
	"strings"
)

type NonceBase64Val struct {
	base64 string
}

func (v *NonceBase64Val) String() string {
	prefix := "cspolicy.src.NonceBase64"

	if len(v.base64) == 0 {
		fmt.Printf("[%s] provided base64 string is empty, skipping...\n", prefix)
		return ""
	}

	return fmt.Sprintf("'nonce-%s'", v.base64)
}

func NonceBase64(base64 string) *NonceBase64Val {
	return &NonceBase64Val{strings.TrimSpace(base64)}
}
