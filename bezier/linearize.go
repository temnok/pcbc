// Copyright © 2025 Alex Temnok. All rights reserved.

package bezier

import (
	"math"
)

// Linearize calls provided callback for selected points along the cubic Bézier path (one or more consecutive curves).
func Linearize(xy []float64, delta float64, callback func(x, y float64)) {
	for ; len(xy) >= 8; xy = xy[6:] {
		linearizeCurve(xy, delta, callback)
	}
}

func linearizeCurve(xy []float64, delta float64, callback func(x, y float64)) {
	x0, y0, x1, y1 := xy[0], xy[1], xy[6], xy[7]
	linearizeSegment(xy, delta, 0, 1, x0, y0, x1, y1, callback)

	callback(x1, y1)
}

func linearizeSegment(xy []float64, delta, t0, t1, x0, y0, x1, y1 float64, callback func(x, y float64)) {
	t := (t0 + t1) / 2
	x, y := cubicBezier(xy, t)

	if d := dist(x0, y0, x1, y1, x, y); math.IsInf(d, 0) || d <= delta {
		return
	}

	linearizeSegment(xy, delta, t0, t, x0, y0, x, y, callback)

	callback(x, y)

	linearizeSegment(xy, delta, t, t1, x, y, x1, y1, callback)
}

func dist(x0, y0, x1, y1, x, y float64) float64 {
	dx, dy := x1-x0, y1-y0
	return math.Abs(x*dy-dx*y+x1*y0-x0*y1) / math.Sqrt(dx*dx+dy*dy)
}
