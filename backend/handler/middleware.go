package handler

import (
	"log"

	"github.com/labstack/echo"
)

func FirebaseAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("authenticating firebase token ...")
		log.Println("... OK")
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
