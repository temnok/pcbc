package shape

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/lab/bitmap"
	"temnok/lab/geom"
	"temnok/lab/path"
	"temnok/lab/util"
	"testing"
)

func TestPie(t *testing.T) {
	bm := bitmap.NewBitmap(400, 400)

	parts := path.Pie(6, 150, 180, 10*geom.Degree)

	for _, part := range parts {
		shape := FromContour(part)
		shape.IterateRowsXY(200, 200, bm.Set1)
	}

	assert.NoError(t, util.SaveTmpPng("pie.png", bm.ToImage(color.White, color.Black)))
}
