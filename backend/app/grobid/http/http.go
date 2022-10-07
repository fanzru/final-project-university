package http

import (
	"backend/app/grobid/domain/param"
	grobidapp "backend/app/grobid/usecase"
	"backend/infrastructure/config"
	"backend/pkg/response"

	"github.com/labstack/echo/v4"
)

type GrobidHandler struct {
	GrobidApp grobidapp.GrobidApp
	Cfg       config.Config
}

func (a GrobidHandler) PdfToTeiParse(ctx echo.Context) error {
	pdfName := ctx.FormValue("pdf_name")
	pdfFile, err := ctx.FormFile("pdf_file")
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = a.GrobidApp.PdfToTeiParse(ctx, param.GrobidUploadParam{
		PdfName: pdfName,
		PdfFile: pdfFile,
	})

	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, nil)
}
