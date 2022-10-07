package param

import "mime/multipart"

type GrobidUploadParam struct {
	PdfName string
	PdfFile *multipart.FileHeader
}
