package handler

import (
	"log"
	"net/http"
	"os"
	"io/ioutil"

	"github.com/labstack/echo"
	"taskiwi/config"
	"taskiwi/utils"
)

func InitRouting(e *echo.Echo) {
	e.GET("/", indexHandler)
	e.GET("/all", allTaskHandler)
	e.GET("/allTags", allTagsHandler)
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
	return c.JSON(http.StatusOK, config.GlobalConf.CData)
}

func allTagsHandler(c echo.Context) error {
	var tags []string
	for _, v := range *config.GlobalConf.CData {
		tags = append(tags, v.Tags...)
	}
	
	return c.JSON(http.StatusOK, utils.Unique(tags))
}
