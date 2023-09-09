package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

func ScriptSrcElem(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("script-src-elem %s;", vals)
}
