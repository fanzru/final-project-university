package main

import (
	accountshandler "backend/app/accounts/http"
	accountsrepo "backend/app/accounts/repo"
	accountsusecase "backend/app/accounts/usecase"
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

	mHandler := registerService(db, cfg)

	e := echo.New()
	e = routes.NewRoutes(mHandler, e)
	log.Fatal(e.Start(":8888"))
}

func registerService(db database.Connection, cfg config.Config) routes.ModuleHandler {

	// accounts services
	accountsDB := accountsrepo.New(accountsrepo.AccountsRepo{
		Db:  db,
		Cfg: cfg,
	})

	accountsApp := accountsusecase.New(accountsusecase.AccountsApp{
		AccountsRepo: accountsDB,
		Cfg:          cfg,
	})

	accountHandler := accountshandler.AccountHandler{
		AccountsApp: accountsApp,
		Cfg:         cfg,
	}
	// xxxxx services
	return routes.ModuleHandler{
		AccountHandler: accountHandler,
	}
}
