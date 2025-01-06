// Copyright © 2025 Alex Temnok. All rights reserved.

package shape

type builder struct {
	rows, cols bounds
}

func (b *builder) addPoint(x, y int) {
	b.rows.addPoint(y, x)
	b.cols.addPoint(x, y)
}

func (b *builder) build() *Shape {
	s := &Shape{}

	minY, maxY := b.rows.getBounds()
	for y := minY; y < maxY; y++ {
		x0, x1 := b.rows.getBound(y)

		for x := x0; x <= x1+1; x++ {
			y0, y1 := b.cols.getBound(x)

			if x <= x1 && y0 <= y && y < y1 {
				continue
			}

			if x0 < x {
				s.rows = append(s.rows, row{
					x0: int32(x0),
					x1: int32(x),
					y:  int32(y),
				})
			}

			x0 = x + 1
		}
	}

	return s
}
