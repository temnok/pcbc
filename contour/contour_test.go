package contour

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/lab/bitmap"
	"temnok/lab/convex"
	"temnok/lab/geom"
	"temnok/lab/path"
	"temnok/lab/util"
	"testing"
)

func TestContours(t *testing.T) {
	bm := bitmap.NewBitmap(1000, 100)

	brush := convex.Circle(3)

	for d := 5.0; d < 24; d++ {
		cnt := Circle(d)

		transform := geom.Move(geom.XY{(d - 4) * 50, 50})
		path.Iterate(cnt, transform, func(x, y int) {
			brush.IterateRows(x, y, bm.Row)
		})
	}

	assert.NoError(t, util.SaveTmpPng("contour.png", bm.ToImage(color.Black, color.White)))
}
