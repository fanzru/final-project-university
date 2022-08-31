package repo

import (
	"backend/infrastructure/config"
	"backend/infrastructure/database"
)

type Impl interface {
}
type AccountsRepo struct {
	Db  database.Connection
	Cfg config.Config
}

func New(a AccountsRepo) AccountsRepo {
	return a
}
