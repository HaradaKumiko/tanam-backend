package controllers

import (
	"net/http"
	"tanam-backend/domains/web/auth"
	"tanam-backend/services"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	service services.AuthService
}

func InitAuthController() AuthController {
	return AuthController{
		service: services.InitAuthService(),
	}
}

func (controller *AuthController) LoginController(c echo.Context) error {
	var request auth.LoginRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	controller.service.LoginService()

	return c.JSON(http.StatusOK, "success")

}
