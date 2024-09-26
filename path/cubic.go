package path

import (
	"math"
	"temnok/lab/geom"
)

// cubicVisit iterates over all pixels for a given cubic Bezier curve.
func cubicVisit(points []geom.XY, visit func(x, y int)) {
	prev := points[0].Round()

	steps := cubicSteps(points)
	for i := 1; i <= steps; i++ {

		t := float64(i) / float64(steps)
		cur := cubicPoint(points, t).Round()
		if cur == prev {
			continue
		}

		visit(int(cur.X), int(cur.Y))
		prev = cur
	}
}

func cubicPoint(p []geom.XY, t float64) geom.XY {
	ab := mix(p[0], p[1], t)
	bc := mix(p[1], p[2], t)
	cd := mix(p[2], p[3], t)
	abc := mix(ab, bc, t)
	bcd := mix(bc, cd, t)
	return mix(abc, bcd, t)
}

func cubicSteps(points []geom.XY) int {
	return 3 * cubicDist(points)
}

func cubicDist(points []geom.XY) int {
	totalD := 0
	for i := 1; i < len(points); i++ {
		a, b := points[i-1], points[i]
		dx, dy := math.Round(math.Abs(a.X-b.X)), math.Round(math.Abs(a.Y-b.Y))
		d := int(max(dx, dy))
		totalD += d
	}
	return totalD
}

//go:inline
func mix(a, b geom.XY, t float64) geom.XY {
	return geom.XY{X: a.X*(1-t) + b.X*t, Y: a.Y*(1-t) + b.Y*t}
}
