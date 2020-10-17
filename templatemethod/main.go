package main

import (
	"log"

	"github.com/gentra/golang-design-patterns/templatemethod/internal/constant"
	"github.com/gentra/golang-design-patterns/templatemethod/internal/provide"
)

func main() {
	// Let's just pretend we get this recipient type from the higher caller in the call-stack
	recipient := constant.SalesReportForInvestor
	salesReportActivity := provide.SalesReporterActivity(recipient)

	salesReporter := provide.SalesReporter()

	err := salesReporter.ReportSales(salesReportActivity)
	if err != nil {
		log.Println(err)
	}
}
