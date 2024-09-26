package bitmap

import (
	"image"
	"image/color"
)

type bitmapsImage struct {
	bitmaps []*Bitmap
	palette color.Palette
}

func NewBitmapImage(bitmap *Bitmap, zero, one color.Color) image.Image {
	return NewBitmapsImage([]*Bitmap{bitmap}, [][2]color.Color{{zero, one}})
}

func NewBitmapsImage(bitmaps []*Bitmap, bitmapColors [][2]color.Color) image.Image {
	palette := make(color.Palette, 1<<len(bitmaps))

	for i := range palette {
		colors := make([]color.Color, len(bitmapColors))

		for j, bc := range bitmapColors {
			color := bc[0]
			if i&(1<<j) != 0 {
				color = bc[1]
			}

			colors[j] = color
		}

		palette[i] = combineColors(colors)
	}

	return &bitmapsImage{
		bitmaps: bitmaps,
		palette: palette,
	}
}

func (bi *bitmapsImage) ColorModel() color.Model {
	return bi.palette
}

func (bi *bitmapsImage) Bounds() image.Rectangle {
	b := bi.bitmaps[0]
	return image.Rect(0, 0, b.w, b.h)
}

func (bi *bitmapsImage) At(x, y int) color.Color {
	return bi.palette[bi.ColorIndexAt(x, y)]
}

func (bi *bitmapsImage) ColorIndexAt(x, y int) uint8 {
	y = bi.bitmaps[0].h - 1 - y

	index := 0

	for i, b := range bi.bitmaps {
		index |= b.Get(x, y) << i
	}

	return uint8(index)
}

func combineColors(colors []color.Color) color.Color {
	const f = 0xFFFF

	var tr, tg, tb float64

	for _, color := range colors {
		r, g, b, a := color.RGBA()
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
