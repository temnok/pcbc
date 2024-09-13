package line

// Visit iterates over all coordinates with integer X and Y values on
// line with end points (x0, y0) and (x1, y1).
func Visit(x0, y0, x1, y1 int, onPix func(x, y int)) {
	dx, ix := x1-x0, 1
	if dx < 0 {
		dx, ix = -dx, -ix
	}

	dy, iy := y1-y0, 1
	if dy < 0 {
		dy, iy = -dy, -iy
	}

	ax, ay := (dy+1)/2, (dx+1)/2
	d := min(dx, dy)

	onPix(x0, y0)

	for x, y := x0, y0; x != x1 || y != y1; {
		ax -= d
		ay -= d

		if ax <= 0 {
			x += ix
			ax += dy
		}

		if ay <= 0 {
			y += iy
			ay += dx
		}

		onPix(x, y)
	}
}
