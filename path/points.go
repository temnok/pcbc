// Copyright © 2025 Alex Temnok. All rights reserved.

package path

import (
	"temnok/pcbc/transform"
)

// Points represent list of arbitrary 2D points, not necessary a path.
type Points []Point

// Apply returns points transformed by a given 2D transformation.
func (points Points) Apply(t transform.T) Path {
	res := make(Path, len(points))

	for i, point := range points {
		res[i].X, res[i].Y = t.Apply(point.X, point.Y)
	}

	return res
}
