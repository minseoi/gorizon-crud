package main

import (
	"github.com/minseoi/gorizon/db"
	"github.com/minseoi/gorizon/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db.Initialize()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
