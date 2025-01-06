// Copyright © 2025 Alex Temnok. All rights reserved.

package shape

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/path"
	"temnok/pcbc/util"
	"testing"
)

func TestBrushes(t *testing.T) {
	bm := bitmap.NewBitmap(1000, 300)

	for d := 1; d < 20; d++ {
		circle := Circle(d)
		circle.IterateRowsXY(50*d, 50, bm.Set1)

		circle = FromContour(path.Circle(float64(d))[0])
		circle.IterateRowsXY(50*d, 100, bm.Set1)

		rect := FromContour(path.RoundRect(float64(d), float64(d)*2, 1+float64(d)/4)[0])
		rect.IterateRowsXY(50*d, 200, bm.Set1)
	}

	assert.NoError(t, util.SavePNG("tmp/brush.png", bm.ToImage(color.Black, color.White)))
}
