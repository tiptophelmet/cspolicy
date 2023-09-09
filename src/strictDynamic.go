package src

type StrictDynamicVal struct {
}

func (v *StrictDynamicVal) String() string {
	return "'strict-dynamic'"
}

func StrictDynamic() *StrictDynamicVal {
	return &StrictDynamicVal{}
}
