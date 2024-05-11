package main

import (
	"tanam-backend/database"
	"tanam-backend/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	database.Migrate()

	e := echo.New()
	routes.InitRoute(e)

	e.Logger.Fatal(e.Start(":1323"))
}
