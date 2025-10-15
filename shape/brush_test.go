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

func TestBrushes(t *testing.T) {
	bm := bitmap.New(1000, 300)

	for d := 1; d < 20; d++ {
		circle := Circle(d)
		circle.ForEachRowWithOffset(50*d, 50, bm.Set1)

		circle = New(path.Circle(float64(d)), transform.I)
		circle.ForEachRowWithOffset(50*d, 100, bm.Set1)

		rect := New(path.RoundRect(float64(d), float64(d)*2, 1+float64(d)/4), transform.I)
		rect.ForEachRowWithOffset(50*d, 200, bm.Set1)
	}

	assert.NoError(t, util.SavePNG("tmp/brush.png", image.NewSingle(bm, color.Black, color.White)))
}
