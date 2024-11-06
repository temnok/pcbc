package path

func Rect(w, h float64) Paths {
	x, y := w/2, h/2

	return Paths{
		Lines([]Point{
			{x, y},
			{x, -y},
			{-x, -y},
			{-x, y},
			{x, y},
		}),
	}
}

func RoundRect(w, h, r float64) Paths {
	x1, y1 := w/2, h/2
	r = min(r, x1, y1)
	x0, y0 := x1-r, y1-r

	m := r * roundK

	return Paths{
		[]Point{
			{x1, y0},
			{x1, y0 + m}, {x0 + m, y1},
			{x0, y1},
			{x0, y1}, {-x0, y1},
			{-x0, y1},
			{-x0 - m, y1}, {-x1, y0 + m},
			{-x1, y0},
			{-x1, y0}, {-x1, -y0},
			{-x1, -y0},
			{-x1, -y0 - m}, {-x0 - m, -y1},
			{-x0, -y1},
			{-x0, -y1}, {x0, -y1},
			{x0, -y1},
			{x0 + m, -y1}, {x1, -y0 - m},
			{x1, -y0},
			{x1, -y0}, {x1, y0},
			{x1, y0},
		},
	}
}

func CutRect(w, h, r float64) Paths {
	x, y := w/2, h/2

	return Paths{
		Lines([]Point{
			{x - r, y}, {-x + r, y},
			{-x, y - r}, {-x, -y + r},
			{-x + r, -y}, {x - r, -y},
			{x, -y + r}, {x, y - r},
			{x - r, y},
		}),
	}
}
