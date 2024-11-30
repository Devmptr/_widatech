package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"widatech_interview/golang/config"
	"widatech_interview/golang/model"
	"widatech_interview/golang/services"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

func Create() *echo.Echo {
	return echo.New()
}

func Register(e *echo.Echo, dbConn *sql.DB) {
	RegisterValidator(e)
	RegisterRoute(e, dbConn)
}

func RegisterValidator(e *echo.Echo) {
	e.Validator = NewCustomValidator()
}

func RegisterRoute(e *echo.Echo, dbConn *sql.DB) {
	invoiceService := services.NewInvoiceService(dbConn)
	productService := services.NewProductService(dbConn)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/create", func(c echo.Context) error {
		req := &model.Invoice{}

		if err := RequestValidate(c, req); err != nil {
			return c.JSON(http.StatusBadRequest, BaseResponse{
				Message: err.Error(),
			})
		}

		res, err := invoiceService.Create(*req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, BaseResponse{
				Message: err.Error(),
			})
		}

		for i := 0; i < len(res.ListOfProduct); i++ {
			product := res.ListOfProduct[i]
			product.InvoiceNo = res.InvoiceNo

			_, err := productService.Create(product)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, BaseResponse{
					Message: err.Error(),
				})
			}
		}

		return c.JSON(http.StatusOK, BaseResponse{
			Message: "success",
			Data:    res,
		})
	})

	e.GET("/read", func(c echo.Context) error {
		req := &model.InvoiceGet{}

		if err := RequestValidate(c, req); err != nil {
			return c.JSON(http.StatusBadRequest, BaseResponse{
				Message: err.Error(),
			})
		}

		result := model.InvoiceSummary{
			TotalProfit:          0,
			TotalCashTransaction: 0,
		}
		invoices := invoiceService.Read(req).([]model.Invoice)

		for _, inv := range invoices {
			inv.ListOfProduct = productService.Read(inv.InvoiceNo).([]model.Product)

			for _, prd := range inv.ListOfProduct {
				result.TotalCashTransaction += (prd.TotalPrice * prd.Quantity)
				result.TotalProfit += (prd.TotalPrice * prd.Quantity) - (prd.TotalCogs * prd.Quantity)
			}

			result.Invoices = append(result.Invoices, inv)
		}

		return c.JSON(http.StatusOK, BaseResponse{
			Message: "success",
			Data:    result,
		})
	})

	e.POST("/update", func(c echo.Context) error {
		req := &model.InvoiceUpdate{}

		if err := RequestValidate(c, req); err != nil {
			return c.JSON(http.StatusBadRequest, BaseResponse{
				Message: err.Error(),
			})
		}

		res, err := invoiceService.Update(req.InvoiceNo, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, BaseResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, BaseResponse{
			Message: "success",
			Data:    res,
		})
	})

	e.DELETE("/delete", func(c echo.Context) error {
		req := struct {
			InvoiceNo string `json:"invoice_no" validate:"required,min=1"`
		}{}

		if err := RequestValidate(c, &req); err != nil {
			return c.JSON(http.StatusBadRequest, BaseResponse{
				Message: err.Error(),
			})
		}

		res, err := invoiceService.Delete(req.InvoiceNo)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, BaseResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, BaseResponse{
			Message: "success",
			Data:    res,
		})
	})

	e.POST("/import", func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		defer src.Close()

		xlsx, err := excelize.OpenReader(src)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		invoiceSheet := xlsx.GetSheetName(0)
		invoiceRows, _ := xlsx.GetRows(invoiceSheet)

		invoiceLists := []model.Invoice{}
		failed := map[string]string{}

		for i := 1; i < len(invoiceRows); i++ {
			row := invoiceRows[i]

			inv := model.Invoice{}
			inv.CreateFromExcel(row)

			if err := RequestValidate[model.Invoice](c, &inv); err != nil {
				failed[inv.InvoiceNo] = err.Error()
			} else {
				invoiceLists = append(invoiceLists, inv)
			}
		}

		productSheet := xlsx.GetSheetName(1)
		productRows, _ := xlsx.GetRows(productSheet)

		productLists := map[string][]model.Product{}

		for i := 1; i < len(productRows); i++ {
			row := productRows[i]

			prd := model.Product{}
			if err := prd.CreateFromExcel(row); err != nil {
				failed[prd.InvoiceNo] = err.Error()
			}

			exists := false
			for _, value := range invoiceLists {
				if value.InvoiceNo == prd.InvoiceNo {
					exists = true
					break
				}
			}
			if !exists {
				failed[prd.InvoiceNo] = "Invoice Not Found"
			}

			if err := RequestValidate[model.Product](c, &prd); err != nil {
				failed[prd.InvoiceNo] = err.Error()
			} else {
				if _, ok := productLists[prd.InvoiceNo]; !ok {
					productLists[prd.InvoiceNo] = []model.Product{}
				}

				productLists[prd.InvoiceNo] = append(productLists[prd.InvoiceNo], prd)
			}
		}

		listProcessed := []model.Invoice{}

		for _, invoiceToProcess := range invoiceLists {
			res, err := invoiceService.Create(invoiceToProcess)
			if err != nil {
				failed[invoiceToProcess.InvoiceNo] = err.Error()
			}
			res.ListOfProduct = productLists[invoiceToProcess.InvoiceNo]

			for i := 0; i < len(res.ListOfProduct); i++ {
				product := res.ListOfProduct[i]
				product.InvoiceNo = res.InvoiceNo

				_, err := productService.Create(product)
				if err != nil {
					failed[invoiceToProcess.InvoiceNo] = err.Error()
				}
			}

			if err == nil {
				listProcessed = append(listProcessed, res)
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":      "File uploaded and processed successfully",
			"failed":       failed,
			"invoiceLists": listProcessed,
		})
	})
}

func Start(e *echo.Echo, serverConfig *config.ServerConfig) {
	e.Logger.Fatal(
		e.Start(fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port)),
	)
}
