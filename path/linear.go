package path

// linearVisit iterates over all line pixels, except the first one.
func linearVisit(a, b Point, iterate func(x, y int)) {
	x0, y0 := a.RoundXY()
	x1, y1 := b.RoundXY()

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

	//iterate(x0, y0)

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

		iterate(x, y)
	}
}
