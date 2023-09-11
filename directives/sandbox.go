package directives

import (
	"fmt"
	"strings"

	"github.com/tiptophelmet/cspolicy/src/sandbox"
)

func Sandbox(val sandbox.SandboxVal) string {
	sandboxVal := strings.TrimSpace(string(val))
	if len(sandboxVal) == 0 {
		return ""
	}

	return fmt.Sprintf("sandbox %s;", sandboxVal)
}
