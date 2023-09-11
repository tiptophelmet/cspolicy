package constraint

import (
	"fmt"
	"reflect"

	"github.com/tiptophelmet/cspolicy/src"
)

type FrameAncestorsConstraint struct {
	validated []src.SourceVal
}

func (c *FrameAncestorsConstraint) validate(source src.SourceVal) bool {
	switch source.(type) {
	case *src.HostVal:
	case *src.SchemeVal:
	case *src.SelfVal:
	case *src.NoneVal:
	default:
		prefix := "cspolicy.constraint.FrameAncestorsConstraint"

		format := "[%s] restricted source %s for frame-ancestors (value: %s), skipping...\n"
		fmt.Printf(format, prefix, reflect.TypeOf(source).Elem(), source.String())
		return false
	}

	return true
}

func (c *FrameAncestorsConstraint) Sources(sources ...src.SourceVal) []src.SourceVal {
	if len(c.validated) != 0 {
		return c.validated
	}

	for _, src := range sources {
		if c.validate(src) {
			c.validated = append(c.validated, src)
		}
	}

	return c.validated
}
