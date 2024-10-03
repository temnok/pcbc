package path

import "temnok/lab/geom"

// Paths represent a sequence of paths.
type Paths []Path

func (paths Paths) Merge(others ...Paths) Paths {
	res := make(Paths, 0, len(paths)+len(others))

	res = append(res, paths...)
	for _, other := range others {
		res = append(res, other...)
	}

	return res
}

// Transform returns list of transformed paths.
func (paths Paths) Transform(transform geom.Transform) Paths {
	res := make(Paths, len(paths))

	for i, path := range paths {
		res[i] = path.Transform(transform)
	}

	return res
}

// Visit calls provided callback for each path.
func (paths Paths) Visit(visit func(x, y int)) {
	for _, path := range paths {
		path.Visit(visit)
	}
}

// Jump calls provided callback for each path.
func (paths Paths) Jump(dist int, jump func(x, y int)) {
	for _, path := range paths {
		path.Jump(dist, jump)
	}
}

// Centers returns center points for paths.
func (paths Paths) Centers() Points {
	centers := make(Points, len(paths))

	for i, path := range paths {
		centers[i] = path.Center()
	}

	return centers
}
