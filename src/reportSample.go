// Sources for Content-Security-Policy directives.
package src

// ReportSampleVal represents the CSP 'report-sample' source.
// The 'report-sample' source indicates that a sample of the violating code should be included in the violation report.
// See: [CSP.Sources.report-sample].
//
// [CSP.Sources.report-sample]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources#report-sample
type ReportSampleVal struct {
}

// Returns the string representation of the 'report-sample' source.
func (v *ReportSampleVal) String() string {
	return "'report-sample'"
}

// Creates and returns a new ReportSampleVal instance.
func ReportSample() *ReportSampleVal {
	return &ReportSampleVal{}
}
