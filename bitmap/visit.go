package bitmap

func VisitDotted(d int, visit func(x, y int)) func(x, y int) {
	type point struct{ x, y int32 }
	var points []point

	return func(x, y int) {
		for i := len(points) - 1; i >= 0; i-- {
			p := points[i]
			dx, dy := x-int(p.x), y-int(p.y)
			if dx*dx+dy*dy < d*d {
				return
			}
		}

		points = append(points, point{int32(x), int32(y)})
		visit(x, y)
	}
}
