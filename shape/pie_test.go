// Copyright © 2025 Alex Temnok. All rights reserved.

package shape

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
	"testing"
)

func TestPie(t *testing.T) {
	bm := bitmap.New(400, 400)

	parts := path.Pie(6, 150, 180, 10)

	for _, part := range parts {
		shape := FromContour(transform.Rotate(90), part)
		shape.IterateRowsXY(200, 200, bm.Set1)
	}

	assert.NoError(t, util.SavePNG("tmp/pie.png", image.NewSingle(bm, color.Black, color.White)))
}

func TestPiePart(t *testing.T) {
	bm := bitmap.New(20, 20)

	part := path.PiePiece(6, 9, 60)
	shape := FromContour(transform.Rotate(150), part)
	shape.IterateRowsXY(10, 10, bm.Set1)

	assert.NoError(t, util.SavePNG("tmp/pie-part.png", image.NewSingle(bm, color.Black, color.White)))
}
