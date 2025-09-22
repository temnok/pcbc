// Copyright © 2025 Alex Temnok. All rights reserved.

package eda

import "temnok/pcbc/path"

func TrackV2(p0 path.Point, dx, dy float64, ds ...float64) Track {
	x0, y0 := p0.XY()
	x, y, sx, sy := x0+dx, y0+dy, sign(dx), sign(dy)

	out := Track{{x0, y0}, {x, y}}

	for _, d := range ds {
		if d < 0 {
			d = -d
			sx, sy = sign(sx-sy), sign(sy+sx)
		} else {
			sx, sy = sign(sx+sy), sign(sy-sx)
		}

		x += sx * d
		y += sy * d

		out = append(out, path.Point{x, y})
	}

	return out
}
