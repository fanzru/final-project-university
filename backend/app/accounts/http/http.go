package http

import (
	accountsapp "backend/app/accounts/usecase"
	"backend/infrastructure/config"
)

type AccountHandler struct {
	AccountsApp accountsapp.Impl
	Cfg         config.Config
}
