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

	AlignCuts, AlignHiddenCuts path.Paths

	Cuts      path.Paths
	CutsWidth float64

	Marks      path.Paths
	MarksWidth float64

	Pads path.Paths

	Tracks      path.Paths
	TracksWidth float64

	ClearWidth float64

	Nested []*Component
}

const (
	ClearOff   = -1
	CutsHidden = -1
)

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

		Back: c.Back,

		AlignCuts:       c.AlignCuts,
		AlignHiddenCuts: c.AlignHiddenCuts,

		Cuts:      c.Cuts,
		CutsWidth: firstNonZero(c.CutsWidth, parent.CutsWidth),

		Marks:      c.Marks,
		MarksWidth: firstNonZero(c.MarksWidth, parent.MarksWidth),

		Pads: c.Pads,

		Tracks:      c.Tracks,
		TracksWidth: firstNonZero(c.TracksWidth, parent.TracksWidth),

		ClearWidth: firstNonZero(c.ClearWidth, parent.ClearWidth),
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

func (c *Component) Clone(n int, dT transform.T) *Component {
	ans := &Component{}

	t := transform.I
	for range n {
		ans.Nested = append(ans.Nested, c.Arrange(t))
		t = t.Multiply(dT)
	}

	return ans
}

func (c *Component) CloneX(n int, dx float64) *Component {
	return c.CloneXY(n, dx, 0)
}

func (c *Component) CloneY(n int, dy float64) *Component {
	return c.CloneXY(n, 0, dy)
}

func (c *Component) CloneXY(n int, dx, dy float64) *Component {
	k := -float64(n-1) / 2
	return c.Arrange(transform.Move(k*dx, k*dy)).Clone(n, transform.Move(dx, dy))
}

func (c *Component) ClearOff() bool {
	return c.ClearWidth <= 0
}

func (c *Component) CutsHidden() bool {
	return c.CutsWidth <= 0
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

func firstNonZero(a, b float64) float64 {
	if a != 0 {
		return a
	}

	return b
}
