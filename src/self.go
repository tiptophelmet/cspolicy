package src

type SelfVal struct {
}

func (v *SelfVal) String() string {
	return "'self'"
}

func Self() *SelfVal {
	return &SelfVal{}
}
