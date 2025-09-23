// Copyright © 2025 Alex Temnok. All rights reserved.

package eda

import (
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

type Components = []*Component

// Component represents vector data for different PCB layers.
// Its data should not be modified after component creation.
// It allows using same components at multiple locations.
type Component struct {
	Transform transform.T

	Back bool

	Cuts      path.Paths
	CutsInner bool
	CutsOuter bool

	Marks path.Paths

	Pads path.Paths

	Tracks      path.Paths
	TracksWidth float64

	ClearWidth    float64
	ClearDisabled bool

	Nested []*Component
}

// Visit calls provided callback for each subcomponent recursively,
// as if every component is isolated (without subcomponents)
func (c *Component) Visit(callback func(*Component)) {
	c.visit(transform.I, &Component{}, callback)
}

func (c *Component) visit(t transform.T, parent *Component, callback func(*Component)) {
	if zero := (transform.T{}); c.Transform != zero {
		t = c.Transform.Multiply(t)
	}

	target := &Component{
		Transform: t,
		Back:      c.Back || parent.Back,

		Cuts:      c.Cuts,
		CutsInner: c.CutsInner || parent.CutsInner,
		CutsOuter: c.CutsOuter || parent.CutsOuter,

		Marks: c.Marks,

		Pads: c.Pads,

		Tracks:      c.Tracks,
		TracksWidth: c.TracksWidth,

		ClearWidth:    c.ClearWidth,
		ClearDisabled: c.ClearDisabled || parent.ClearDisabled,
	}

	if target.TracksWidth == 0 {
		target.TracksWidth = parent.TracksWidth
	}

	if target.ClearWidth == 0 {
		target.ClearWidth = parent.ClearWidth
	}

	callback(target)

	for _, sub := range c.Nested {
		sub.visit(t, target, callback)
	}
}

func (c *Component) PadCenters() []path.Point {
	var centers []path.Point

	c.Visit(func(component *Component) {
		centers = append(centers, component.Pads.Centers(component.Transform)...)
	})

	return centers
}

func (c *Component) Arrange(t transform.T) *Component {
	return &Component{
		Transform: t,
		Nested:    Components{c},
	}
}

func (c *Component) Clone(n int, dx, dy float64) *Component {
	res := &Component{}
	for i := range n {
		k := float64(i) - float64(n-1)/2
		clone := c.Arrange(transform.Move(k*dx, k*dy))
		res.Nested = append(res.Nested, clone)
	}

	return res
}

func ComponentGrid(cols int, dx, dy float64, comps ...*Component) *Component {
	rows := (len(comps) + cols - 1) / cols

	grid := &Component{}

	for i, comp := range comps {
		c := float64(i%cols) - float64(cols-1)/2
		r := float64(i/cols) - float64(rows-1)/2

		grid.Nested = append(grid.Nested, comp.Arrange(transform.Move(c*dx, -r*dy)))
	}

	return grid
}
