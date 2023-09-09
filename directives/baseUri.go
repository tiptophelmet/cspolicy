package directives

import (
	"fmt"

	"github.com/tiptophelmet/cspolicy/src"
)

func BaseURI(sources ...src.SourceVal) string {
	if len(sources) == 0 {
		return ""
	}

	vals := src.Join(sources...)
	return fmt.Sprintf("base-uri %s;", vals)
}
