package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"widatech_interview/golang/config"
	"widatech_interview/golang/model"
	"widatech_interview/golang/services"

	"github.com/labstack/echo/v4"
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
			c.JSON(http.StatusBadRequest, BaseResponse{
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
}

func Start(e *echo.Echo, serverConfig *config.ServerConfig) {
	e.Logger.Fatal(
		e.Start(fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port)),
	)
}
