package monodraw

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"image/color"
	"image/png"
	"os"
	"temnok/lab/glyph"
	"testing"
)

func TestGlyph(t *testing.T) {
	bm := NewBitmap(1000, 1000)

	gb := new(glyph.Builder)

	gb.AddContourPoint(400, 400)
	gb.AddContourPoint(600, 400)
	gb.AddContourPoint(500, 600)
	gb.AddContourPoint(400, 400)
	gb.FinishContour()

	gb.Rasterize(func(x0, x1, y int) {
		fmt.Printf("x0=%v, x1=%v, y=%v\n", x0, x1, y)
		bm.Segment(x0, x1, y)
	})

	f, err := os.Create("glyph.png")
	assert.NoError(t, err)

	im := bm.ToImage(color.Black, color.White)
	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())
}
