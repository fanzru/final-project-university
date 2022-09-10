package services

import (
	accountshandler "backend/app/accounts/http"
	accountsrepo "backend/app/accounts/repo"
	accountsusecase "backend/app/accounts/usecase"
	"backend/infrastructure/config"
	"backend/infrastructure/database"
)

func RegisterServiceAccounts(db database.Connection, cfg config.Config) accountshandler.AccountHandler {
	accountsDB := accountsrepo.New(accountsrepo.AccountsRepo{
		MySQL: db,
		Cfg:   cfg,
	})

	accountsApp := accountsusecase.New(accountsusecase.AccountsApp{
		AccountsRepo: accountsDB,
		Cfg:          cfg,
	})

	accountHandler := accountshandler.AccountHandler{
		AccountsApp: accountsApp,
		Cfg:         cfg,
	}

	return accountHandler
}
