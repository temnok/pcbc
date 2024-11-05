package font

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/shape"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
	"testing"
)

func TestFont_SavePng(t *testing.T) {
	const scale = 100.0

	bm := bitmap.NewBitmap(16*scale*Width, 20*scale)

	normalBrush := shape.Circle(Normal * scale)

	d := Bold * scale
	boldBrush := shape.Circle(int(d))

	extraBoldBrush := shape.Circle(ExtraBold * scale)

	for i := 0; i < 14; i++ {
		for j := 0; j < 16; j++ {
			c := (i+2)*16 + j

			tx := transform.Move(float64(j)*Width, float64(-i)).Scale(scale, -scale)
			normalBrush.IterateContours(Paths[c].Apply(tx), bm.Set1)

			tx = transform.Move(float64(j)*Width, float64(-i-7)).Scale(scale, -scale)
			boldBrush.IterateContours(Paths[c].Apply(tx), bm.Set1)

			tx = transform.Move(float64(j)*Width, float64(-i-14)).Scale(scale, -scale)
			extraBoldBrush.IterateContours(Paths[c].Apply(tx), bm.Set1)
		}
	}

	assert.NoError(t, util.SavePng("out/font.png", bm.ToImage(color.Black, color.White)))
}
