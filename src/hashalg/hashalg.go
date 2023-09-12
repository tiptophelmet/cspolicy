// Enforces hashing algorithms for '<hash-algorithm>-<base64-value>' CSP source.
package hashalg

// Constraint to prevent random values for hashing algorithm.
// '<hash-algorithm>-<base64-value>' allows specific algorithms, see [CSP.Sources.hash-algorithm-base64-value].
// 
// [CSP.Sources.hash-algorithm-base64-value]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#hash-algorithm-base64-value
type HashAlgorithm interface {
	GetAlgorithmID() string
}

type sha256 struct {
}

// SHA-256's interface implementation.
func (alg *sha256) GetAlgorithmID() string {
	return "sha256"
}

// Creates SHA-256 hashing algorithm constraint. 
func Sha256() *sha256 {
	return &sha256{}
}

type sha384 struct {
}

// SHA-384's interface implementation.
func (alg *sha384) GetAlgorithmID() string {
	return "sha384"
}

// Creates SHA-384 hashing algorithm constraint. 
func Sha384() *sha384 {
	return &sha384{}
}

type sha512 struct {
}

// SHA-512's interface implementation.
func (alg *sha512) GetAlgorithmID() string {
	return "sha512"
}

// Creates SHA-512 hashing algorithm constraint. 
func Sha512() *sha512 {
	return &sha512{}
}
