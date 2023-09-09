package src

import "fmt"

type NonceBase64Val struct {
	base64 string
}

func (v *NonceBase64Val) String() string {
	return fmt.Sprintf("'nonce-%s'", v.base64)
}

func NonceBase64(base64 string) *NonceBase64Val {
	return &NonceBase64Val{base64}
}
