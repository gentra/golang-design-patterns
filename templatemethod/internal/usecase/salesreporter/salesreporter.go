package salesreporter

import (
	"github.com/gentra/golang-design-patterns/templatemethod/internal"
	"log"
)

type SalesReporter struct {
}

func NewSalesReporter() *SalesReporter {
	return &SalesReporter{}
}

func (s *SalesReporter) ReportSales(action internal.SalesReportActivity) error {
	orders, err := action.FetchOrderData()
	if err != nil {
		return err
	}

	report, err := action.CompileReport(orders)
	if err != nil {
		return err
	}

	err = action.StoreReportData(report)
	if err != nil {
		return err
	}

	err = action.SendReport(report)
	if err != nil {
		return err
	}

	log.Println("Storing report activity results to dev-internal monitoring system")

	return nil
}
