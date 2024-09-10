package bezier

import (
	"math"
	"temnok/lab/t2d"
)

func CubicPoint(p []Point, t float64) Point {
	ab := p[0].Mix(p[1], t)
	bc := p[1].Mix(p[2], t)
	cd := p[2].Mix(p[3], t)
	abc := ab.Mix(bc, t)
	bcd := bc.Mix(cd, t)
	return abc.Mix(bcd, t)
}

func TransformPoints(t t2d.Transform, points []Point) []Point {
	res := make([]Point, len(points))

	for i, p := range points {
		r := t.Point(t2d.Vector{p.X, p.Y})
		res[i].X, res[i].Y = r[0], r[1]
	}

	return res
}

func CubicVisit(allPoints []Point, visit func(x, y int)) {
	prev := allPoints[0].Round()
	visit(int(prev.X), int(prev.Y))

	for s := 0; s+3 < len(allPoints); s += 3 {
		points := allPoints[s : s+4]
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
