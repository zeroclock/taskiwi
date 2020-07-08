package main

import (
	"taskiwi/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"flag"
	"log"
	"fmt"
	"os"
	"io/ioutil"
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

	e := echo.New()
	e.Static("/", "./web/taskiwi/build")
	handler.InitRouting(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(handler.FirebaseAuthMiddleware)
	e.Logger.Fatal(e.Start(":8080"))
}
