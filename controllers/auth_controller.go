package controllers

import (
	"net/http"
	"tanam-backend/domains/web/auth"
	"tanam-backend/helpers/response"
	"tanam-backend/services"

	_ "github.com/golang-jwt/jwt"
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

func (controller *AuthController) ProfileController(c echo.Context) error {
	authId := c.Get("authId").(string)

	auth, err := controller.service.ProfileService(authId)
	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := response.SuccessSingularFormatter("Data Profil", auth)
	return c.JSON(http.StatusOK, response)
}

func (controller *AuthController) LoginController(c echo.Context) error {
	var request auth.LoginRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	auth, err := controller.service.LoginService(request)
	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := response.SuccessSingularFormatter("Berhasil Login", auth)
	return c.JSON(http.StatusOK, response)
}

func (controller *AuthController) RegisterController(c echo.Context) error {
	var request auth.RegisterRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	auth, err := controller.service.RegisterService(request)
	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := response.SuccessSingularFormatter("Berhasil Registrasi", auth)
	return c.JSON(http.StatusOK, response)

}
