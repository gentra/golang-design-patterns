package internal

import (
	"github.com/gentra/legosamples/templatemethod/internal/entity"
)

type SalesReportActivity interface {
	FetchOrderData() ([]entity.Order, error)
	CompileReport([]entity.Order) (entity.SalesReport, error)
	StoreReportData(report entity.SalesReport) error
	SendReport(report entity.SalesReport) error
}

type SalesReporter interface {
	ReportSales(action SalesReportActivity) error
}
