// Copyright © 2025 Alex Temnok. All rights reserved.

package path

import (
	"temnok/pcbc/transform"
)

// Paths represent a sequence of paths.
type Paths []Path

func Join(paths ...Paths) Paths {
	var res Paths
	for _, p := range paths {
		res = append(res, p...)
	}
	return res
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

func (paths Paths) Clone(n int, dx, dy float64) Paths {
	res := make(Paths, 0, n*len(paths))

	for i := 0; i < n; i++ {
		k := float64(i) - float64(n-1)/2
		res = append(res, paths.Apply(transform.Move(dx*k, dy*k))...)
	}

	return res
}

// Centers returns center points for paths.
func (paths Paths) Centers(t transform.Transform) Points {
	centers := make(Points, len(paths))

	for i, path := range paths {
		centers[i] = path.Center(t)
	}

	return centers
}
