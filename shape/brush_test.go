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

func TestBrushes(t *testing.T) {
	bm := bitmap.New(1000, 300)

	for d := 1; d < 20; d++ {
		circle := Circle(d)
		circle.IterateRowsXY(50*d, 50, bm.Set1)

		circle = FromContour(transform.I, path.Circle(float64(d)))
		circle.IterateRowsXY(50*d, 100, bm.Set1)

		rect := FromContour(transform.I, path.RoundRect(float64(d), float64(d)*2, 1+float64(d)/4))
		rect.IterateRowsXY(50*d, 200, bm.Set1)
	}

	assert.NoError(t, util.SavePNG("tmp/brush.png", image.NewSingle(bm, color.Black, color.White)))
}
