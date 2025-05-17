// Copyright © 2025 Alex Temnok. All rights reserved.

package bezier

import "math"

// Rasterize calls provided callback for each consecutive pixel along the cubic Bézier curve.
// Pixels are points with integer coordinates. Pixels A and B are consecutive if A != B && |Ax-Bx| <= 1 && |Ay-By| <= 1.
// The callback is NOT called for the very first pixel, to prevent duplicate points for adjacent curves.
// Cubic Bézier curve is represented by four points in xy array in form [Ax, Ay, Bx, By, Cx, Cy, Dx, Dy]
// where A and D are start and end points of the curve and B and C are control points.
func Rasterize(xy []float64, callback func(x, y int)) {
	ax, ay := round(xy[0]), round(xy[1])
	dx, dy := round(xy[6]), round(xy[7])

	recurse(xy, 0, 1, ax, ay, dx, dy, callback)
}

func recurse(xy []float64, i0, i1 float64, x0, y0, x1, y1 int, callback func(x, y int)) {
	if x0 == x1 && y0 == y1 {
		return
	}

	if abs(x0-x1) <= 1 && abs(y0-y1) <= 1 {
		callback(x1, y1)
		return
	}

	i := (i0 + i1) / 2

	abX, abY := mix(xy[0], xy[1], xy[2], xy[3], i)
	bcX, bcY := mix(xy[2], xy[3], xy[4], xy[5], i)
	cdX, cdY := mix(xy[4], xy[5], xy[6], xy[7], i)

	abcX, abcY := mix(abX, abY, bcX, bcY, i)
	bcdX, bcdY := mix(bcX, bcY, cdX, cdY, i)

	abcdX, abcdY := mix(abcX, abcY, bcdX, bcdY, i)
	x, y := round(abcdX), round(abcdY)

	recurse(xy, i0, i, x0, y0, x, y, callback)
	recurse(xy, i, i1, x, y, x1, y1, callback)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func mix(x0, y0, x1, y1, i float64) (float64, float64) {
	return x0*(1-i) + x1*i, y0*(1-i) + y1*i
}

func round(a float64) int {
	return int(math.Round(a))
}
