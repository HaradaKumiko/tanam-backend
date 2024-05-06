package constant

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var ErrInsertDatabase error = errors.New("invalid Add Data in Database")
var ErrInvalidRequest error = errors.New("invalid Request")
var ErrEmptyInput error = errors.New("input cannot be empty")
var ErrDuplicatedData error = errors.New("duplicated data")
var ErrNotFound error = errors.New("not found")
var ErrInvalidEmailOrPassword error = errors.New("invalid email or password")
var ErrNotAuthorized error = errors.New("not authorized")
var ErrInternalServer error = errors.New("internal server error")
var ErrJobNotOpened error = errors.New("job not opened")
var ErrJobAlreadyFull error = errors.New("job already full")
var ErrFailedUpdate error = errors.New("failed to update the data")
var ErrHelperAlreadyTakeTheJob error = errors.New("helper already take the job")

func CustomHTTPErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not a valid email",
					err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s value must be greater than or equal to %s",
					err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s value must be less than or equal to %s",
					err.Field(), err.Param())
			}

			break
		}
	}

	c.Logger().Error(report)
	c.JSON(report.Code, report)
}
