package contour

import (
	"math"
	"temnok/lab/geom"
	"temnok/lab/path"
)

var magic = 4 * (math.Sqrt(2) - 1) / 3

func Rect(w, h float64) path.Path {
	x, y := w/2, h/2

	return Lines([]geom.XY{
		{x, y},
		{x, -y},
		{-x, -y},
		{-x, y},
		{x, y},
	})
}

func RoundRect(w, h, r float64) path.Path {
	x1, y1 := w/2, h/2
	r = min(r, x1, y1)
	x0, y0 := x1-r, y1-r

	m := r * magic

	return []geom.XY{
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

func CutRect(w, h, r float64) path.Path {
	x, y := w/2, h/2
	return Lines([]geom.XY{
		{x - r, y}, {-x + r, y},
		{-x, y - r}, {-x, -y + r},
		{-x + r, -y}, {x - r, -y},
		{x, -y + r}, {x, y - r},
		{x - r, y},
	})
}
