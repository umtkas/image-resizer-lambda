package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/umtkas/image-resizer-lambda/internal/resizer"

	"github.com/aws/aws-lambda-go/events"
	"github.com/umtkas/image-resizer-lambda/configs"
)

// Handler handles aws lambda action
func Handler(ctx context.Context, s3Event events.S3Event) {
	configuration, err := configs.GetConfiguration()

	if err != nil {
		exitErrorf("configuration is not valid %v", err)
	}

	for _, record := range s3Event.Records {
		objectKey := record.S3.Object.Key
		sourceImage := DownloadFile(*configuration, objectKey)
		resizedImages := resizer.GetResizedImages(*configuration, sourceImage)
		UploadFiles(*configuration, resizedImages)
	}

	resizer.RemoveImages(*configuration)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
