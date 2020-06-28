package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"log"
)

func InitRouting(e *echo.Echo) {
	e.GET("/", helloHandler)
}

func helloHandler(c echo.Context) error {
	log.Println("hello action")
	return c.String(http.StatusOK, "Hello")
}
