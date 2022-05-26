package utils

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	Sess *session.Session
}

var S3 *S3Client

func init() {
	S3 = new(S3Client)
}

func (t *S3Client) NewSession() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			""), // token can be left blank for now
	})

	if err != nil {
		log.Println("Failed to create new session", err)
	}
	t.Sess = sess
}

// UploadFileToS3 saves a file to aws bucket and returns the url to // the file and an error if there's any
func (t *S3Client) UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader, id string) (string, error) {
	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	// create a unique file name for the file
	tempFileName := "pictures/" + strings.Split(fileHeader.Filename, ".")[0] + id + filepath.Ext(fileHeader.Filename)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(t.Sess).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(os.Getenv("AWS_BUCKET")),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	return tempFileName, err
}

// GenerateUrl generates a url to the file in s3
func (t *S3Client) GenerateUrl(keyName string) (string, error) {
	req, _ := s3.New(t.Sess).GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(keyName),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
		return "", err
	}
	return urlStr, nil
}
