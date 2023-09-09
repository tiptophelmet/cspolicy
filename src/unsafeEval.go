package src

type UnsafeEvalVal struct {
}

func (v *UnsafeEvalVal) String() string {
	return "'unsafe-eval'"
}

func UnsafeEval() *UnsafeEvalVal {
	return &UnsafeEvalVal{}
}
