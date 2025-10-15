// Copyright © 2025 Alex Temnok. All rights reserved.

package path

import (
	"github.com/temnok/pcbc/transform"
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

// Transform returns list of transformed paths.
func (paths Paths) Transform(t transform.T) Paths {
	res := make(Paths, len(paths))

	for i, path := range paths {
		res[i] = path.Transform(t)
	}

	return res
}

// Rasterize calls provided callback for each path.
func (paths Paths) Rasterize(t transform.T, callback func(x, y int)) {
	for _, path := range paths {
		path.Rasterize(t, callback)
	}
}

// RasterizeIntermittently calls provided callback for each path.
func (paths Paths) RasterizeIntermittently(t transform.T, dist float64, callback func(x, y int)) {
	for _, path := range paths {
		path.RasterizeIntermittently(t, dist, callback)
	}
}

func (paths Paths) Clone(n int, dx, dy float64) Paths {
	res := make(Paths, 0, n*len(paths))

	for i := 0; i < n; i++ {
		k := float64(i) - float64(n-1)/2
		res = append(res, paths.Transform(transform.Move(dx*k, dy*k))...)
	}

	return res
}

// Centers returns center points for paths.
func (paths Paths) Centers(t transform.T) []Point {
	centers := make([]Point, len(paths))

	for i, path := range paths {
		centers[i] = path.Center(t)
	}

	return centers
}
