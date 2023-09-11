package directives

import (
	"testing"
)

func TestEmptyReportTo(t *testing.T) {
	gotVal := ReportTo("   ")
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestReportTo(t *testing.T) {
	gotVal := ReportTo("csp-collect-endpoint")

	wantVal := "report-to csp-collect-endpoint;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
