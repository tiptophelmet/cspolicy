package src

type NoneVal struct {
}

func (v *NoneVal) String() string {
	return "'none'"
}

func None() *NoneVal {
	return &NoneVal{}
}
