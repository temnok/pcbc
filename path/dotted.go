package path

import "temnok/lab/geom"

func IterateDotted(path []geom.XY, transform geom.Transform, d int, iterator func(x, y int)) {
	Iterate(path, transform, wrapDotted(d, iterator))
}

func wrapDotted(d int, visit func(x, y int)) func(x, y int) {
	var px, py int
	started := false

	return func(x, y int) {
		if !started {
			started = true
		} else {
			dx, dy := x-px, y-py
			if dx*dx+dy*dy < d*d {
				return
			}
		}

		visit(x, y)
		px, py = x, y
	}
}
