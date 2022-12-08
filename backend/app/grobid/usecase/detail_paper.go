package usecase

import (
	"backend/app/grobid/domain/resp"

	"github.com/labstack/echo/v4"
)

func (g *GrobidApp) GetDetailPaper(ctx echo.Context, paperId int64) (*resp.PDFToTEI, error) {
	papersUsers, err := g.Repo.GetPaperUsers(ctx, paperId)
	if err != nil {
		return nil, err
	}
	sentencesLabel, err := g.Repo.GetSentencesLabels(ctx, papersUsers.Id)
	if err != nil {
		return nil, err
	}
	response := &resp.PDFToTEI{}
	response.MapToResponse(papersUsers, sentencesLabel)
	return response, nil
}

func (g *GrobidApp) GetSentencesLabelsCSV(ctx echo.Context, paperId int64) (*resp.CSVresp, error) {
	papersUsers, err := g.Repo.GetPaperUsers(ctx, paperId)
	if err != nil {
		return nil, err
	}
	sentencesLabel, err := g.Repo.GetSentencesLabels(ctx, papersUsers.Id)
	if err != nil {
		return nil, err
	}

	return &resp.CSVresp{PaperDetail: *papersUsers, SentencesLabel: *sentencesLabel}, nil
}
