package convex

func Circle(d int) *Shape {
	shape := new(Shape)

	if d <= 0 {
		return shape
	}

	off := (d + 1) % 2
	dd := (d - 1 - off) * (d - off)
	for y, r := 0, (d+1)/2; y < r; y++ {
		x := 0
		for ((x+1)*(x+1)+y*y)*4 <= dd {
			x++
		}

		shape.AddPoint(-x-off, -y-off)
		shape.AddPoint(x, -y-off)
		shape.AddPoint(-x-off, y)
		shape.AddPoint(x, y)
	}

	return shape
}
