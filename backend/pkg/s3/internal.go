package s3

import (
	"crypto/rand"
	"errors"
	"io/ioutil"
	"math/big"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func (s *S3Object) createObjectOutput(objectPath string) objectOutput {
	fullPath := "http"
	if s.useSSL {
		fullPath = fullPath + "s"
	}
	fullPath = fullPath + "://" + s.endpoint + "/" + s.bucketName + "/" + objectPath
	return objectOutput{
		EndpointPath: fullPath,
		Path:         objectPath,
		Endpoint:     s.endpoint,
	}
}

// Generate Random Object Path Name
func (s *S3Object) generateObjectPathName(data []byte) (string, error) {
	fileName, err := s.generateFileName(data)
	if err != nil {
		return "", err
	}
	// path, err := s.generatePath()
	// if err != nil {
	// 	return "", err
	// }
	return "riset/" + fileName, nil
}

// Create Temporary Data File
func (s *S3Object) createTempFile(data []byte) (string, error) {
	objectPathName, err := s.generateFileName(data)
	if err != nil {
		return "", err
	}
	tempFile := "./temp/" + objectPathName
	err = ioutil.WriteFile(tempFile, data, 0644)
	if err != nil {
		return "", err
	}
	return tempFile, nil
}

// Deleting Temp File
func (s *S3Object) deleteTempFile(tempFile string) error {
	// NOTE: Probably will cause bug if there's a concurent connection on uploading
	return os.RemoveAll(tempFile)
}

// Generate Random Path
func (s *S3Object) generatePath() (string, error) {
	finalPath := ""
	for i := 0; i < 4; i++ {
		path, err := s.randomString(15)
		if err != nil {
			return "", err
		}
		finalPath = finalPath + path + "/"
	}
	return finalPath, nil
}

// Generate UUID File Name
func (s *S3Object) generateFileName(data []byte) (string, error) {
	fileExtension, err := s.getFileExtension(data)
	if err != nil {
		return "", err
	}
	return uuid.New().String() + fileExtension, nil
}

// Detect Content Type outputing mime
func (s *S3Object) detectContentType(data []byte) string {
	return http.DetectContentType(data)
}

// Get registered file extension by bytes
func (s *S3Object) getFileExtension(data []byte) (string, error) {
	mimeType := s.detectContentType(data)
	fileExtension := ""
	switch mimeType {
	case "image/jpeg":
		fileExtension = fileExtension + ".jpg"
	case "image/png":
		fileExtension = fileExtension + ".png"
	case "application/pdf":
		fileExtension = fileExtension + ".pdf"
	case "video/mp4":
		fileExtension = fileExtension + ".mp4"
	case "application/zip":
		fileExtension = fileExtension + ".zip"
	case "application/octet-stream":
		return "", errors.New("unsupported file type")
	}
	return fileExtension, nil
}

// Convert multipart byte to one part byte
func (s *S3Object) multiPartToByte(file multipart.File) ([]byte, error) {
	var finalByte []byte
	b := make([]byte, 100)
	for {
		n, err := file.Read(b)
		finalByte = append(finalByte, b...)
		if n == 0 {
			break
		}
		if err != nil {
			return []byte{}, err
		}
	}
	return finalByte, nil
}

// Generate random string
func (s *S3Object) randomString(n int) (string, error) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		b[i] = letterBytes[idx.Int64()]
	}
	return string(b), nil
}
