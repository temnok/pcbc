// Copyright © 2025 Alex Temnok. All rights reserved.

package path

func Rect(w, h float64) Path {
	x, y := w/2, h/2

	return Linear([]Point{
		{x, y},
		{x, -y},
		{-x, -y},
		{-x, y},
		{x, y},
	})
}

func RoundRect(w, h, r float64) Path {
	x1, y1 := w/2, h/2
	r = min(r, x1, y1)
	x0, y0 := x1-r, y1-r

	m := r * circleK

	return Path{
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

func CutRect(w, h, r float64) Path {
	x, y := w/2, h/2

	return Linear([]Point{
		{x - r, y}, {-x + r, y},
		{-x, y - r}, {-x, -y + r},
		{-x + r, -y}, {x - r, -y},
		{x, -y + r}, {x, y - r},
		{x - r, y},
	})
}
