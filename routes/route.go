package routes

import (
	"tanam-backend/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	authController := controllers.InitAuthController()

	api := e.Group("/api")
	auth := api.Group("/auth")
	auth.POST("/login", authController.LoginController)
}
