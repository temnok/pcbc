// Copyright © 2025 Alex Temnok. All rights reserved.

package eda

import (
	"math"
	"temnok/pcbc/path"
)

func TrackV2(p0, p1 path.Point, steps ...float64) Track {
	if p0 == p1 {
		return Track{}
	}

	x, y := p0.XY()
	x1, y1 := p1.XY()
	sx, sy := sign(x1-x), sign(y1-y)

	out := Track{{x, y}}

	for _, step := range steps {
		if turnLeft := math.Signbit(step); turnLeft {
			step = -step
			sx, sy = sign(sx-sy), sign(sy+sx)
		} else {
			sx, sy = sign(sx+sy), sign(sy-sx)
		}

		x += sx * step
		y += sy * step
		out = append(out, path.Point{x, y})
	}

	if x != x1 || y != y1 {
		dx, dy := x1-x, y1-y
		sdx, sdy := sign(dx), sign(dy)
		dx, dy = dx*sdx, dy*sdy
		//fmt.Printf("dx=%v dy=%v sdx=%v sdy=%v sx=%v sy=%v\n", dx, dy, sdx, sdy, sx, sy)

		if dx != 0 && dy != 0 && dx != dy {
			if sx != 0 && sy != 0 {
				if dx > dy {
					dx = sdx * dy
				} else {
					dy = sdy * dx
				}
			} else {
				if dx > dy {
					dx, dy = sdx*(dx-dy), 0
				} else {
					dy, dx = sdy*(dy-dx), 0
				}
			}

			out = append(out, path.Point{x + dx, y + dy})
		}

		out = append(out, path.Point{x1, y1})
	}

	return out
}
