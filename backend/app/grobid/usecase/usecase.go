package usecase

import (
	"backend/app/grobid/domain/outbound"
	"backend/app/grobid/domain/param"
	"backend/app/grobid/domain/resp"
	"backend/app/grobid/repo"
	"backend/infrastructure/config"
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

type Impl interface {
	PdfToTeiParse(ctx echo.Context, Param param.GrobidUploadParam) (*resp.PDFToTEI, error)
}

type GrobidApp struct {
	GrobidRepo repo.GrobidRepo
	Cfg        config.Config
}

func New(g GrobidApp) GrobidApp {
	return g
}

func (g *GrobidApp) PdfToTeiParse(ctx echo.Context, Param param.GrobidUploadParam) (*resp.PDFToTEI, error) {

	src, err := Param.PdfFile.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	srcFile := "../../../temp" + Param.PdfFile.Filename
	dst, err := os.Create(srcFile)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	defer os.Remove(srcFile)

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	// Read file
	fileBytes, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return nil, err
	}

	client := resty.New()
	client.SetDisableWarn(true)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client.SetTransport(tr)
	responseGrobid, err := client.R().
		SetMultipartFields(
			&resty.MultipartField{
				Param:       "input",
				FileName:    srcFile,
				ContentType: "application/pdf",
				Reader:      bytes.NewReader(fileBytes),
			}).
		SetContentLength(true).
		Post(g.Cfg.Grobid.GrobidUrlPdfToTei)
	if err != nil {
		return nil, err
	}

	// @TODO : Parsing response from grobid to golang and save to DB Mysql
	responseResult := &outbound.TEI{}
	err = xml.Unmarshal(responseGrobid.Body(), &responseResult)
	if err != nil {
		return nil, err
	}

	result := &resp.PDFToTEI{}
	result.MapToTEIParse(responseResult)

	return result, nil
}
