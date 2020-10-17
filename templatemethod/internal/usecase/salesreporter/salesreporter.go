package salesreporter

import (
	"github.com/gentra/legosamples/templatemethod/internal"
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

	return nil
}
