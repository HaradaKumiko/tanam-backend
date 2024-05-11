package routes

import (
	"tanam-backend/controllers"
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
	plantController := controllers.InitPlantController()

	api := e.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/login", authController.LoginController)

	plant := api.Group("")
	plant.POST("/upload", plantController.UploadFileController)
	plant.GET("/plants", plantController.GetAllPlantController)
	plant.POST("/plant", plantController.CreatePlantController)
	plant.GET("/plant/:plantId", plantController.GetPlantByPlantIdController)
	plant.PUT("/plant/:plantId", plantController.EditPlantByPlantIdController)
	plant.DELETE("/plant/:plantId", plantController.DeletePlantByIdController)
}
