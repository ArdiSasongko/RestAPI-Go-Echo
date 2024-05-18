package helper

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateBind(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)

	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if CastedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range CastedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s must filled", err.Field())
				report.Code = http.StatusBadRequest
			case "email":
				report.Message = fmt.Sprintf("%s invalid", err.Field())
				report.Code = http.StatusBadRequest
			}
		}
	}

	c.Logger().Error(report.Message)
	c.JSON(report.Code, ResponseClient(report.Code, report.Message.(string), nil))
}
