// Copyright © 2025 Alex Temnok. All rights reserved.

package shape

import (
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

type row = struct {
	x0, x1, y int32
}

type Shape struct {
	rows []row
}

func New(contour path.Path, t transform.T) *Shape {
	b := new(builder)
	contour.Rasterize(t, b.addPoint)
	return b.build()
}

func (s *Shape) ForEachRow(iterator func(x0, x1, y int)) {
	s.ForEachRowWithOffset(0, 0, iterator)
}

func (s *Shape) ForEachRowWithOffset(x, y int, iterator func(x0, x1, y int)) {
	for _, row := range s.rows {
		iterator(x+int(row.x0), x+int(row.x1), y+int(row.y))
	}
}

func (s *Shape) ForEachPathsPixel(paths path.Paths, t transform.T, iterator func(x0, x1, y int)) {
	paths.Rasterize(t, func(x, y int) {
		s.ForEachRowWithOffset(x, y, iterator)
	})
}

func ForEachRow(shapeContours path.Paths, t transform.T, iterator func(x0, x1, y int)) {
	for _, contour := range shapeContours {
		New(contour, t).ForEachRow(iterator)
	}
}
