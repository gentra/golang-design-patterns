package salesreportactivity

import (
	"log"
	"time"

	"github.com/gentra/legosamples/templatemethod/internal/entity"
)

type Finance struct {
}

func NewFinance() *Finance {
	return &Finance{}
}

func (f *Finance) FetchOrderData() ([]entity.Order, error) {
	log.Println("Fetching order-data from Big Data")
	return []entity.Order{{ID: "1"}}, nil
}

func (f *Finance) CompileReport(orders []entity.Order) (entity.SalesReport, error) {
	log.Println("Compiling report with special format for Finance Department")
	return entity.SalesReport{
		Header:    "Sales Report Oct 2020 - Finance Dept",
		SubHeader: "October 2020",
		Content:   "Sales report content for finance department",
		Footer:    "Finance Dept",
		Orders:    orders,
		CreatedAt: time.Now(),
	}, nil
}

func (f *Finance) StoreReportData(report entity.SalesReport) error {
	log.Println("Storing report data to Finance department database")
	log.Println("Storing report data as Ms. Word on BizDev's file storage")
	return nil
}

func (f *Finance) SendReport(report entity.SalesReport) error {
	log.Println("Sending report to Finance Dept mailing group")
	return nil
}
