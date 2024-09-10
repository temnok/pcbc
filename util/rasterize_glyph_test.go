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

	glyph := [][]twod.Coord{
		{
			{-1, -1}, {-1, -1},
			{1, -1}, {1, -1}, {1, -1},
			{0, 1}, {0, 1}, {0, 1},
			{-1, -1}, {-1, -1},
		}, {
			{-0.5, -0.5}, {-0.5, 0.5},
			{0.5, 0.5}, {0.5, -0.5}, {0.5, -0.5},
			{-0.5, -0.5}, {-0.5, -0.5},
		},
	}

	RasterizeGlyph(bm, TransformAllPoints(transform, glyph))

	savePng(t, "glyph.png", bm.ToImage(color.Black, color.White))
}

func savePng(t *testing.T, name string, im image.Image) {
	_ = os.Mkdir("tmp", 0770)

	f, err := os.Create("tmp/" + name)
	assert.NoError(t, err)

	assert.NoError(t, png.Encode(f, im))
	assert.NoError(t, f.Close())
}

func TestCircleGlyph(t *testing.T) {
	var magic = 4 * (math.Sqrt(2) - 1) / 3

	glyph := [][]twod.Coord{{
		{1, 0},
		{1, magic}, {magic, 1},
		{0, 1},
		{-magic, 1}, {-1, magic},
		{-1, 0},
		{-1, -magic}, {-magic, -1},
		{0, -1},
		{magic, -1}, {1, -magic},
		{1, 0},
	}}

	bm := bitmap.NewBitmap(1500, 100)

	for d := 1.0; d < 30; d++ {
		transform := twod.Move(twod.Coord{X: 50 * d, Y: 50}).ScaleLocked(d)
		RasterizeGlyph(bm, TransformAllPoints(transform, glyph))
	}

	savePng(t, "circle.png", bm.ToImage(color.Black, color.White))
}

func TestRoundedRectGlyph(t *testing.T) {
	glyph := [][]twod.Coord{
		RoundedRectContour(3, 4, 1),
		RoundedRectContour(3-0.4, 4-0.4, 0.8),
	}

	bm := bitmap.NewBitmap(1500, 100)

	for d := 1.0; d < 20; d++ {
		transform := twod.Move(twod.Coord{X: 75 * d, Y: 50}).ScaleLocked(d)
		RasterizeGlyph(bm, TransformAllPoints(transform, glyph))
	}

	savePng(t, "rect.png", bm.ToImage(color.Black, color.White))
}
