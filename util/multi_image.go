package util

import (
	"image"
	"image/color"
)

type MultiImage struct {
	Images []image.Image
}

func (mi *MultiImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (mi *MultiImage) Bounds() image.Rectangle {
	return mi.Images[0].Bounds()
}

func (mi *MultiImage) At(x, y int) color.Color {
	var tr, tg, tb float64

	for _, i := range mi.Images {
		r, g, b, a := i.At(x, y).RGBA()
		k := float64(a) / 0xFFFF
		tr += (float64(r) / 0xFFFF) * k
		tg += (float64(g) / 0xFFFF) * k
		tb += (float64(b) / 0xFFFF) * k
	}

	return color.RGBA{byte(tr * 0xFF), byte(tg * 0xFF), byte(tb * 0xFF), 0xFF}
}
