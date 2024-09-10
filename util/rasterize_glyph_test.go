package util

import (
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"temnok/lab/bitmap"
	"temnok/lab/twod"
	"testing"
)

func TestGlyph(t *testing.T) {
	bm := bitmap.NewBitmap(1000, 1000)

	transform := twod.Move(twod.Coord{X: 500, Y: 500}).ScaleLocked(200).Rotate(150 * math.Pi / 180)

	glyph := TransformAllPoints(transform, [][]twod.Coord{
		{
			{-1, -1}, {-1, -1},
			{1, -1}, {1, -1}, {1, -1},
			{0, 1}, {0, 1}, {0, 1},
			{-1, -1}, {-1, -1},
			//}, {
			//	{-0.5, -0.5}, {-0.5, 0.5},
			//	{0.5, 0.5}, {0.5, -0.5}, {0.5, -0.5},
			//	{-0.5, -0.5}, {-0.5, -0.5},
		},
	})

	RasterizeGlyph(bm, glyph)

	savePng(t, "glyph.png", bm.ToImage(color.Black, color.White))
}

func savePng(t *testing.T, name string, im image.Image) {
	_ = os.Mkdir("tmp", 0770)

	f, err := os.Create("tmp/" + name)
	assert.NoError(t, err)

	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())
}
