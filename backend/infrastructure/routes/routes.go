package routes

import (
	"log"
	"net/http"

	accounts "backend/app/accounts/http"
	grobid "backend/app/grobid/http"
	"backend/infrastructure/middleware"

	"github.com/labstack/echo/v4"
)

type ModuleHandler struct {
	AccountHandler accounts.AccountHandler
	GrobidHandler  grobid.GrobidHandler
	MiddlewareAuth middleware.MiddlewareAuth
}

func NewRoutes(h ModuleHandler, app *echo.Echo) *echo.Echo {

	log.Println("Starting to create new routing...")

	// test api connect or not
	app.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "FANZRU PASTI LULUS S1 INFORMATIKA 200 OK")
	})

	// accounts
	accounts := app.Group("/accounts")
	accounts.POST("/login", h.AccountHandler.Login)
	accounts.POST("/register", h.AccountHandler.Register)
	accounts.GET("/profile", h.MiddlewareAuth.BearerTokenMiddleware(h.AccountHandler.Profile))

	//grobid
	grobid := app.Group("/grobid")
	grobid.POST("/pdf-to-tei", h.MiddlewareAuth.BearerTokenMiddleware(h.GrobidHandler.PdfToTeiParse))

	return app
}
