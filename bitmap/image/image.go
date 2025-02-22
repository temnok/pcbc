// Copyright © 2025 Alex Temnok. All rights reserved.

package image

import (
	"image"
	"image/color"
	"temnok/pcbc/bitmap"
)

type bitmapsImage struct {
	bitmaps []*bitmap.Bitmap
	palette color.Palette
	flipY   bool
}

func NewSingle(bm *bitmap.Bitmap, zero, one color.Color) image.Image {
	return New([]*bitmap.Bitmap{bm}, [][2]color.Color{{zero, one}}, false)
}

func New(bitmaps []*bitmap.Bitmap, bitmapColors [][2]color.Color, flipY bool) image.Image {
	palette := make(color.Palette, 1<<len(bitmaps))

	for i := range palette {
		colors := make([]color.Color, len(bitmapColors))

		for j, bc := range bitmapColors {
			c := bc[0]
			if i&(1<<j) != 0 {
				c = bc[1]
			}

			colors[j] = c
		}

		palette[i] = combineColors(colors)
	}

	return &bitmapsImage{
		bitmaps: bitmaps,
		palette: palette,
		flipY:   flipY,
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

func (bi *bitmapsImage) ColorIndexAt(x, y int) uint8 {
	if bi.flipY {
		y = bi.bitmaps[0].Height() - 1 - y
	}

	index := 0

	for i, b := range bi.bitmaps {
		index |= b.Get(x, y) << i
	}

	return uint8(index)
}

func combineColors(colors []color.Color) color.Color {
	if len(colors) == 1 {
		return colors[0]
	}

	const f = 0xFFFF

	var tr, tg, tb float64

	for _, c := range colors {
		r, g, b, a := c.RGBA()
		k := float64(a) / f
		tr = tr*(1-k) + (float64(r)/f)*k
		tg = tg*(1-k) + (float64(g)/f)*k
		tb = tb*(1-k) + (float64(b)/f)*k
	}

	return color.RGBA64{
		R: uint16(tr * f),
		G: uint16(tg * f),
		B: uint16(tb * f),
		A: f,
	}
}
