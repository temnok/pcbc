package font

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/lab/bitmap"
	"temnok/lab/geom"
	"temnok/lab/shape"
	"temnok/lab/util"
	"testing"
)

func TestFont_SavePng(t *testing.T) {
	const scale = 200.0

	bm := bitmap.NewBitmap(16*scale*Width, 6*scale)
	brush := shape.Circle(Normal * scale)

	for i := 0; i < 14; i++ {
		for j := 0; j < 16; j++ {
			c := (i+2)*16 + j

			transform := geom.ScaleK(scale).Move(geom.XY{float64(j) * Width, float64(6 - i)})
			brush.IterateContours(Paths[c], transform, bm.SetRow1)
		}
	}

	assert.NoError(t, util.SaveTmpPng("font.png", bm.ToImage(color.White, color.Black)))
}
