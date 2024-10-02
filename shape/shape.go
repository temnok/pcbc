package shape

import (
	"temnok/lab/path"
)

type row = struct{ x0, x1 int32 }

type Shape struct {
	y0           int
	lower, upper []row
}

func FromContour(contour path.Path) *Shape {
	shape := new(Shape)
	contour.Visit(shape.AddPoint)
	return shape
}

func (s *Shape) AddPoint(x, y int) {
	if s.upper == nil && s.lower == nil {
		s.y0 = y
	}

	y -= s.y0
	rows := &s.upper
	if y < 0 {
		rows = &s.lower
		y = ^y
	}

	for y >= len(*rows) {
		*rows = append(*rows, row{})
	}

	r := &(*rows)[y]
	if r.x0 == r.x1 {
		*r = row{int32(x), int32(x + 1)}
	} else {
		*r = row{min(r.x0, int32(x)), max(r.x1, int32(x+1))}
	}
}

func (s *Shape) IterateRows(iterator func(x0, x1, y int)) {
	s.IterateRowsXY(0, 0, iterator)
}

func (s *Shape) IterateRowsXY(x0, y0 int, iterator func(x0, x1, y int)) {
	for y := -len(s.lower); y < len(s.upper); y++ {
		var r *row
		if y < 0 {
			r = &s.lower[^y]
		} else {
			r = &s.upper[y]
		}

		if r.x0 < r.x1 {
			iterator(x0+int(r.x0), x0+int(r.x1), y0+s.y0+y)
		}
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
