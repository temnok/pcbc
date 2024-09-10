package bezier

import (
	"math"
	"temnok/lab/twod"
)

func CubicPoint(p []twod.Coord, t float64) twod.Coord {
	ab := mix(p[0], p[1], t)
	bc := mix(p[1], p[2], t)
	cd := mix(p[2], p[3], t)
	abc := mix(ab, bc, t)
	bcd := mix(bc, cd, t)
	return mix(abc, bcd, t)
}

func CubicVisit(allPoints []twod.Coord, visit func(x, y int)) {
	prev := round(allPoints[0])
	visit(int(prev.X), int(prev.Y))

	for s := 0; s+3 < len(allPoints); s += 3 {
		points := allPoints[s : s+4]
		steps := cubicSteps(points)
		for i := 1; i <= steps; i++ {
			t := float64(i) / float64(steps)
			cur := round(CubicPoint(points, t))
			if cur == prev {
				continue
			}

			if math.Abs(cur.Y-prev.Y) > 1 {
				panic("CubicVisit(): interpolation error")
			}

			visit(int(cur.X), int(cur.Y))
			prev = cur
		}
	}
}

func cubicSteps(points []twod.Coord) int {
	return 3 * totalDist(points)
}

func totalDist(points []twod.Coord) int {
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
func mix(a, b twod.Coord, t float64) twod.Coord {
	return twod.Coord{X: a.X*(1-t) + b.X*t, Y: a.Y*(1-t) + b.Y*t}
}

//go:inline
func round(a twod.Coord) twod.Coord {
	return twod.Coord{X: math.Round(a.X), Y: math.Round(a.Y)}
}
