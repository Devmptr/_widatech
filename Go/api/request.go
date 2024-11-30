package api

import (
	"github.com/labstack/echo/v4"
)

func RequestValidate[T any](c echo.Context, req *T) error {
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	return nil
}
