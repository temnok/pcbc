package path

import (
	"temnok/lab/geom"
)

func Iterate(path []geom.XY, transform geom.Transform, iterator func(x, y int)) {
	if len(path) == 0 {
		return
	}

	a := transform.Point(path[0])
	iterator(a.Ints())

	for i := 0; i+3 < len(path); i += 3 {
		c1, c2, b := transform.Point(path[i+1]), transform.Point(path[i+2]), transform.Point(path[i+3])

		if a == c1 && c2 == b {
			linearIterate(a, b, iterator)
		} else {
			cubicIterate([]geom.XY{a, c1, c2, b}, iterator)
		}

		a = b
	}
}

func IterateAll(paths [][]geom.XY, trans geom.Transform, iterator func(x, y int)) {
	for _, path := range paths {
		Iterate(path, trans, iterator)
	}
}
