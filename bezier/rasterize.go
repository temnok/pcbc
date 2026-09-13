// Copyright © 2025 Alex Temnok. All rights reserved.

package bezier

import "math"

// Rasterize calls provided callback for each consecutive pixel along the cubic Bézier curve.
// Pixels are points with integer coordinates. Pixels A and B are consecutive if A != B && |Ax-Bx| <= 1 && |Ay-By| <= 1.
// The callback is NOT called for the very first pixel, to prevent duplicate points for adjacent curves.
// Cubic Bézier curve is represented by four points in xy array in form [Ax, Ay, Bx, By, Cx, Cy, Dx, Dy]
// where A and D are start and end points of the curve and B and C are control points.
func Rasterize(xy []float64, callback func(x, y int)) {
	aX, aY, bX, bY, cX, cY, dX, dY := xy[0], xy[1], xy[2], xy[3], xy[4], xy[5], xy[6], xy[7]

	parametricCurve(func(t float64) (int, int) {
		abX, abY := mix(t, aX, aY, bX, bY)
		bcX, bcY := mix(t, bX, bY, cX, cY)
		cdX, cdY := mix(t, cX, cY, dX, dY)

		abcX, abcY := mix(t, abX, abY, bcX, bcY)
		bcdX, bcdY := mix(t, bcX, bcY, cdX, cdY)

		x, y := mix(t, abcX, abcY, bcdX, bcdY)
		return int(math.Round(x)), int(math.Round(y))
	}, callback)
}

func mix(t, x0, y0, x1, y1 float64) (x, y float64) {
	return x0*(1-t) + x1*t, y0*(1-t) + y1*t
}
