package pkg

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/lab/bitmap"
	"temnok/lab/contour"
	"temnok/lab/eda/lib/pkg/qfn16"
	"temnok/lab/geom"
	"temnok/lab/shape"
	"temnok/lab/util"
	"testing"
)

func TestPadRow_SavePNG(t *testing.T) {
	bm := bitmap.NewBitmap(1000, 1000)
	transform := geom.MoveXY(500, 500).RotateD(45).ScaleK(100)

	shape.IterateContoursRows(qfn16.PadContours, transform, bm.SetRow1)
	shape.Circle(3).IterateContour(contour.Rect(3, 3), transform, bm.SetRow1)
	shape.Circle(20).IterateContour(contour.Lines([]geom.XY{{-3, 3}, {3, -3}}), transform, bm.SetRow1)

	assert.NoError(t, util.SaveTmpPng("footprint.png", bm.ToImage(color.Black, color.White)))
}
