package configs

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// ImageSize stores image size
type ImageSize struct {
	WidthHeight string
	Width       int
	Height      int
}

// Configuration stores environment variables, and s3 details
type Configuration struct {
	LocalImageDirectory   string
	UploadDirectory       string
	ImageExtension        string
	IsSaveWithAspectRatio bool
	ImageSizes            []ImageSize
	Region                string
	Bucket                string
}

// GetConfiguration reads env variables and returns Configuration storate
func GetConfiguration() (*Configuration, error) {

	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")
	envImagesSizes := os.Getenv("IMAGE_SIZES")
	saveWithAspectRatioStr := os.Getenv("SAVE_WITH_ASPECT_RATIO")
	imageExtension := os.Getenv("IMAGE_EXTENSION")
	isSaveWithAspectRatio, _ := strconv.ParseBool(saveWithAspectRatioStr)
	uploadDirectory := os.Getenv("UPLOAD_DIRECTORY")

	imageSizes, err := parseImageSizes(envImagesSizes)

	if err != nil {
		return nil, err
	}

	configuration := Configuration{
		UploadDirectory:       uploadDirectory,
		ImageExtension:        imageExtension,
		IsSaveWithAspectRatio: isSaveWithAspectRatio,
		ImageSizes:            imageSizes,
		Region:                region,
		Bucket:                bucket,
		LocalImageDirectory:   "/tmp",
	}

	return &configuration, nil
}

func parseImageSizes(envImageSizes string) ([]ImageSize, error) {

	var imageSizes []ImageSize

	if len(envImageSizes) == 0 {
		return imageSizes, nil
	}

	arr := strings.Split(envImageSizes, ",")

	for _, envImageSize := range arr {
		sizeArr := strings.Split(envImageSize, "x")
		height, err := strconv.Atoi(sizeArr[0])
		if err != nil {
			return nil, errors.New("height is not found")
		}

		width, err := strconv.Atoi(sizeArr[1])
		if err != nil {
			return nil, errors.New("width is not found")
		}

		imageSize := ImageSize{
			Height:      height,
			Width:       width,
			WidthHeight: envImageSize,
		}

		imageSizes = append(imageSizes, imageSize)
	}

	return imageSizes, nil

}
