package bitmap

import (
	"image/color"
	"testing"
)

func TestBrushes(t *testing.T) {
	bm := NewBitmap(1000, 100)

	for r := 1; r < 20; r++ {
		brush := NewRoundBrush(r)
		bm.Segments(50*r, 50, brush)
	}

	savePng(t, "brush.png", bm.ToImage(color.Black, color.White))
}
