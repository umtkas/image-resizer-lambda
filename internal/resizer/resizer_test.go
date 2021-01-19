package resizer

import (
	"testing"

	"github.com/umtkas/image-resizer-lambda/configs"
)

func Test_getResizedImageName(t *testing.T) {
	imageSize := configs.ImageSize{Height: 200, Width: 200, WidthHeight: "200x200"}
	sourceImagePath := "../../tmp/image.png"
	configuration := configs.Configuration{
		ImageExtension:      "jpeg",
		LocalImageDirectory: "tmp",
	}

	expected := "image_" + imageSize.WidthHeight + "." + configuration.ImageExtension
	actual := getResizedImageName(configuration, imageSize, sourceImagePath)

	if expected != actual {
		t.Fatalf("actual image name is not same as resized image name \nactual: %s\nexpected: %s", actual, expected)
	}
}

func Test_resize(t *testing.T) {
	imageSize := configs.ImageSize{Height: 200, Width: 200, WidthHeight: "200x200"}
	sourceImagePath := "../../tmp/testimage.jpg"
	configuration := configs.Configuration{
		ImageExtension:      "png",
		LocalImageDirectory: "/tmp",
	}

	expected := "testimage_" + imageSize.WidthHeight + "." + configuration.ImageExtension
	actual := resize(configuration, imageSize, sourceImagePath)

	if expected != actual {
		t.Fatalf("actual image name is not same as resized image name \nactual: %s\nexpected: %s", actual, expected)
	}
}

func Test_GetResizedImages(t *testing.T) {
	expectedImageSizes := []configs.ImageSize{configs.ImageSize{Height: 200, Width: 200, WidthHeight: "200x200"}, configs.ImageSize{Height: 300, Width: 300, WidthHeight: "300x300"}}
	sourceImagePath := "../../tmp/testimage.jpg"

	configuration := configs.Configuration{
		IsSaveWithAspectRatio: true,
		ImageExtension:        "png",
		ImageSizes:            expectedImageSizes,
		LocalImageDirectory:   "/tmp",
	}

	var expectedImageNames []string

	for _, imageSize := range expectedImageSizes {
		expected := "testimage_" + imageSize.WidthHeight + "." + configuration.ImageExtension
		expectedImageNames = append(expectedImageNames, expected)
	}

	actualImageNames := GetResizedImages(configuration, sourceImagePath)

	for index, actualImageName := range actualImageNames {
		expectedImageName := expectedImageNames[index]
		if actualImageName != expectedImageName {
			t.Fatalf("actual image name is not same as resized image name \nactual: %s\nexpected: %s", actualImageName, expectedImageName)
		}
	}
}
