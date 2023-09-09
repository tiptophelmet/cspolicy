package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src/sandbox"
)

func Sandbox(val sandbox.SandboxVal) string {
	if len(val) == 0 {
		return ""
	}

	return fmt.Sprintf("sandbox %s;", val)
}
