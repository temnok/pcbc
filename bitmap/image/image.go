// Copyright © 2025 Alex Temnok. All rights reserved.

package image

import (
	"github.com/temnok/pcbc/bitmap"
	"image"
	"image/color"
	"math"
)

type bitmapsImage struct {
	bitmaps []*bitmap.Bitmap
	palette color.Palette
}

func NewSingle(bm *bitmap.Bitmap, zero, one color.Color) image.Image {
	return New([]*bitmap.Bitmap{bm}, [][2]color.Color{{zero, one}})
}

func New(bitmaps []*bitmap.Bitmap, bitmapColors [][2]color.Color) image.Image {
	return &bitmapsImage{
		bitmaps: bitmaps,
		palette: createPalette(bitmaps, bitmapColors),
	}
}

func (bi *bitmapsImage) ColorModel() color.Model {
	return bi.palette
}

func (bi *bitmapsImage) Bounds() image.Rectangle {
	b := bi.bitmaps[0]
	return image.Rect(0, 0, b.Width(), b.Height())
}

func (bi *bitmapsImage) At(x, y int) color.Color {
	return bi.palette[bi.ColorIndexAt(x, y)]
}

func (bi *bitmapsImage) ColorIndexAt(x, y int) byte {
	y = bi.bitmaps[0].Height() - 1 - y

	index := 0
	for i, b := range bi.bitmaps {
		index |= b.Get(x, y) << i
	}

	return byte(index)
}

func createPalette(bitmaps []*bitmap.Bitmap, bitmapColors [][2]color.Color) color.Palette {
	palette := make(color.Palette, 1<<len(bitmaps))

	for i := range palette {
		colors := make([]color.Color, len(bitmapColors))
		for j, bc := range bitmapColors {
			colors[j] = bc[(i>>j)&1]
		}

		palette[i] = mixColors(colors)
	}

	return palette
}

func mixColors(colors []color.Color) color.Color {
	if len(colors) == 1 {
		return colors[0]
	}

	var r, g, b float64
	for _, c := range colors {
		r, g, b = mixInColor(r, g, b, c)
	}

	return color.RGBA64{
		R: uint16(r * math.MaxUint16),
		G: uint16(g * math.MaxUint16),
		B: uint16(b * math.MaxUint16),
		A: math.MaxUint16,
	}
}

func mixInColor(tr, tg, tb float64, c color.Color) (float64, float64, float64) {
	r, g, b, a := c.RGBA()
	k := float64(a) / math.MaxUint16
	tr = tr*(1-k) + (float64(r)/math.MaxUint16)*k
	tg = tg*(1-k) + (float64(g)/math.MaxUint16)*k
	tb = tb*(1-k) + (float64(b)/math.MaxUint16)*k
	return tr, tg, tb
}
