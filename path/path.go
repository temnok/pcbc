package path

import (
	"temnok/lab/bezier"
	"temnok/lab/geom"
	"temnok/lab/line"
)

// Visit iterates over all coordinates on cubic Bezier path with integer X and Y values.
// Cubic Bezier path is represented by 3*n+1 points, like p1,c1,c2,p2,c3,c4,p3,c5,c6,p4,...
// where pI are points on the path and cI are cubic-Bezier control points.
func Visit(path []geom.XY, visit func(x, y int)) {
	prev := path[0].Round()
	visit(int(prev.X), int(prev.Y))

	visitNext := func(x, y int) {
		cur := geom.XY{X: float64(x), Y: float64(y)}
		if cur != prev {
			visit(x, y)

			prev = cur
		}
	}

	for s := 0; s+3 < len(path); s += 3 {
		points := path[s : s+4]

		if points[0] == points[1] && points[2] == points[3] { // straight line
			a, b := points[0].Round(), points[2].Round()

			line.Visit(int(a.X), int(a.Y), int(b.X), int(b.Y), visitNext)
		} else {
			bezier.Visit(points, visitNext)
		}
	}
}
