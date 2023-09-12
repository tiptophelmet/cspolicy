package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/directives/constraint"
	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyFrameAncestors(t *testing.T) {
	constr := &constraint.FrameAncestorsConstraint{}

	gotVal := FrameAncestors(constr)
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestFrameAncestors(t *testing.T) {
	constr := &constraint.FrameAncestorsConstraint{}
	constr.Sources(src.None())

	gotVal := FrameAncestors(constr)
	wantVal := "frame-ancestors 'none';"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
