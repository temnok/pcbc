package bitmap

import (
	"image/color"
	"math"
	"temnok/lab/bezier"
	"temnok/lab/glyph"
	"temnok/lab/t2d"
	"testing"
)

func TestGlyph(t *testing.T) {
	bm := NewBitmap(1000, 1000)

	//gb := &glyph.Interpolator{Builder: new(glyph.Builder)}
	gb := new(glyph.Builder)
	transform := t2d.Move(t2d.Vector{500, 500}).ScaleLocked(200).Rotate(150 * math.Pi / 180)

	path := bezier.TransformPoints(transform, []bezier.Point{
		{-1, -1}, {-1, -1},
		{1, -1}, {1, -1}, {1, -1},
		{0, 1}, {0, 1}, {0, 1},
		{-1, -1}, {-1, -1},
	})
	bezier.CubicVisit(path, func(x, y int) {
		gb.AddContourPoint(x, y)
	})
	path = bezier.TransformPoints(transform, []bezier.Point{
		{-0.5, -0.5}, {-0.5, 0.5},
		{0.5, 0.5}, {0.5, -0.5}, {0.5, -0.5},
		{-0.5, -0.5}, {-0.5, -0.5},
	})
	bezier.CubicVisit(path, func(x, y int) {
		gb.AddContourPoint(x, y)
	})

	gb.FinishContour()

	gb.Rasterize(func(x0, x1, y int) {
		bm.Segment(x0, x1, y)
	})

	savePng(t, "glyph.png", bm.ToImage(color.Black, color.White))
}
