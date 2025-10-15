// Copyright © 2025 Alex Temnok. All rights reserved.

package shape

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/pcbc/bitmap"
	"github.com/temnok/pcbc/bitmap/image"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
	"github.com/temnok/pcbc/util"
	"image/color"
	"testing"
)

func TestPie(t *testing.T) {
	bm := bitmap.New(400, 400)

	parts := path.Pie(6, 150, 180, 10)

	for _, part := range parts {
		shape := New(part, transform.RotateDegrees(90))
		shape.ForEachRowWithOffset(200, 200, bm.Set1)
	}

	assert.NoError(t, util.SavePNG("tmp/pie.png", image.NewSingle(bm, color.Black, color.White)))
}

func TestPiePart(t *testing.T) {
	bm := bitmap.New(20, 20)

	part := path.PiePiece(6, 9, 60)
	shape := New(part, transform.RotateDegrees(150))
	shape.ForEachRowWithOffset(10, 10, bm.Set1)

	assert.NoError(t, util.SavePNG("tmp/pie-part.png", image.NewSingle(bm, color.Black, color.White)))
}
