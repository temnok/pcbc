// Copyright © 2025 Alex Temnok. All rights reserved.

package path

import "github.com/temnok/pcbc/transform"

type Bounds struct {
	initialized bool
	lb, rt      Point
}

func (b *Bounds) includePoint(p Point) {
	if !b.initialized {
		b.initialized = true
		b.lb, b.rt = p, p
		return
	}

	b.lb.X, b.lb.Y = min(b.lb.X, p.X), min(b.lb.Y, p.Y)
	b.rt.X, b.rt.Y = max(b.rt.X, p.X), max(b.rt.Y, p.Y)
}

func (b *Bounds) AddPath(t transform.T, path Path) {
	for _, p := range path {
		b.includePoint(p.Apply(t))
	}
}

func (b *Bounds) AddPaths(t transform.T, paths Paths) {
	for _, p := range paths {
		b.AddPath(t, p)
	}
}

func (b *Bounds) Center() Point {
	return Point{(b.lb.X + b.rt.X) / 2, (b.lb.Y + b.rt.Y) / 2}
}
