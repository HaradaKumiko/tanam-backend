package routes

import (
	"tanam-backend/controllers"
	"tanam-backend/controllers/admin"
	"tanam-backend/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()

	e.Use(loggerMiddleware)

	authController := controllers.InitAuthController()
	adminPlantController := admin.InitPlantController()
	plantController := controllers.InitPlantController()

	api := e.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/login", authController.LoginController)
	auth.POST("/register", authController.RegisterController)
	auth.GET("/profile", authController.ProfileController, middlewares.JWTMiddleware)

	plant := api.Group("")
	plant.GET("/plants", plantController.GetAllPlantController)
	plant.GET("/plant/:plantId", plantController.GetPlantByPlantIdController)

	admin := api.Group("/admin")
	admin.GET("/plants", adminPlantController.GetAllPlantController)
	admin.POST("/plant", adminPlantController.CreatePlantController)
	admin.GET("/plant/:plantId", adminPlantController.GetPlantByPlantIdController)
	admin.PUT("/plant/:plantId", adminPlantController.EditPlantByPlantIdController)
	admin.DELETE("/plant/:plantId", adminPlantController.DeletePlantByIdController)
}
