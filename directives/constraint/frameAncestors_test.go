package constraint

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestFrameAncestorsConstraintWithInvalidSources(t *testing.T) {
	constr := &FrameAncestorsConstraint{}
	sources := constr.Sources(
		src.ReportSample(),
		src.StrictDynamic(),
		src.WasmUnsafeEval(),
	)

	gotVal := src.Join(sources...)
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestFrameAncestorsConstraintWithInvalidSource(t *testing.T) {
	constr := &FrameAncestorsConstraint{}
	sources := constr.Sources(
		src.Self(),
		src.Host("*.example.com"),
		src.UnsafeHashes(),
	)

	gotVal := src.Join(sources...)
	wantVal := "'self' *.example.com"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestFrameAncestorsConstraint(t *testing.T) {
	constr := &FrameAncestorsConstraint{}
	sources := constr.Sources(
		src.None(),
		src.Self(),
		src.Scheme("https:"),
		src.Host("example.com:12/path/to/file.js"),
	)

	gotVal := src.Join(sources...)
	wantVal := "'none' 'self' https: example.com:12/path/to/file.js"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
