package main

import (
	"taskiwi/handler"
	"taskiwi/config"
	"taskiwi/validation"

	"flag"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/go-playground/validator/v10"
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
	log.Println(config.GlobalConf.IData.FileContent)

	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	e.Static("/", "./web/taskiwi/build")
	handler.InitRouting(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":8080"))
}
