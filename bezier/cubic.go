package bezier

import "math"

func CubicPoint(p []Point, t float64) Point {
	ab := p[0].Mix(p[1], t)
	bc := p[1].Mix(p[2], t)
	cd := p[2].Mix(p[3], t)
	abc := ab.Mix(bc, t)
	bcd := bc.Mix(cd, t)
	return abc.Mix(bcd, t)
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
	return 3 * totalDist(points)
}

func totalDist(points []Point) int {
	totalD := 0
	for i := 1; i < len(points); i++ {
		a, b := points[i-1], points[i]
		dx, dy := math.Round(math.Abs(a.X-b.X)), math.Round(math.Abs(a.Y-b.Y))
		d := int(max(dx, dy))
		totalD += d
	}
	return totalD
}
