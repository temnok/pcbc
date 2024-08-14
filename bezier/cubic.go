package bezier

import "math"

func CubicPoint(p []Point, t float64) Point {
	ab := p[0].Mix(p[1], t)
	cd := p[2].Mix(p[3], t)
	return ab.Mix(cd, t)
}

func CubicVisit(points []Point, visit func(x, y int)) {
	prev := points[0].Round()
	visit(int(prev.X), int(prev.Y))

	steps := cubicSteps(points)
	for i := 1; i <= steps; i++ {
		t := float64(i) / float64(steps)
		cur := CubicPoint(points, t).Round()
		if cur == prev {
			continue
		}

		visit(int(cur.X), int(cur.Y))
		prev = cur
	}
}

func cubicSteps(points []Point) int {
	return 3 * int(maxDist(points))
}

func maxDist(points []Point) float64 {
	maxD := 0.0
	for i := 1; i < len(points); i++ {
		a, b := points[i-1], points[i]
		if d := max(math.Abs(a.X-b.X), math.Abs(a.Y-b.Y)); d > maxD {
			maxD = d
		}
	}
	return maxD
}
