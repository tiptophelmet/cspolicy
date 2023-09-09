package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/directives/constraint"
	"github.com/tiptophelmet/cspolicy/src"
)

func FrameAncestors(c *constraint.FrameAncestorsConstraint) string {
	if len(c.Sources()) == 0 {
		return ""
	}

	vals := src.Join(c.Sources()...)
	return fmt.Sprintf("frame-ancestors %s;", vals)
}
