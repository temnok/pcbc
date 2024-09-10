package util

import (
	"math"
	"temnok/lab/bezier"
	"temnok/lab/bitmap"
	"temnok/lab/glyph"
	"temnok/lab/twod"
)

func RasterizeGlyph(bm *bitmap.Bitmap, glph [][]twod.Coord) {
	gb := new(glyph.Builder)

	for _, contour := range glph {
		bezier.CubicVisit(contour, gb.AddContourPoint)

		gb.FinishContour()
	}

	gb.Rasterize(bm.Segment)
}

func RoundedRectContour(w, h, r float64) []twod.Coord {
	x1, y1 := w/2, h/2
	r = min(r, x1, y1)
	m := r * 4 * (math.Sqrt(2) - 1) / 3
	x0, y0 := x1-r, y1-r

	return []twod.Coord{
		{x1, y0},
		{x1, y0 + m}, {x0 + m, y1},
		{x0, y1},
		{x0, y1}, {-x0, y1},
		{-x0, y1},
		{-x0 - m, y1}, {-x1, y0 + m},
		{-x1, y0},
		{-x1, y0}, {-x1, -y0},
		{-x1, -y0},
		{-x1, -y0 - m}, {-x0 - m, -y1},
		{-x0, -y1},
		{-x0, -y1}, {x0, -y1},
		{x0, -y1},
		{x0 + m, -y1}, {x1, -y0 - m},
		{x1, -y0},
		{x1, -y0}, {x1, y0},
		{x1, y0},
	}
}
