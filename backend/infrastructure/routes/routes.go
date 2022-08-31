package routes

import (
	"log"
	"net/http"

	accounts "backend/app/accounts/http"

	"github.com/labstack/echo/v4"
)

type ModuleHandler struct {
	AccountHandler accounts.AccountHandler
}

func NewRoutes(h ModuleHandler, app *echo.Echo) *echo.Echo {

	log.Println("Starting to create new routing...")

	// test api connect or not
	app.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "FANZRU PASTI LULUS S1 INFORMATIKA 200 OK")
	})

	return app
}
