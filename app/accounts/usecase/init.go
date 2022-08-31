package usecase

import (
	"backend/app/accounts/repo"
	"backend/infrastructure/config"
)

type Impl interface {
}

type AccountsApp struct {
	AccountsRepo repo.Impl
	Cfg          config.Config
}

func New(accounts AccountsApp) AccountsApp {
	return accounts
}
