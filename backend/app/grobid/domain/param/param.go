package param

import "mime/multipart"

type GrobidUploadParam struct {
	DomainPaper string
	PdfName     string
	PdfFile     *multipart.FileHeader
}

type GrobidUploadWithSummaParam struct {
	DomainPaper string
	PdfName     string
	PdfFile     *multipart.FileHeader
	TuwienSumma string
}
