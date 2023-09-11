package hashalg

type HashAlgorithm interface {
	GetAlgorithmID() string
}

type sha256 struct {
}

func (alg *sha256) GetAlgorithmID() string {
	return "sha256"
}

func Sha256() *sha256 {
	return &sha256{}
}

type sha384 struct {
}

func (alg *sha384) GetAlgorithmID() string {
	return "sha384"
}

func Sha384() *sha384 {
	return &sha384{}
}

type sha512 struct {
}

func (alg *sha512) GetAlgorithmID() string {
	return "sha512"
}

func Sha512() *sha512 {
	return &sha512{}
}
