// Copyright © 2026 Alex Temnok. All rights reserved.

package bezier

func parametricCurve(txy func(t float64) (x, y int), render func(x, y int)) {
	t0, t1 := 0.0, 1.0
	x0, y0 := txy(t0)
	x1, y1 := txy(t1)

	parametricRecurse(txy, render, t0, t1, x0, y0, x1, y1)
}

func parametricRecurse(txy func(t float64) (x, y int), render func(x, y int), t0, t1 float64, x0, y0, x1, y1 int) {
	if dx, dy := x0-x1, y0-y1; -1 <= dx && dx <= 1 && -1 <= dy && dy <= 1 {
		if x0 != x1 || y0 != y1 {
			render(x1, y1)
		}

		return
	}

	t := (t0 + t1) / 2
	x, y := txy(t)

	parametricRecurse(txy, render, t0, t, x0, y0, x, y)
	parametricRecurse(txy, render, t, t1, x, y, x1, y1)
}
