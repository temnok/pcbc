package font

import (
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"image/png"
	"os"
	"temnok/lab/bitmap"
	"temnok/lab/geom"
	"temnok/lab/path"
	"testing"
)

func TestFont_SavePng(t *testing.T) {
	const scale = 200.0

	b := bitmap.NewBitmap(16*scale*Width, 6*scale)
	brush := bitmap.NewRoundBrush(Normal * scale)

	for i := 0; i < 14; i++ {
		for j := 0; j < 16; j++ {
			c := (i+2)*16 + j

			transform := geom.ScaleLocked(scale).Move(geom.XY{float64(j) * Width, float64(i)})
			path.IterateAll(Paths[c], transform, func(x, y int) {
				b.Segments(x, y, brush)
			})
		}
	}

	savePng(t, "font.png", b.ToImage(color.White, color.Black))
}

func savePng(t *testing.T, name string, im image.Image) {
	_ = os.Mkdir("tmp", 0770)

	f, err := os.Create("tmp/" + name)
	assert.NoError(t, err)

	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())
}
