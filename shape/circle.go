// Copyright © 2025 Alex Temnok. All rights reserved.

package shape

func Circle(d int) *Shape {
	b := new(builder)

	if d <= 0 {
		return b.build()
	}

	off := (d + 1) % 2
	dd := (d - 1 - off) * (d - off)
	for y, r := 0, (d+1)/2; y < r; y++ {
		x := 0
		for ((x+1)*(x+1)+y*y)*4 <= dd {
			b.addPoint(-x-off, -y-off)
			b.addPoint(x, -y-off)
			b.addPoint(-x-off, y)
			b.addPoint(x, y)

			x++
		}

		b.addPoint(-x-off, -y-off)
		b.addPoint(x, -y-off)
		b.addPoint(-x-off, y)
		b.addPoint(x, y)
	}

	return b.build()
}
