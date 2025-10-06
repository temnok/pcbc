// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/util"
)

func saveOverview(config *config.Config, fileName string, copper, mask, silk, stencil *bitmap.Bitmap) error {
	fileName = config.SavePath + fileName

	bitmaps := []*bitmap.Bitmap{
		copper,
		mask,
		silk,
	}

	colors := [][2]color.Color{
		{color.RGBA{R: 0xC0, G: 0x60, A: 0xFF}, color.RGBA{G: 0x40, B: 0x10, A: 0xFF}},
		{color.RGBA{}, color.RGBA{R: 0x80, G: 0x80, B: 0xFF, A: 0x80}},
		{color.RGBA{}, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xA0}},
	}

	if stencil != nil {
		bitmaps = append(bitmaps, stencil)
		colors = append(colors, [2]color.Color{color.RGBA{}, color.RGBA{R: 0x80, G: 0xFF, B: 0x80, A: 0x80}})
	}

	im := image.New(
		bitmaps,
		colors,
	)
	if err := util.SavePNG(fileName, im); err != nil {
		return err
	}

	return nil
}
