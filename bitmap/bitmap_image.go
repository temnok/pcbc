package bitmap

import (
	"image"
	"image/color"
)

type bitmapImage struct {
	bitmap  *Bitmap
	palette color.Palette
}

func (bi *bitmapImage) ColorModel() color.Model {
	return bi.palette
}

func (bi *bitmapImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, bi.bitmap.w, bi.bitmap.h)
}

func (bi *bitmapImage) At(x, y int) color.Color {
	return bi.palette[bi.ColorIndexAt(x, y)]
}

func (bi *bitmapImage) ColorIndexAt(x, y int) uint8 {
	b := bi.bitmap
	y = b.h - 1 - y

	if x < 0 || x >= b.w || y < 0 || y >= b.h || b.elems[b.addr(x, y)]&(1<<(x%64)) == 0 {
		return 0
	}

	return 1
}
