package main

import (
	"backend/cmd/services"
	"backend/infrastructure/config"
	"backend/infrastructure/database"
	"backend/infrastructure/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Start Services....")

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	accountsHandler := services.RegisterService(db, cfg)

	mHandler := routes.ModuleHandler{
		AccountHandler: accountsHandler,
	}

	e := echo.New()
	e = routes.NewRoutes(mHandler, e)
	log.Fatal(e.Start(":8888"))
}
