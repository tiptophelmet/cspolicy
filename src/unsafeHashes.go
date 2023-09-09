package src

type UnsafeHashesVal struct {
}

func (v *UnsafeHashesVal) String() string {
	return "'unsafe-hashes'"
}

func UnsafeHashes() *UnsafeHashesVal {
	return &UnsafeHashesVal{}
}
