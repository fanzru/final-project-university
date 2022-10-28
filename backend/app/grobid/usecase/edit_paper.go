package usecase

import (
	"backend/app/grobid/domain/resp"

	"github.com/labstack/echo/v4"
)

func (g *GrobidApp) EditPaper(ctx echo.Context, request resp.PDFToTEI, isSubmit bool) error {
	err := g.Repo.BulkUpdateSentences(ctx, request, isSubmit)
	if err != nil {
		return err
	}
	return nil
}
