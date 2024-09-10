package util

import (
	"temnok/lab/bezier"
	"temnok/lab/bitmap"
	"temnok/lab/glyph"
	"temnok/lab/twod"
)

func RasterizeGlyph(bm *bitmap.Bitmap, g [][]twod.Coord) {
	gb := new(glyph.Builder)

	for _, contour := range g {
		bezier.CubicVisit(contour, gb.AddContourPoint)

		gb.FinishContour()
	}

	gb.Rasterize(bm.Segment)
}
