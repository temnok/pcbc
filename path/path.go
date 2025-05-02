// Copyright © 2025 Alex Temnok. All rights reserved.

package path

import (
	"temnok/pcbc/cbc"
	"temnok/pcbc/transform"
)

// Path consists of cubic Bézier curves as a sequence of on-path points separated by pairs of control points,
// for example {p1,c1,c2,p2,c3,c4,p3}. Length of path should be 0 or 1+n*3, where n >= 0.
// If the last point is the same as the first one, the path represents a closed contour.
type Path []Point

// Apply returns new path transformed by a given 2D transformation.
func (path Path) Apply(t transform.T) Path {
	res := make(Path, len(path))

	for i, point := range path {
		res[i].X, res[i].Y = t.Apply(point.X, point.Y)
	}

	return res
}

// ForEachPixel calls provided callback for each interpolated point on the path with integer coordinates (pixel).
// The callback is called at least one time for a non-empty path.
func (path Path) ForEachPixel(t transform.T, visit func(x, y int)) {
	if len(path) == 0 {
		return
	}

	a := path[0].Apply(t)
	visit(a.RoundXY())

	for i := 0; i+3 < len(path); i += 3 {
		c1, c2, b := path[i+1].Apply(t), path[i+2].Apply(t), path[i+3].Apply(t)

		cbc.Rasterize([]float64{a.X, a.Y, c1.X, c1.Y, c2.X, c2.Y, b.X, b.Y}, visit)

		a = b
	}
}

// ForEachPixelDist calls provided callback for interpolated points on the path, separated by given distance.
// For example, it could be used to draw a dotted line.
func (path Path) ForEachPixelDist(t transform.T, dist int, jump func(x, y int)) {
	var prevX, prevY int
	started := false

	path.ForEachPixel(t, func(x, y int) {
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

func (path Path) Clone(n int, dx, dy float64) Paths {
	res := make(Paths, 0, n)

	for i := 0; i < n; i++ {
		k := float64(i) - float64(n-1)/2
		res = append(res, path.Apply(transform.Move(dx*k, dy*k)))
	}

	return res
}

func (path Path) Center(t transform.T) Point {
	var b Bounds
	b.IncludePath(t, path)
	return b.Center()
}
