package http

import (
	"backend/app/grobid/domain/param"
	"backend/app/grobid/domain/resp"
	grobidapp "backend/app/grobid/usecase"
	"backend/infrastructure/config"
	"backend/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GrobidHandler struct {
	GrobidApp grobidapp.GrobidApp
	Cfg       config.Config
}

func (a GrobidHandler) PdfToTeiParse(ctx echo.Context) error {
	pdfName := ctx.FormValue("pdf_name")
	domainPaper := ctx.FormValue("domain_paper")
	pdfFile, err := ctx.FormFile("pdf_file")
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	result, err := a.GrobidApp.PdfToTeiParse(ctx, param.GrobidUploadParam{
		DomainPaper: domainPaper,
		PdfName:     pdfName,
		PdfFile:     pdfFile,
	})

	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, result)
}

func (a GrobidHandler) GetDetailPaperById(ctx echo.Context) error {
	s := ctx.Param("id")
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	resp, err := a.GrobidApp.GetDetailPaper(ctx, id)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, resp)
}

func (a GrobidHandler) EditPaper(ctx echo.Context) error {
	s := ctx.Param("isSubmit")
	isSubmit, err := strconv.ParseBool(s)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	var req resp.PDFToTEI
	err = ctx.Bind(&req)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = a.GrobidApp.EditPaper(ctx, req, isSubmit)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, nil)
}
