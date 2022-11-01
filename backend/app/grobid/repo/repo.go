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
	GetPaperUsers(ctx echo.Context, paperId int64) (*models.PapersUsers, error)
	GetSentencesLabels(ctx echo.Context, paperId int64) (*[]models.SentencesLabel, error)
	BulkUpdateSentences(ctx echo.Context, request resp.PDFToTEI, isSubmit bool) error
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

	err := tx.Table("papers_users").Create(&papersUsers).Error
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

	err = tx.Table("sentences_labels").CreateInBatches(&sentencesLabel, 100).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (g *GrobidRepo) GetPaperUsers(ctx echo.Context, paperId int64) (*models.PapersUsers, error) {
	papersUsers := &models.PapersUsers{}
	err := g.MySQL.DB.Table("papers_users").Where("id = ? AND user_id = ?", paperId, ctx.Get("user_id")).First(papersUsers).Error
	if err != nil {
		return nil, err
	}
	return papersUsers, nil
}

func (g *GrobidRepo) GetSentencesLabels(ctx echo.Context, paperId int64) (*[]models.SentencesLabel, error) {
	sentencesLabel := &[]models.SentencesLabel{}
	err := g.MySQL.DB.Table("sentences_labels").Where("paper_id = ?", paperId).Find(sentencesLabel).Error
	if err != nil {
		return nil, err
	}
	return sentencesLabel, nil
}

func (g *GrobidRepo) BulkUpdateSentences(ctx echo.Context, request resp.PDFToTEI, isSubmit bool) error {
	tx := g.MySQL.DB.Begin()
	sentences := &[]models.SentencesLabel{}
	sentUpdates := []models.SentencesLabel{}
	err := tx.Table("sentences_labels").Where("paper_id = ?", request.PaperId).Find(&sentences).Error
	if err != nil {
		return err
	}

	for _, s := range *sentences {
		for _, hr := range request.Body {
			for _, sr := range hr.Sentences {
				if s.Id == sr.SentID {
					if s.IsImportant != sr.IsImportant {
						sentUpdates = append(sentUpdates, models.SentencesLabel{
							Id:          sr.SentID,
							IsImportant: sr.IsImportant,
						})
					}
				}
			}
		}
	}

	for _, su := range sentUpdates {
		err := tx.Table("sentences_labels").Where("id = ?", su.Id).Updates(map[string]interface{}{"is_important": su.IsImportant}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if isSubmit {
		err := tx.Table("papers_users").Where("id = ?", request.PaperId).Updates(models.PapersUsers{IsDone: true}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
