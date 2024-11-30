package model

import (
	"widatech_interview/golang/helpers"
)

type PaymentType string

const (
	PaymentType_CASH   PaymentType = "CASH"
	PaymentType_CREDIT PaymentType = "CREDIT"
)

type Invoice struct {
	InvoiceNo       string      `json:"invoice_no" validate:"required,min=1"`
	Date            string      `json:"date" validate:"required,date"`
	CustomerName    string      `json:"customer_name" validate:"required,min=2"`
	SalesPersonName string      `json:"sales_person_name" validate:"required,min=2"`
	PaymentType     PaymentType `json:"payment_type" validate:"required,oneof=CASH CREDIT"`
	Notes           string      `json:"notes" validate:"omitempty,min=5"`
	ListOfProduct   []Product   `json:"list_of_product" validate:"omitempty,min=1"`
}

func (i *Invoice) CreateFromExcel(row []string) {
	i.InvoiceNo = row[0]
	i.Date = helpers.ReformatDateExcel(row[1])
	i.CustomerName = row[2]
	i.SalesPersonName = row[3]
	i.PaymentType = PaymentType(row[4])
	i.Notes = row[5]
	// i.ListOfProduct = []Product{}
}

type InvoiceGet struct {
	Date string `json:"date" validate:"required,date"`
	Size int    `json:"size" validate:"required,numeric,min=1"`
	Page int    `json:"page" validate:"required,numeric,min=1"`
}

type InvoiceUpdate struct {
	InvoiceNo       string      `json:"invoice_no" validate:"required,min=1"`
	Date            string      `json:"date" validate:"required,date"`
	CustomerName    string      `json:"customer_name" validate:"required,min=2"`
	SalesPersonName string      `json:"sales_person_name" validate:"required,min=2"`
	PaymentType     PaymentType `json:"payment_type" validate:"required,oneof=CASH CREDIT"`
	Notes           string      `json:"notes" validate:"omitempty,min=5"`
}

func (i *InvoiceGet) GetPageOffset() int {
	return (i.Page - 1) * i.Size
}

type InvoiceSummary struct {
	TotalProfit          int
	TotalCashTransaction int
	Invoices             []Invoice
}
