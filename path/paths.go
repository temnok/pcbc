package path

import (
	"temnok/pcbc/transform"
)

// Paths represent a sequence of paths.
type Paths []Path

func (paths Paths) Append(others ...Paths) Paths {
	for _, other := range others {
		paths = append(paths, other...)
	}
	return paths
}

// Apply returns list of transformed paths.
func (paths Paths) Apply(t transform.Transform) Paths {
	res := make(Paths, len(paths))

	for i, path := range paths {
		res[i] = path.Apply(t)
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

func (paths Paths) Bounds() (lt, rb Point) {
	if len(paths) == 0 {
		return
	}

	lt, rb = paths[0].Bounds()
	for _, p := range paths[1:] {
		a, b := p.Bounds()

		lt.X, lt.Y = min(lt.X, a.X), min(lt.Y, a.Y)
		rb.X, rb.Y = max(rb.X, b.X), max(rb.Y, b.Y)
	}

	return
}

func (paths Paths) Resize(delta float64) Paths {
	res := make(Paths, len(paths))

	for i, path := range paths {
		res[i] = path.Resize(delta)
	}

	return res
}
