package contour

import (
	"temnok/lab/path"
)

func Circle(r float64) path.Path {
	m := r * magic

	return path.Path{
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
