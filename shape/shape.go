package shape

import (
	"temnok/lab/path"
)

type row = struct {
	x0, x1, y int32
}

type Shape struct {
	rows []row
}

func FromContour(contour path.Path) *Shape {
	builder := new(Builder)
	contour.Visit(builder.AddPoint)
	return builder.Build()
}

func (s *Shape) IterateRows(iterator func(x0, x1, y int)) {
	s.IterateRowsXY(0, 0, iterator)
}

func (s *Shape) IterateRowsXY(x0, y0 int, iterator func(x0, x1, y int)) {
	for _, row := range s.rows {
		iterator(x0+int(row.x0), x0+int(row.x1), y0+int(row.y))
	}
}

func (s *Shape) IterateContour(contour path.Path, iterator func(x0, x1, y int)) {
	contour.Visit(func(x, y int) {
		s.IterateRowsXY(x, y, iterator)
	})
}

func (s *Shape) IterateContours(contours path.Paths, iterator func(x0, x1, y int)) {
	for _, contour := range contours {
		s.IterateContour(contour, iterator)
	}
}

func IterateContourRows(contour path.Path, iterator func(x0, x1, y int)) {
	FromContour(contour).IterateRows(iterator)
}

func IterateContoursRows(contours path.Paths, iterator func(x0, x1, y int)) {
	for _, contour := range contours {
		IterateContourRows(contour, iterator)
	}
}
