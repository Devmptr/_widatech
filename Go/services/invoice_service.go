package services

import (
	"database/sql"
	"fmt"
	"widatech_interview/golang/helpers"
	"widatech_interview/golang/model"
)

type invoiceService struct {
	BaseService[model.Invoice]
	dbConn *sql.DB
}

func (s *invoiceService) Create(data model.Invoice) (model.Invoice, error) {
	notes := "NULL"
	if len(data.Notes) > 0 {
		notes = fmt.Sprintf(`"%s"`, data.Notes)
	}
	_, err := s.dbConn.Exec(
		fmt.Sprintf(
			"INSERT INTO invoices (invoice_no, date, customer_name, sales_person_name, payment_type, notes) VALUES ('%s', '%s', '%s', '%s', '%s', %s)",
			data.InvoiceNo,
			helpers.ReformatDate(data.Date),
			data.CustomerName,
			data.SalesPersonName,
			data.PaymentType,
			notes,
		),
	)

	if err != nil {
		return model.Invoice{}, err
	}

	return data, nil
}

func (s *invoiceService) Read(param interface{}) interface{} {
	data := param.(*model.InvoiceGet)
	lists := []model.Invoice{}

	query := fmt.Sprintf(
		`SELECT * FROM invoices WHERE date = "%s" LIMIT %d OFFSET %d`,
		helpers.ReformatDate(data.Date),
		data.Size,
		data.GetPageOffset(),
	)

	rows, err := s.dbConn.Query(query)
	if err == nil {
		defer rows.Close()

		for rows.Next() {
			var InvoiceNo string
			var Date string
			var CustomerName string
			var SalesPersonName string
			var PaymentType model.PaymentType
			var Notes string

			rows.Scan(&InvoiceNo, &Date, &CustomerName, &SalesPersonName, &PaymentType, &Notes)
			lists = append(lists, model.Invoice{
				InvoiceNo:       InvoiceNo,
				Date:            Date,
				CustomerName:    CustomerName,
				SalesPersonName: SalesPersonName,
				PaymentType:     PaymentType,
				Notes:           Notes,
			})
		}
	}

	return lists
}

func (s *invoiceService) Update(id string, request interface{}) (model.Invoice, error) {
	data := request.(*model.InvoiceUpdate)
	notes := "NULL"
	if len(data.Notes) > 0 {
		notes = fmt.Sprintf(`"%s"`, data.Notes)
	}

	query := fmt.Sprintf(
		`UPDATE invoices SET customer_name = "%s", date = "%s", sales_person_name = "%s", payment_type = "%s", notes = %s WHERE invoice_no = "%s"`,
		data.CustomerName,
		helpers.ReformatDate(data.Date),
		data.SalesPersonName,
		data.PaymentType,
		notes,
		data.InvoiceNo,
	)
	_, err := s.dbConn.Exec(query)
	if err != nil {
		return model.Invoice{}, err
	}

	return model.Invoice{
		InvoiceNo:       data.InvoiceNo,
		CustomerName:    data.CustomerName,
		SalesPersonName: data.SalesPersonName,
		Date:            data.Date,
		PaymentType:     data.PaymentType,
		Notes:           data.Notes,
	}, nil
}

func (s *invoiceService) Delete(id string) (model.Invoice, error) {
	query := fmt.Sprintf(
		`DELETE FROM invoices WHERE invoice_no = "%s"`,
		id,
	)
	_, err := s.dbConn.Exec(query)
	if err != nil {
		return model.Invoice{}, err
	}

	return model.Invoice{}, nil
}

func NewInvoiceService(dbConn *sql.DB) BaseService[model.Invoice] {
	return &invoiceService{
		dbConn: dbConn,
	}
}
