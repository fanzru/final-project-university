package repo

import (
	"backend/infrastructure/config"
	"backend/infrastructure/database"
)

type GrobidRepo struct {
	MySQL database.Connection
	Cfg   config.Config
}

func New(g GrobidRepo) GrobidRepo {
	return g
}
