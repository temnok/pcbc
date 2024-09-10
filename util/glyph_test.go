package util

import (
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"temnok/lab/bezier"
	"temnok/lab/bitmap"
	"temnok/lab/t2d"
	"testing"
)

func TestGlyph(t *testing.T) {
	bm := bitmap.NewBitmap(1000, 1000)

	transform := t2d.Move(t2d.Vector{500, 500}).ScaleLocked(200).Rotate(150 * math.Pi / 180)

	glyph := TransformAllPoints(transform, [][]bezier.Point{
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
	os.Mkdir("tmp", 0770)

	f, err := os.Create("tmp/" + name)
	assert.NoError(t, err)

	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())
}
