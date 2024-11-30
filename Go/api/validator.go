package api

import (
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	validator := validator.New()
	validator.RegisterValidation("date", DateFormatValidate)

	customValidator := &CustomValidator{validator: validator}

	return customValidator

}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func DateFormatValidate(fl validator.FieldLevel) bool {
	dateFormat := `^\d{1,2}\/\d{1,2}\/\d{4}$`

	re := regexp.MustCompile(dateFormat)

	return re.MatchString(fl.Field().String())
}
