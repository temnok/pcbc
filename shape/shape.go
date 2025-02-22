// Copyright © 2025 Alex Temnok. All rights reserved.

package shape

import (
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

type row = struct {
	x0, x1, y int32
}

type Shape struct {
	rows []row
}

func FromContour(t transform.T, contour path.Path) *Shape {
	b := new(builder)
	contour.Visit(t, b.addPoint)
	return b.build()
}

func (s *Shape) IterateRows(iterator func(x0, x1, y int)) {
	s.IterateRowsXY(0, 0, iterator)
}

func (s *Shape) IterateRowsXY(x0, y0 int, iterator func(x0, x1, y int)) {
	for _, row := range s.rows {
		iterator(x0+int(row.x0), x0+int(row.x1), y0+int(row.y))
	}
}

func (s *Shape) IterateContour(t transform.T, contour path.Path, iterator func(x0, x1, y int)) {
	contour.Visit(t, func(x, y int) {
		s.IterateRowsXY(x, y, iterator)
	})
}

func (s *Shape) IterateContours(t transform.T, contours path.Paths, iterator func(x0, x1, y int)) {
	for _, contour := range contours {
		s.IterateContour(t, contour, iterator)
	}
}

func IterateContourRows(t transform.T, contour path.Path, iterator func(x0, x1, y int)) {
	FromContour(t, contour).IterateRows(iterator)
}

func IterateContoursRows(t transform.T, contours path.Paths, iterator func(x0, x1, y int)) {
	for _, contour := range contours {
		IterateContourRows(t, contour, iterator)
	}
}
