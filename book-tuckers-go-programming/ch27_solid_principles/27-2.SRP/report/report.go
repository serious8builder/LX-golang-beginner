package main

import "fmt"

type Report interface {
	Report() string
}

type FinanceReport struct {
	report string
}

func (r *FinanceReport) Report() string {
	return r.report
}

type ReportSender struct {
	method string
}

func (s *ReportSender) SendReport(report Report) {
	fmt.Printf("Send %s : %s", s.method, report.Report())
}

func main() {
	r := FinanceReport{"Report 2022-07: ..."}
	rs := ReportSender{
		method: "Email",
	}

	rs.SendReport(&r)
}
