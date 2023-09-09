package src

type UnsafeInlineVal struct {
}

func (v *UnsafeInlineVal) String() string {
	return "'unsafe-inline'"
}

func UnsafeInline() *UnsafeInlineVal {
	return &UnsafeInlineVal{}
}
