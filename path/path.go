package path

import (
	"temnok/pcbc/transform"
)

// Path consists of cubic-Bezier curves as a sequence of on-path points separated by pairs of control points,
// for example {p1,c1,c2,p2,c3,c4,p3}. Length of path should be 0 or n*3-2, where n > 0.
// If the last point is the same as the first one, the path represents a closed contour.
type Path []Point

// Apply returns new path transformed by a given 2D transformation.
func (path Path) Apply(t transform.Transform) Path {
	return Points(path).Apply(t)
}

// Visit calls provided callback for each interpolated point on the path with integer coordinates.
// The callback is called at least one time for a non-empty path.
func (path Path) Visit(visit func(x, y int)) {
	if len(path) == 0 {
		return
	}

	a := path[0]
	visit(a.RoundXY())

	for i := 0; i+3 < len(path); i += 3 {
		c1, c2, b := path[i+1], path[i+2], path[i+3]

		if a == c1 && c2 == b {
			linearVisit(a, b, visit)
		} else {
			cubicVisit([]Point{a, c1, c2, b}, visit)
		}

		a = b
	}
}

// Jump calls provided callback for interpolated points on the path, separated by given distance.
// For example, it could be used to draw a dotted line.
func (path Path) Jump(dist int, jump func(x, y int)) {
	var prevX, prevY int
	started := false

	path.Visit(func(x, y int) {
		if !started {
			started = true
		} else {
			dx, dy := x-prevX, y-prevY
			if dx*dx+dy*dy < dist*dist {
				return
			}
		}

		jump(x, y)
		prevX, prevY = x, y
	})
}

func (path Path) Center(t transform.Transform) Point {
	var b Bounds
	b.AddPath(t, path)
	return b.Center()
}
