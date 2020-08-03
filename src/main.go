package main

import (
	"taskiwi/config"
	"taskiwi/handler"
	"taskiwi/validation"

	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var pidFile = flag.String("pid-file", "", "./")
	flag.Parse()

	if len(*pidFile) > 0 {
		if err := ioutil.WriteFile(*pidFile, []byte(fmt.Sprintf("%d", os.Getpid())), 0664); err != nil {
			log.Printf("[WARNING] Failed to write pid file. %v\n", err)
		}
		defer func() {
			if err := os.Remove(*pidFile); err != nil {
				log.Printf("[WARNING] Failed to delete pid file. %v\n", err)
			}
		}()
	}

	path := "./test.csv"
	config.GlobalConf = config.InitConfig(path)

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
