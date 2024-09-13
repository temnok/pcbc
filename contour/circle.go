package contour

import "temnok/lab/geom"

func Circle(r float64) []geom.XY {
	m := r * magic

	return []geom.XY{
		{r, 0},
		{r, m}, {m, r},
		{0, r},
		{-m, r}, {-r, m},
		{-r, 0},
		{-r, -m}, {-m, -r},
		{0, -r},
		{m, -r}, {r, -m},
		{r, 0},
	}
}
