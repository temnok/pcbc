// Copyright © 2025 Alex Temnok. All rights reserved.

package eda

import (
	"github.com/temnok/pcbc/path"
	"math"
)

func LinearTrack(p0, p1 path.Point, steps ...float64) path.Path {
	if p0 == p1 {
		return nil
	}

	x, y := p0.XY()
	x1, y1 := p1.XY()
	sx, sy := sign(x1-x), sign(y1-y)

	out := []path.Point{{x, y}}

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

			out = append(out, path.Point{x, y})
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

			out = append(out, path.Point{x + dx, y + dy})
		}

		out = append(out, path.Point{x1, y1})
	}

	return path.Linear(out)
}

func sign(val float64) float64 {
	switch {
	case val < 0:
		return -1
	case val > 0:
		return 1
	default:
		return 0
	}
}
