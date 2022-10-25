package s3

import (
	"context"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Object struct {
	client     *minio.Client
	endpoint   string
	bucketName string
	useSSL     bool
	ctx        context.Context
}

type S3ObjectI interface {
	GetObjectPresigned(objectPath string) (string, error)
	UploadFileMultipart(file multipart.File) (objectOutput, error)
	UploadFileFromPath(filePath string) (objectOutput, error)
	UploadFileFromPathNamed(fileName, filePath string) (objectOutput, error)
	DeleteObject(objectPath string) error
	ListObjectParentDir() []string
	GetObjectPath(fullPathEndpoint string) string
}

type objectOutput struct {
	EndpointPath string
	Path         string
	Endpoint     string
}

// Create new instace of S3 Object with MiniIO APIs
// This also will create a "./temp" folder for uploading from memory file (multipart)
func NewS3Object(endpoint, accessKeyID, secretAcessKey, bucketName string, useSSL bool) (S3ObjectI, error) {
	ctx := context.Background()

	err := os.Mkdir("./temp", 0644)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return nil, err
	}

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAcessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return &S3Object{}, err
	}

	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return &S3Object{}, err
	}

	if !exists {
		return &S3Object{}, errors.New("bucket does not exist")
	}

	return &S3Object{
		client:     client,
		endpoint:   endpoint,
		bucketName: bucketName,
		useSSL:     useSSL,
		ctx:        ctx,
	}, nil
}

// Generate temporary object URL for fetching, will expire in 1 minute
func (s *S3Object) GetObjectPresigned(objectPath string) (string, error) {
	objectPath = s.GetObjectPath(objectPath)
	url, err := s.client.PresignedGetObject(s.ctx, s.bucketName, objectPath, time.Second*60, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

// Upload file using local file instances. This will generate random path to the file.
func (s *S3Object) UploadFileMultipart(file multipart.File) (objectOutput, error) {
	data, err := s.multiPartToByte(file)
	if err != nil {
		return objectOutput{}, err
	}
	tempFile, err := s.createTempFile(data)
	if err != nil {
		return objectOutput{}, err
	}

	object, err := s.UploadFileFromPath(tempFile)
	if err != nil {
		return objectOutput{}, err
	}
	err = s.deleteTempFile(tempFile)
	if err != nil {
		return objectOutput{}, err
	}
	return object, nil
}

// Upload file using local file paths. This will generate random path to the file.
func (s *S3Object) UploadFileFromPath(filePath string) (objectOutput, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return objectOutput{}, err
	}
	objectPathName, err := s.generateObjectPathName(data)
	if err != nil {
		return objectOutput{}, err
	}

	_, err = s.client.FPutObject(s.ctx, s.bucketName, objectPathName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return objectOutput{}, err
	}

	return s.createObjectOutput(objectPathName), nil
}

// Upload file using local file paths. This will generate random path to the file.
func (s *S3Object) UploadFileFromPathNamed(fileName, filePath string) (objectOutput, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return objectOutput{}, err
	}
	ext, err := s.getFileExtension(data)
	if err != nil {
		return objectOutput{}, err
	}
	fileName = fileName + ext
	_, err = s.client.FPutObject(s.ctx, s.bucketName, fileName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return objectOutput{}, err
	}

	return s.createObjectOutput(fileName), nil
}

// Convert from "https://example.com/bucket/file.jpg" to "file.jpg"
// or from "bucket/file.jpg" to "file.jpg"
func (s *S3Object) GetObjectPath(fullPathEndpoint string) string {
	targetString := s.bucketName + "/"
	idx := strings.Index(fullPathEndpoint, targetString)
	if idx == -1 {
		return fullPathEndpoint
	}
	return fullPathEndpoint[idx+len(targetString):]
}

func (s *S3Object) ListObjectParentDir() []string {
	var obj []string
	objectCh := s.client.ListObjects(s.ctx, s.bucketName, minio.ListObjectsOptions{})
	for object := range objectCh {
		if object.Err == nil {
			obj = append(obj, object.Key)
		}
	}
	return obj
}

// Delete based on object path or full path
func (s *S3Object) DeleteObject(objectPath string) error {
	objectPath = s.GetObjectPath(objectPath)
	return s.client.RemoveObject(s.ctx, s.bucketName, objectPath, minio.RemoveObjectOptions{
		ForceDelete: true,
	})
}
