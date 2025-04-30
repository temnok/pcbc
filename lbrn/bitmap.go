// Copyright © 2025 Alex Temnok. All rights reserved.

package lbrn

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
)

type Base64Bitmap struct {
	W, H int
	Data string
}

func NewBase64Bitmap(im image.Image) *Base64Bitmap {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, im); err != nil {
		panic(err)
	}

	return &Base64Bitmap{
		W:    im.Bounds().Dx(),
		H:    im.Bounds().Dy(),
		Data: base64.StdEncoding.EncodeToString(buf.Bytes()),
	}
}
