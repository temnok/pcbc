package pkg

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/lab/bitmap"
	"temnok/lab/contour"
	"temnok/lab/convex"
	"temnok/lab/geom"
	"temnok/lab/path"
	"temnok/lab/util"
	"testing"
)

func TestPadRow_SavePNG(t *testing.T) {
	bm := bitmap.NewBitmap(1000, 1000)
	transform := geom.Move(geom.XY{500, 500}).RotateD(45).ScaleK(100)

	convex.IterateContoursRows(QFN16PadContours, transform, bm.SetRow)

	brush := convex.Circle(3)
	path.Iterate(contour.Rect(3, 3), transform, func(x, y int) {
		brush.IterateRowsXY(x, y, bm.SetRow)
	})

	brush = convex.Circle(20)
	path.Iterate([]geom.XY{{-3, 3}, {-3, 3}, {3, -3}, {3, -3}}, transform, func(x, y int) {
		brush.IterateRowsXY(x, y, bm.SetRow)
	})

	assert.NoError(t, util.SaveTmpPng("footprint.png", bm.ToImage(color.Black, color.White)))
}
