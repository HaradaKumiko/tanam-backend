package google

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var upload *ClientUploader

const (
	ProjectID  = "octopuslab-365307" // Fill in with your project ID
	BucketName = "payroll_anggi"     // Fill in with your bucket name
)

const MaxFileSize = 5 * 1024 * 1024 // 5 MB as an example limit

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func init() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("key.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	upload = &ClientUploader{
		cl:         client,
		bucketName: BucketName,
		projectID:  ProjectID,
		uploadPath: "test/",
	}
	log.Println("init called")
}

func getFileSize(file multipart.File) (int64, error) {
	// Seek to the end of the file to get its size
	size, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}
	// Seek back to the beginning of the file
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(ctx context.Context, file multipart.File, object string) (string, error) {
	fileSize, err := getFileSize(file)
	if err != nil {
		return "", fmt.Errorf("failed to get file size: %v", err)
	}
	if fileSize > MaxFileSize {
		return "", fmt.Errorf("file size exceeds the limit: %v", MaxFileSize)
	}
	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}
	return "", nil
}
