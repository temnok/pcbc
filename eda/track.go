// Copyright © 2025 Alex Temnok. All rights reserved.

package eda

import (
	"math"
	"temnok/pcbc/path"
)

func Track(p0, p1 path.Point, steps ...float64) path.Path {
	if p0 == p1 {
		return path.Path{}
	}

	x, y := p0.XY()
	x1, y1 := p1.XY()
	sx, sy := sign(x1-x), sign(y1-y)

	p := path.Point{x, y}
	out := path.Path{p}

	for i, step := range steps {
		if i > 0 {
			if turnLeft := math.Signbit(step); turnLeft {
				step = -step
				sx, sy = sign(sx-sy), sign(sy+sx)
			} else {
				sx, sy = sign(sx+sy), sign(sy-sx)
			}
		}

		if step != 0 {
			x += sx * step
			y += sy * step

			out = append(out, p)
			p = path.Point{x, y}
			out = append(out, p, p)
		}
	}

	if x != x1 || y != y1 {
		dx, dy := x1-x, y1-y
		sdx, sdy := sign(dx), sign(dy)
		dx, dy = dx*sdx, dy*sdy
		//fmt.Printf("dx=%v dy=%v sdx=%v sdy=%v sx=%v sy=%v\n", dx, dy, sdx, sdy, sx, sy)

		if dx != 0 && dy != 0 && dx != dy {
			if sx != 0 && sy != 0 {
				if dx > dy {
					dx, dy = sdx*dy, sdy*dy
				} else {
					dy, dx = sdy*dx, sdx*dx
				}
			} else {
				if dx > dy {
					dx, dy = sdx*(dx-dy), 0
				} else {
					dy, dx = sdy*(dy-dx), 0
				}
			}

			out = append(out, p)
			p = path.Point{x + dx, y + dy}
			out = append(out, p, p)
		}

		out = append(out, p)
		p = path.Point{x1, y1}
		out = append(out, p, p)
	}

	return out
}
