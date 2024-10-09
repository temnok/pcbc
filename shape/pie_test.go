package shape

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
	"temnok/pcbc/util"
	"testing"
)

func TestPie(t *testing.T) {
	bm := bitmap.NewBitmap(400, 400)

	parts := path.Pie(6, 150, 180, 10*geom.Degree).Transform(geom.RotateD(90))

	for _, part := range parts {
		shape := FromContour(part)
		shape.IterateRowsXY(200, 200, bm.Set1)
	}

	assert.NoError(t, util.SaveTmpPng("pie.png", bm.ToImage(color.Black, color.White)))
}

func TestPiePart(t *testing.T) {
	bm := bitmap.NewBitmap(20, 20)

	part := path.PiePiece(6, 9, 60*geom.Degree).Transform(geom.RotateD(150))
	shape := FromContour(part)
	shape.IterateRowsXY(10, 10, bm.Set1)

	assert.NoError(t, util.SaveTmpPng("pie-part.png", bm.ToImage(color.Black, color.White)))
}
