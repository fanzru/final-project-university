package usecase

import (
	"backend/app/grobid/repo"
	"backend/infrastructure/config"
)

type GrobidApp struct {
	GrobidRepo repo.GrobidRepo
	Cfg        config.Config
}

func New(accounts GrobidApp) GrobidApp {
	return accounts
}
