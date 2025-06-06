// Copyright © 2025 Alex Temnok. All rights reserved.

package bezier

func cubicBezier(xy []float64, t float64) (x, y float64) {
	aX, aY, bX, bY, cX, cY, dX, dY := xy[0], xy[1], xy[2], xy[3], xy[4], xy[5], xy[6], xy[7]
	abX, abY := mix(aX, aY, bX, bY, t)
	bcX, bcY := mix(bX, bY, cX, cY, t)
	cdX, cdY := mix(cX, cY, dX, dY, t)

	abcX, abcY := mix(abX, abY, bcX, bcY, t)
	bcdX, bcdY := mix(bcX, bcY, cdX, cdY, t)

	return mix(abcX, abcY, bcdX, bcdY, t)
}

func mix(x0, y0, x1, y1, i float64) (float64, float64) {
	return x0*(1-i) + x1*i, y0*(1-i) + y1*i
}
