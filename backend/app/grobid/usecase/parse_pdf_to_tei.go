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
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
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

	sentencesLabel := []models.SentencesLabel{}
	for _, v := range result.Body {
		for _, s := range v.Sentences {
			sentencesLabel = append(sentencesLabel, models.SentencesLabel{
				PaperId:     0,
				Head:        v.Head,
				Text:        s.Text,
				IsImportant: s.IsImportant,
			})
		}
	}

	// @TODO : Save sentencesLabel to MySQL and save paper to database
	log.Println(sentencesLabel)

	/*
		Transaction
		1. Upload PDF to S3
		2. Create Data to table papers_user
		3. Create Bulk Data to table sentences_label
	*/

	// @TODO : upload to S3
	cloudStorage, err := s3.NewS3Object(g.Cfg.S3.Endpoint, g.Cfg.S3.AccessKeyIdS3, g.Cfg.S3.SecretAccessKeyS3, g.Cfg.S3.BucketName, true)
	if err != nil {
		return nil, err
	}

	a, err := cloudStorage.UploadFileFromPath(srcFile)
	if err != nil {
		return nil, err
	}

	log.Println("--------------------------------", a.EndpointPath)
	userId, err := converter.CtxToInt64(ctx, "user_id")
	if err != nil {
		return nil, err
	}
	log.Println("--------------------------------", userId)
	papersUsers := models.PapersUsers{
		UserId:    userId,
		PaperName: param.PdfName,
		LinkPdf:   a.EndpointPath,
		IsDone:    false,
		CreatedAt: time.Now(),
	}

	log.Println(papersUsers)
	err = g.Repo.SaveUserPapersAndBulkInsertSentencesWithTx(ctx, papersUsers, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
