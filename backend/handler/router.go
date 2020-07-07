package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"taskiwi/config"
)

func InitRouting(e *echo.Echo) {
	e.GET("/", helloHandler)
}

type Employee struct {
	Name string
}

func helloHandler(c echo.Context) error {
	log.Println("hello action")

	ctx := context.Background()
	config := config.GetConfig()
	e := &Employee{
		Name: "tarou tanaka",
	}

	return c.String(http.StatusOK, "Hello")
}
