// Copyright © 2025 Alex Temnok. All rights reserved.

package cbc

import "math"

// Rasterize calls provided callback for each consecutive pixel along the CBC (cubic Bézier curve).
// Pixels are points with integer coordinates. Pixels A and B are consecutive if A != B && |Ax-Bx| <= 1 && |Ay-By| <= 1.
// The callback is NOT called for the very first pixel, to prevent duplicate calls for adjacent curves.
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

	if max(x0-x1, x1-x0) <= 1 && max(y0-y1, y1-y0) <= 1 {
		callback(x1, y1)
		return
	}

	i := (i0 + i1) / 2

	abx, aby := mix(xy[0], xy[1], xy[2], xy[3], i)
	bcx, bcy := mix(xy[2], xy[3], xy[4], xy[5], i)
	cdx, cdy := mix(xy[4], xy[5], xy[6], xy[7], i)

	abcx, abcy := mix(abx, aby, bcx, bcy, i)
	bcdx, bcdy := mix(bcx, bcy, cdx, cdy, i)

	abcdx, abcdy := mix(abcx, abcy, bcdx, bcdy, i)
	x, y := round(abcdx), round(abcdy)

	recurse(xy, i0, i, x0, y0, x, y, callback)
	recurse(xy, i, i1, x, y, x1, y1, callback)
}

func mix(x0, y0, x1, y1, i float64) (float64, float64) {
	return x0*(1-i) + x1*i, y0*(1-i) + y1*i
}

func round(a float64) int {
	return int(math.Round(a))
}
