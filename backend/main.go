package main

import (
	"taskiwi/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	handler.InitRouting(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(handler.FirebaseAuthMiddleware)
	e.Logger.Fatal(e.Start(":8080"))
}
