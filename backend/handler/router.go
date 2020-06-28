package handler

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
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
	datastoreClient, _ := datastore.NewClient(ctx, config.DSConfig.ProjectID)
	e := &Employee{
		Name: "tarou tanaka",
	}
	k := datastore.IncompleteKey("Employee", nil)
	_, err := datastoreClient.Put(ctx, k, e)
	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, "Hello")
}
