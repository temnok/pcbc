package monodraw

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"image/png"
	"os"
	"temnok/lab/bezier"
	"temnok/lab/font"
	"testing"
)

func TestFont_SavePng(t *testing.T) {
	const scale = 200.0

	b := NewBitmap(16*scale*font.Width, 6*scale)
	brush := NewRoundBrush(font.Normal * scale)

	fontData := font.Lines

	for i := 0; i < 14; i++ {
		for j := 0; j < 16; j++ {
			c := (i+2)*16 + j

			if fontData[c] == nil {
				continue
			}

			x0, y0 := float64(j)*scale*font.Width, float64(i)*scale

			for _, stroke := range fontData[c] {
				var px, py float64

				for _, p := range stroke {
					x := x0 + p.X*scale
					y := y0 + p.Y*scale

					if px != 0 {
						bezier.CubicVisit([]bezier.Point{{px, py}, {px, py}, {x, y}, {x, y}}, func(x, y int) {
							b.Segments(x, y, brush)
						})
					}
					px, py = x, y
				}
			}
		}
	}

	f, err := os.Create("font.png")
	assert.NoError(t, err)

	im := b.ToImage(color.White, color.Black)
	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())

}
