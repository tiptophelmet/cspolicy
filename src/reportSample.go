package src

type ReportSampleVal struct {
}

func (v *ReportSampleVal) String() string {
	return "'report-sample'"
}

func ReportSample() *ReportSampleVal {
	return &ReportSampleVal{}
}
