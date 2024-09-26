package contour

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/lab/bitmap"
	"temnok/lab/geom"
	"temnok/lab/path"
	"temnok/lab/shape"
	"temnok/lab/util"
	"testing"
)

func TestPie(t *testing.T) {
	bm := bitmap.NewBitmap(4000, 400)

	brush := shape.Circle(1)

	for d := 1.0; d < 10; d++ {
		c := Pie(8, 80, 120, d*2*geom.Degree)

		transform := geom.MoveXY(d*400, 200)
		path.IterateAll(c, transform, func(x, y int) {
			brush.IterateRowsXY(x, y, bm.Set1)
		})
	}

	assert.NoError(t, util.SaveTmpPng("pie.png", bm.ToImage(color.Black, color.White)))
}
