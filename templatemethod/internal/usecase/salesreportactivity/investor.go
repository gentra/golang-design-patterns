package salesreportactivity

import (
	"log"
	"time"

	"github.com/gentra/legosamples/templatemethod/internal/entity"
)

type Investor struct {
}

func NewInvestor() *Investor {
	return &Investor{}
}

func (i *Investor) FetchOrderData() ([]entity.Order, error) {
	log.Println("Fetching order-data from Big Data")
	return []entity.Order{{ID: "1"}}, nil
}

func (i *Investor) CompileReport(orders []entity.Order) (entity.SalesReport, error) {
	log.Println("Compiling report with special format for Investor")
	return entity.SalesReport{
		Header:    "Earning Report Oct 2020",
		SubHeader: "October 2020",
		Content:   "Sales report content for Investor",
		Footer:    "Earning Report",
		Orders:    orders,
		CreatedAt: time.Now(),
	}, nil
}

func (i *Investor) StoreReportData(report entity.SalesReport) error {
	log.Println("Storing report data to special Investor-Report's application database")
	return nil
}

func (i *Investor) SendReport(report entity.SalesReport) error {
	log.Println("Sending report to Investor Relations webpage with status Draft")
	log.Println("Notifying Investor-Relations Manager to review the report")
	return nil
}
