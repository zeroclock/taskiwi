package handler

import (
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"

	"github.com/labstack/echo"
	"taskiwi/config"
)

func InitRouting(e *echo.Echo) {
	e.GET("/", indexHandler)
	e.GET("/all", allTaskHandler)
}

type Employee struct {
	Name string
}

func indexHandler(c echo.Context) error {
	f, err := os.Open("./web/taskiwi/build/index.html")
	if err != nil {
		log.Println("[WARNING] Failed to load index.html")
	}
	defer f.Close()
	
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("[WARNING] Failed to read from file buffer")
	}
	
	return c.HTML(http.StatusOK, string(b))
}

func allTaskHandler(c echo.Context) error {
	jsonData, err := json.Marshal(config.GlobalConf.CData)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, string(jsonData))
}
