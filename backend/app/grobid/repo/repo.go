package repo

import (
	"backend/app/grobid/domain/models"
	"backend/app/grobid/domain/resp"
	"backend/infrastructure/config"
	"backend/infrastructure/database"

	"github.com/labstack/echo/v4"
)

type Repo interface {
	SaveUserPapersAndBulkInsertSentencesWithTx(ctx echo.Context, papersUsers models.PapersUsers, pdfTEI *resp.PDFToTEI) error
}

type GrobidRepo struct {
	MySQL database.Connection
	Cfg   config.Config
}

func New(g GrobidRepo) GrobidRepo {
	return g
}

func (g *GrobidRepo) SaveUserPapersAndBulkInsertSentencesWithTx(ctx echo.Context, papersUsers models.PapersUsers, pdfTEI *resp.PDFToTEI) error {
	tx := g.MySQL.DB.Begin()

	err := tx.Table("papers_user").Create(&papersUsers).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	sentencesLabel := []models.SentencesLabel{}
	for _, v := range pdfTEI.Body {
		for _, s := range v.Sentences {
			sentencesLabel = append(sentencesLabel, models.SentencesLabel{
				PaperId:     papersUsers.Id,
				Head:        v.Head,
				Text:        s.Text,
				IsImportant: s.IsImportant,
			})
		}
	}

	err = tx.Table("sentences_label").CreateInBatches(&sentencesLabel, 100).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
