package path

import "temnok/lab/geom"

// Points represent list of arbitrary 2D points, not necessary a path.
type Points []geom.XY

// Transform returns points transformed by a given 2D transformation.
func (points Points) Transform(transform geom.Transform) Path {
	res := make(Path, len(points))

	for i, point := range points {
		res[i] = transform.Point(point)
	}

	return res
}
