package src

import "testing"

func TestReportSample(t *testing.T) {
	gotVal := ReportSample().String()
	wantVal := "'report-sample'"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
