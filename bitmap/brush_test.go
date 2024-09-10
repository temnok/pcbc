package bitmap

import (
	"image/color"
	"testing"
)

func TestBrushes(t *testing.T) {
	bm := NewBitmap(1000, 100)

	for d := 1; d < 20; d++ {
		brush := NewRoundBrush(d)
		bm.Segments(50*d, 50, brush)
	}

	savePng(t, "brush.png", bm.ToImage(color.Black, color.White))
}
