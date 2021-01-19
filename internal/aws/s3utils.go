package aws

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/umtkas/image-resizer-lambda/configs"
)

func createS3Downloader(configuration configs.Configuration) *s3manager.Downloader {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(configuration.Region),
	})

	if err != nil {
		exitErrorf("cannot create s3 session, %v", err)
	}

	return s3manager.NewDownloader(sess)
}

func createS3Uploader(configuration configs.Configuration) *s3manager.Uploader {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(configuration.Region),
	})

	if err != nil {
		exitErrorf("cannot create s3 session, %v", err)
	}

	return s3manager.NewUploader(sess)
}

// DownloadFile downloads file from s3 bucket
func DownloadFile(configuration configs.Configuration, objectKey string) string {
	downloader := createS3Downloader(configuration)

	file, err := os.Create(configuration.LocalImageDirectory + "/" + filepath.Base(objectKey))

	if err != nil {
		exitErrorf("DownloadFile:::Unable to open file, %v", err)
	}

	defer file.Close()

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(configuration.Bucket),
			Key:    aws.String(objectKey),
		})

	if err != nil {
		exitErrorf("Unable to download item %q, %v", objectKey, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	return configuration.LocalImageDirectory + "/" + filepath.Base(objectKey)
}

// UploadFiles uploads resized images to s3 bucket
func UploadFiles(configuration configs.Configuration, fileNames []string) {
	for _, fileName := range fileNames {
		uploadFile(configuration, fileName)
	}
}

// uploads resized image to s3 bucket
func uploadFile(configuration configs.Configuration, filename string) {
	uploader := createS3Uploader(configuration)

	filePath := configuration.LocalImageDirectory + "/" + filename
	file, err := os.Open(filePath)

	if err != nil {
		exitErrorf("uploadFile::: Unable to open file, %v", err)
	}

	defer file.Close()

	bucketDir := configuration.UploadDirectory + "/" + filename
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(configuration.Bucket),
		Key:    aws.String(bucketDir),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, configuration.Bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, configuration.Bucket)
}
