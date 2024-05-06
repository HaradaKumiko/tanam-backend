package main

import (
	"fmt"
	"net/http"
	"tanam-backend/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = func(err error, c echo.Context) {
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
					report.Message = fmt.Sprintf("%s is not valid email",
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param())
				}

			}
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}
	routes.InitRoute(e)

	// e.GET("/", func(c echo.Context) error {

	// 	token, _ := middlewares.GenerateTokenJWT("randomUd", "email@gmail.com")
	// 	fmt.Println("this is value: ", token)
	// 	return c.String(http.StatusOK, "hello world")
	// })
	e.Logger.Fatal(e.Start(":1323"))
}
