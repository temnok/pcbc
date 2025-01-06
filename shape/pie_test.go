// Copyright © 2025 Alex Temnok. All rights reserved.

package shape

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
	"testing"
)

func TestPie(t *testing.T) {
	bm := bitmap.NewBitmap(400, 400)

	parts := path.Pie(6, 150, 180, 10).Apply(transform.Rotate(90))

	for _, part := range parts {
		shape := FromContour(part)
		shape.IterateRowsXY(200, 200, bm.Set1)
	}

	assert.NoError(t, util.SavePNG("tmp/pie.png", bm.ToImage(color.Black, color.White)))
}

func TestPiePart(t *testing.T) {
	bm := bitmap.NewBitmap(20, 20)

	part := path.PiePiece(6, 9, 60).Apply(transform.Rotate(150))
	shape := FromContour(part)
	shape.IterateRowsXY(10, 10, bm.Set1)

	assert.NoError(t, util.SavePNG("tmp/pie-part.png", bm.ToImage(color.Black, color.White)))
}
