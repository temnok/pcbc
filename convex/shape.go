package convex

type row = struct{ x0, x1 int32 }

type Shape struct {
	y0           int
	lower, upper []row
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

func (s *Shape) VisitRows(visit func(x0, x1, y int)) {
	for y := -len(s.lower); y < len(s.upper); y++ {
		var r *row
		if y < 0 {
			r = &s.lower[^y]
		} else {
			r = &s.upper[y]
		}

		if r.x0 < r.x1 {
			visit(int(r.x0), int(r.x1), s.y0+y)
		}
	}
}
