package path

import (
	"temnok/pcbc/geom"
	"temnok/pcbc/transform"
)

// Points represent list of arbitrary 2D points, not necessary a path.
type Points []geom.XY

// Apply returns points transformed by a given 2D transformation.
func (points Points) Apply(t transform.Transform) Path {
	res := make(Path, len(points))

	for i, point := range points {
		res[i].X, res[i].Y = t.Apply(point.X, point.Y)
	}

	return res
}
