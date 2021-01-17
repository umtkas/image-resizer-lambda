package configs

import (
	"fmt"
	"testing"
)

func Test_parseImageSizes_WithEmptyInput(t *testing.T) {
	input := ""
	output, err := parseImageSizes(input)

	if err != nil {
		t.Error(err)
	}

	if len(output) != 0 {
		t.Error("empty input is not parsed")
	}
}

func Test_parseImageSizes_WithParsableText(t *testing.T) {
	input := "200x200,300x300"
	output, err := parseImageSizes(input)

	if err != nil {
		t.Error(err)
	}

	if len(output) != 2 {
		t.Error("empty input is not parsed")
	}
}

func Test_parseImageSizes_WithNotParsableX(t *testing.T) {
	input := "200x200,300v300"
	_, err := parseImageSizes(input)

	if err == nil {
		t.Error(err)
	}
}

func Test_parseImageSizes_WithNotParsableComma(t *testing.T) {
	input := "200x200-300x300"
	_, err := parseImageSizes(input)

	if err == nil {
		t.Error(err)
	}
}

func Test_parseImageSizes_WithParameters(t *testing.T) {
	expectedImageSizes := []ImageSize{ImageSize{Height: 200, Width: 220}, ImageSize{Height: 330, Width: 300}}
	input := "200x220,330x300"
	output, err := parseImageSizes(input)

	if err != nil {
		t.Error(err)
	}

	for index, imageSize := range output {

		expectedImageSize := expectedImageSizes[index]

		if imageSize.Width != expectedImageSize.Width || imageSize.Height != expectedImageSize.Height {
			t.Error(fmt.Printf("expected and actual image sizes are not mathched --- expected imageSize: height=%d, width: %d ----- actual: height: %d, width: %d", expectedImageSize.Height, expectedImageSize.Width, imageSize.Height, imageSize.Width))
			return
		}
	}
}
