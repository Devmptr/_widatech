package services

import (
	"database/sql"
	"fmt"
	"widatech_interview/golang/model"
)

type productService struct {
	BaseService[model.Product]
	dbConn *sql.DB
}

func (s *productService) Create(data model.Product) (model.Product, error) {
	_, err := s.dbConn.Exec(
		fmt.Sprintf(
			"INSERT INTO products (invoice_no, item_name, quantity, total_cogs, total_price) VALUES ('%s', '%s', '%d', '%d', '%d')",
			data.InvoiceNo,
			data.ItemName,
			data.Quantity,
			data.TotalCogs,
			data.TotalPrice,
		),
	)

	if err != nil {
		return model.Product{}, err
	}

	return data, nil
}

func (s *productService) Read(param interface{}) interface{} {
	invoiceNo := param.(string)
	lists := []model.Product{}

	query := fmt.Sprintf(
		`SELECT * FROM products WHERE invoice_no = "%s"`,
		invoiceNo,
	)

	rows, err := s.dbConn.Query(query)
	if err == nil {
		defer rows.Close()

		for rows.Next() {
			var InvoiceNo string
			var ItemName string
			var Quantity int
			var TotalCogs int
			var TotalPrice int

			rows.Scan(&InvoiceNo, &ItemName, &Quantity, &TotalCogs, &TotalPrice)
			lists = append(lists, model.Product{
				InvoiceNo:  InvoiceNo,
				ItemName:   ItemName,
				Quantity:   Quantity,
				TotalCogs:  TotalCogs,
				TotalPrice: TotalPrice,
			})
		}

		return lists
	}

	return []model.Invoice{}
}

func (s *productService) Update(id string, data interface{}) (model.Product, error) {
	return model.Product{}, nil
}

func (s *productService) Delete(id string) (model.Product, error) {
	return model.Product{}, nil
}

func NewProductService(dbConn *sql.DB) BaseService[model.Product] {
	return &productService{
		dbConn: dbConn,
	}
}
