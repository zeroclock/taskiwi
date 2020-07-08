package handler

import (
	"log"
	"net/http"
	"os"
	"io/ioutil"

	"github.com/labstack/echo"
	//	"taskiwi/config"
)

func InitRouting(e *echo.Echo) {
	e.GET("/", helloHandler)
}

type Employee struct {
	Name string
}

func helloHandler(c echo.Context) error {
	log.Println("hello action")

	f, err := os.Open("./web/taskiwi/build/index.html")
	if err != nil {
		log.Println("[WARNING] Failed to load index.bhtml")
	}
	defer f.Close()
	
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("[WARNING] Failed to read from file buffer")
	}
	
	// return c.String(http.StatusOK, "Hello!??!!?!?!?!?!?!?")
	return c.HTML(http.StatusOK, string(b))
}
