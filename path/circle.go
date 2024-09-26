package path

import "math"

var roundK = 4 * (math.Sqrt(2) - 1) / 3

func Circle(r float64) Path {
	m := r * roundK

	return Path{
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
