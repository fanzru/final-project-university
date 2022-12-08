package usecase

import (
	"backend/app/grobid/domain/param"
	"backend/app/grobid/domain/resp"
	"backend/app/grobid/repo"
	"backend/infrastructure/config"

	"github.com/labstack/echo/v4"
)

type Impl interface {
	PdfToTeiParse(ctx echo.Context, Param param.GrobidUploadParam) (*resp.PDFToTEI, error)
	GetDetailPaper(ctx echo.Context, paperId int64) (*resp.PDFToTEI, error)
	EditPaper(ctx echo.Context, request resp.PDFToTEI, isSubmit bool) error
	GetSentencesLabelsCSV(ctx echo.Context, paperId int64) (*resp.CSVresp, error)
}

type GrobidApp struct {
	Repo repo.GrobidRepo
	Cfg  config.Config
}

func New(g GrobidApp) GrobidApp {
	return g
}
