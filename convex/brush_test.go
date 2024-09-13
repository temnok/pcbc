package convex

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/lab/bitmap"
	"temnok/lab/contour"
	"temnok/lab/geom"
	"temnok/lab/path"
	"temnok/lab/util"
	"testing"
)

func TestBrushes(t *testing.T) {
	bm := bitmap.NewBitmap(1000, 200)

	for d := 1; d < 20; d++ {
		circle := Circle(d)
		circle.IterateRowsXY(50*d, 50, bm.SetRow)

		rect := new(Shape)
		path.Iterate(
			contour.RoundRect(float64(d), float64(d)*2, 1+float64(d)/4),
			geom.Identity(),
			rect.AddPoint,
		)

		rect.IterateRowsXY(50*d, 100, bm.SetRow)
	}

	assert.NoError(t, util.SaveTmpPng("brush.png", bm.ToImage(color.Black, color.White)))
}
