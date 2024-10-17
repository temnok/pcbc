package font

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/geom"
	"temnok/pcbc/shape"
	"temnok/pcbc/util"
	"testing"
)

func TestFont_SavePng(t *testing.T) {
	const scale = 200.0

	bm := bitmap.NewBitmap(16*scale*Width, 6*scale)
	brush := shape.Circle(Normal * scale)

	for i := 0; i < 14; i++ {
		for j := 0; j < 16; j++ {
			c := (i+2)*16 + j

			transform := geom.ScaleXY(scale, -scale).MoveXY(float64(j)*Width, float64(-i))
			brush.IterateContours(Paths[c].Transform(transform), bm.Set1)
		}
	}

	assert.NoError(t, util.SavePng("gen/font.png", bm.ToImage(color.Black, color.White)))
}
