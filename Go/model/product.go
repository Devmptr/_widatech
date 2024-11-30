package model

type Product struct {
	InvoiceNo  string `json:"invoice_no" validate:"omitempty"`
	ItemName   string `json:"item_name" validate:"required,min=5"`
	Quantity   int    `json:"quantity" validate:"required,numeric,min=1"`
	TotalCogs  int    `json:"total_cogs" validate:"required,numeric,min=0"`
	TotalPrice int    `json:"total_price" validate:"required,numeric,min=0"`
}
