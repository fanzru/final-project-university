package usecase

import (
	"backend/app/grobid/domain/models"
	"backend/app/grobid/domain/outbound"
	"backend/app/grobid/domain/param"
	"backend/app/grobid/domain/resp"
	"backend/pkg/converter"
	"backend/pkg/s3"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

const (
	URLArtu = "https://ir-group.ec.tuwien.ac.at/artu_summarize/summarize_article"
)

func (g *GrobidApp) PdfToTeiParse(ctx echo.Context, param param.GrobidUploadParam) (*resp.PDFToTEI, error) {
	src, err := param.PdfFile.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	srcFile := "../../../temp" + param.PdfFile.Filename
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

	responseResult := &outbound.TEI{}
	err = xml.Unmarshal(responseGrobid.Body(), &responseResult)
	if err != nil {
		return nil, err
	}

	result := &resp.PDFToTEI{}
	result.MapToTEIParse(responseResult)

	cloudStorage, err := s3.NewS3Object(g.Cfg.S3.Endpoint, g.Cfg.S3.AccessKeyIdS3, g.Cfg.S3.SecretAccessKeyS3, g.Cfg.S3.BucketName, true)
	if err != nil {
		return nil, err
	}

	a, err := cloudStorage.UploadFileFromPath(srcFile)
	if err != nil {
		return nil, err
	}

	userId, err := converter.CtxToInt64(ctx, "user_id")
	if err != nil {
		return nil, err
	}

	papersUsers := models.PapersUsers{
		DomainPaper: param.DomainPaper,
		UserId:      userId,
		PaperName:   param.PdfName,
		LinkPdf:     a.EndpointPath,
		IsDone:      false,
		CreatedAt:   time.Now(),
	}

	paperId, err := g.Repo.SaveUserPapersAndBulkInsertSentencesWithTx(ctx, papersUsers, result)
	if err != nil {
		return nil, err
	}

	result.PaperId = paperId
	result.LinkPdf = a.EndpointPath
	return result, nil
}

func (g *GrobidApp) PdfToTeiParseAsParagraf(ctx echo.Context, param param.GrobidUploadParam) (*resp.PDFToTEI, error) {
	src, err := param.PdfFile.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	srcFile := "../../../temp" + param.PdfFile.Filename
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
	// tuwien services
	client := resty.New()
	client.SetDisableWarn(true)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.SetTransport(tr)
	responsetuwein, err := client.R().
		SetMultipartFields(
			&resty.MultipartField{
				Param:       "pdf_article",
				FileName:    srcFile,
				ContentType: "application/pdf",
				Reader:      bytes.NewReader(fileBytes),
			},
			&resty.MultipartField{
				Param:       "paper_id",
				FileName:    "",
				ContentType: "text/plain",
				Reader:      strings.NewReader("pdf"),
			}).
		SetContentLength(true).
		Post(URLArtu)
	if err != nil {
		return nil, err
	}
	log.Println("-------------------------------- TUWIEN MID : ", responsetuwein)
	if responsetuwein.StatusCode() != http.StatusOK {
		return nil, errors.New(responsetuwein.Status())
	}
	responseResultTuwien := resp.DataPaperArtuSummary{}
	err = json.Unmarshal(responsetuwein.Body(), &responseResultTuwien)
	if err != nil {
		return nil, err
	}

	log.Println("-------------------------------- GROBID REQ")
	// grobid services fanzru
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

	responseResult := &outbound.TEI{}
	err = xml.Unmarshal(responseGrobid.Body(), &responseResult)
	if err != nil {
		return nil, err
	}

	result := &resp.PDFToTEI{}
	result.MapToTEIParseParagraft(responseResult)
	log.Println("-------------------------------- GROBID DONE")
	log.Println("-------------------------------- TUWIEN REQ")

	log.Println("-------------------------------- TUWIEN DONE")
	modelsLmjs := []models.TuwienSummaLmjm{}
	for _, summary := range responseResultTuwien.Summaries {
		if summary.Method == "lmjm" {
			for _, sentence := range summary.Summary {
				modelsLmjs = append(modelsLmjs, models.TuwienSummaLmjm{
					Id:        0,
					PaperId:   0,
					Sent:      sentence,
					CreatedAt: time.Now().UTC(),
				})
			}
		}
	}

	modelsBM25 := []models.TuwienSummaBM25{}
	for _, summary := range responseResultTuwien.Summaries {
		if summary.Method == "bm25" {
			for _, sentence := range summary.Summary {
				modelsBM25 = append(modelsBM25, models.TuwienSummaBM25{
					Id:        0,
					PaperId:   0,
					Sent:      sentence,
					CreatedAt: time.Now().UTC(),
				})
			}
		}
	}

	/*
		1. Parse grobid to database, and tuwien result [done]
		2. TX begin
		3. Insert dataset_paper
		4. Insert Bulk grobid_extraction
		5. Insert Bulk lmjm
		6. Insert Bulk bm25
		7. Tx Commit
	*/

	// cloudStorage, err := s3.NewS3Object(g.Cfg.S3.Endpoint, g.Cfg.S3.AccessKeyIdS3, g.Cfg.S3.SecretAccessKeyS3, g.Cfg.S3.BucketName, true)
	// if err != nil {
	// 	return nil, err
	// }

	// a, err := cloudStorage.UploadFileFromPath(srcFile)
	// if err != nil {
	// 	return nil, err
	// }

	// userId, err := converter.CtxToInt64(ctx, "user_id")
	// if err != nil {
	// 	return nil, err
	// }

	// papersUsers := models.PapersUsers{
	// 	DomainPaper: param.DomainPaper,
	// 	UserId:      userId,
	// 	PaperName:   param.PdfName,
	// 	LinkPdf:     a.EndpointPath,
	// 	IsDone:      false,
	// 	CreatedAt:   time.Now(),
	// }

	// paperId, err := g.Repo.SaveUserPapersAndBulkInsertSentencesWithTx(ctx, papersUsers, result)
	// if err != nil {
	// 	return nil, err
	// }

	result.Lmjm = modelsLmjs
	result.BM25 = modelsBM25
	result.PaperId = 0
	result.LinkPdf = "https://fanzru.dev/"
	return result, nil
}
