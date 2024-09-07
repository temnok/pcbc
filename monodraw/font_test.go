package monodraw

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"image/png"
	"os"
	"temnok/lab/bezier"
	"testing"
)

func TestFont_SavePng(t *testing.T) {
	b := NewBitmap(16*130, 16*200)
	brush := NewRoundBrush(24) // light:16, norm:24, bold:32

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			c := i*16 + j

			if c >= len(font) || len(font[c]) == 0 {
				continue
			}

			x0, y0 := float64(j*130), float64(i*200)

			for _, stroke := range font[c] {
				var px, py float64

				for _, p := range stroke {
					x := x0 + float64(10+(p%10)*20)
					y := y0 + float64(10+(p/10)*20)

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
