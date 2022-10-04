package http

import (
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
	return response.ResponseSuccessOK(ctx, nil)
}
