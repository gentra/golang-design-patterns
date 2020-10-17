package salesreportactivity

import (
	"log"
	"time"

	"github.com/gentra/golang-design-patterns/templatemethod/internal/entity"
)

type BusinessDev struct {
}

func NewBusinessDev() *BusinessDev {
	return &BusinessDev{}
}

func (b *BusinessDev) FetchOrderData() ([]entity.Order, error) {
	log.Println("Fetching order-data from application's database")
	return []entity.Order{{ID: "1"}}, nil
}

func (b *BusinessDev) CompileReport(orders []entity.Order) (entity.SalesReport, error) {
	log.Println("Compiling report with special format for Business Dev Department")
	return entity.SalesReport{
		Header:    "Sales Report Oct 2020 - BizDev Dept",
		SubHeader: "October 2020",
		Content:   "Sales report content for business development department",
		Footer:    "BizDev Dept",
		Orders:    orders,
		CreatedAt: time.Now(),
	}, nil
}

func (b *BusinessDev) StoreReportData(report entity.SalesReport) error {
	log.Println("Storing report data to BizDev department database")
	log.Println("Storing report data as Excel on BizDev's file storage")
	return nil
}

func (b *BusinessDev) SendReport(report entity.SalesReport) error {
	log.Println("Sending report to Business Development mailing group")
	log.Println("Sending report to BizDev's Head WhatsApp number")
	return nil
}
