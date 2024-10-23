package path

import (
	"math"
)

// cubicVisit iterates over all pixels for a given cubic Bezier curve.
func cubicVisit(points []Point, visit func(x, y int)) {
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

func cubicPoint(p []Point, t float64) Point {
	ab := mix(p[0], p[1], t)
	bc := mix(p[1], p[2], t)
	cd := mix(p[2], p[3], t)
	abc := mix(ab, bc, t)
	bcd := mix(bc, cd, t)
	return mix(abc, bcd, t)
}

func cubicSteps(points []Point) int {
	return 3 * cubicDist(points)
}

func cubicDist(points []Point) int {
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
func mix(a, b Point, t float64) Point {
	return Point{X: a.X*(1-t) + b.X*t, Y: a.Y*(1-t) + b.Y*t}
}
