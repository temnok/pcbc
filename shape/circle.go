package shape

func Circle(d int) *Shape {
	builder := new(Builder)

	if d <= 0 {
		return builder.Build()
	}

	off := (d + 1) % 2
	dd := (d - 1 - off) * (d - off)
	for y, r := 0, (d+1)/2; y < r; y++ {
		x := 0
		for ((x+1)*(x+1)+y*y)*4 <= dd {
			x++
		}

		builder.AddPoint(-x-off, -y-off)
		builder.AddPoint(x, -y-off)
		builder.AddPoint(-x-off, y)
		builder.AddPoint(x, y)
	}

	return builder.Build()
}
