package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := NewRouter()
	e.Logger.Fatal(e.Start(":8080"))
}

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/hello", helloHandler)

	return e
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}
