// Copyright © 2025 Alex Temnok. All rights reserved.

package util

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"os"
	"testing"
)

type imageMock struct {
	colorModel color.Model
	bounds     image.Rectangle
	at         func(x, y int) color.Color
}

func (im *imageMock) ColorModel() color.Model {
	return im.colorModel
}

func (im *imageMock) Bounds() image.Rectangle {
	return im.bounds
}

func (im *imageMock) At(x, y int) color.Color {
	return im.at(x, y)
}

func TestSavePNG(t *testing.T) {
	tests := []struct {
		name          string
		fileName      string
		image         image.Image
		expectedError string
	}{
		{
			name:     "success",
			fileName: "success/image.png",
			image: &imageMock{
				colorModel: color.RGBAModel,
				bounds:     image.Rect(0, 0, 4, 4),
				at: func(x, y int) color.Color {
					return color.White
				},
			},
		},
	}

	tempDir := os.TempDir() + "/" + uuid.NewString()
	assert.NoError(t, os.Mkdir(tempDir, 0700))
	defer func() {
		assert.NoError(t, os.RemoveAll(tempDir))
	}()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			errText := ""
			err := SavePNG(tempDir+"/"+test.fileName, test.image)
			if err != nil {
				errText = err.Error()
			}

			assert.Equal(t, test.expectedError, errText)
		})
	}
}
