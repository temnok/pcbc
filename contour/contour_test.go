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

func TestContours(t *testing.T) {
	bm := bitmap.NewBitmap(1000, 100)

	brush := shape.Circle(3)

	for d := 5.0; d < 24; d++ {
		cnt := Circle(d)

		transform := geom.MoveXY((d-4)*50, 50)
		path.Path(cnt).Transform(transform).Visit(func(x, y int) {
			brush.IterateRowsXY(x, y, bm.Set1)
		})
	}

	assert.NoError(t, util.SaveTmpPng("contour.png", bm.ToImage(color.Black, color.White)))
}
