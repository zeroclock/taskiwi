package main

import (
	"taskiwi/config"
	"taskiwi/handler"
	"taskiwi/validation"

	"flag"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	var csvpath = flag.String("path", "", "csv path")
	flag.Parse()

	config.GlobalConf = config.InitConfig(*csvpath)

	v := validator.New()
	v.RegisterValidation("is_date", validation.DateValidation)

	e := echo.New()
	
	e.Validator = &validation.CustomValidator{Validator: v}
	e.Static("/", "./web/taskiwi/build")
	handler.InitRouting(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use((middleware.CORS()))
	e.Logger.Fatal(e.Start(":8080"))
}
