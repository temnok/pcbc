// Copyright © 2025 Alex Temnok. All rights reserved.

package path

import "math"

var roundK = 4 * (math.Sqrt(2) - 1) / 3

func Circle(d float64) Paths {
	r := d * 0.5
	m := r * roundK

	return Paths{
		Path{
			{r, 0},
			{r, m}, {m, r},
			{0, r},
			{-m, r}, {-r, m},
			{-r, 0},
			{-r, -m}, {-m, -r},
			{0, -r},
			{m, -r}, {r, -m},
			{r, 0},
		},
	}
}
