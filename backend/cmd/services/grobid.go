package services

import (
	grobidhandler "backend/app/grobid/http"
	grobidrepo "backend/app/grobid/repo"
	grobidusecase "backend/app/grobid/usecase"
	"backend/infrastructure/config"
	"backend/infrastructure/database"
)

func RegisterServiceGrobid(db database.Connection, cfg config.Config) grobidhandler.GrobidHandler {
	grobidDB := grobidrepo.New(grobidrepo.GrobidRepo{
		MySQL: db,
		Cfg:   cfg,
	})

	grobidApp := grobidusecase.New(grobidusecase.GrobidApp{
		GrobidRepo: grobidDB,
		Cfg:        cfg,
	})

	grobidHandler := grobidhandler.GrobidHandler{
		GrobidApp: grobidApp,
		Cfg:       cfg,
	}

	return grobidHandler
}
