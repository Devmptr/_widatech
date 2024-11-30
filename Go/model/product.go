package model

import (
	"strconv"
)

type Product struct {
	InvoiceNo  string `json:"invoice_no" validate:"omitempty"`
	ItemName   string `json:"item_name" validate:"required,min=5"`
	Quantity   int    `json:"quantity" validate:"required,numeric,min=1"`
	TotalCogs  int    `json:"total_cogs" validate:"required,numeric,min=0"`
	TotalPrice int    `json:"total_price" validate:"required,numeric,min=0"`
}

func (p *Product) CreateFromExcel(row []string) error {
	p.InvoiceNo = row[0]
	p.ItemName = row[1]

	if num, err := strconv.Atoi(row[2]); err != nil {
		return err
	} else {
		p.Quantity = num
	}

	if num, err := strconv.Atoi(row[3]); err != nil {
		return err
	} else {
		p.TotalCogs = num
	}

	if num, err := strconv.Atoi(row[4]); err != nil {
		return err
	} else {
		p.TotalPrice = num
	}

	return nil
}
