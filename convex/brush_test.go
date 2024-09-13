package convex

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"temnok/lab/bitmap"
	"temnok/lab/util"
	"testing"
)

func TestBrushes(t *testing.T) {
	bm := bitmap.NewBitmap(1000, 100)

	for d := 1; d < 20; d++ {
		brush := Circle(d)
		brush.IterateRows(50*d, 50, bm.Segment)
	}

	assert.NoError(t, util.SaveTmpPng("brush.png", bm.ToImage(color.Black, color.White)))
}
