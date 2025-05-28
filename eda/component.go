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

	// Remove: copper strokes
	// Cut: substrate (leave tabs)
	Cuts path.Paths

	// Remove: copper strokes, mask strokes
	// Add: copper shapes
	// Cut: stencil
	Pads path.Paths

	// Remove: copper strokes
	// Add: copper strokes
	Tracks path.Paths

	// Add: copper strokes
	GroundTracks path.Paths

	// Add: mark strokes
	Marks path.Paths

	TrackWidth float64

	NoClear bool

	Components []*Component
}

// Visit calls provided callback for each subcomponent recursively,
// as if every component is isolated (without subcomponents)
func (c *Component) Visit(callback func(*Component)) {
	c.visit(transform.I, &Component{}, callback)
}

func (c *Component) visit(t transform.T, parent *Component, callback func(*Component)) {
	if c.Transform != (transform.T{}) {
		t = c.Transform.Multiply(t)
	}

	comp := &Component{
		Transform:    t,
		Cuts:         c.Cuts,
		Pads:         c.Pads,
		Tracks:       c.Tracks,
		GroundTracks: c.GroundTracks,
		Marks:        c.Marks,
		TrackWidth:   c.TrackWidth,
		NoClear:      c.NoClear || parent.NoClear,
	}

	if comp.TrackWidth == 0 {
		comp.TrackWidth = parent.TrackWidth
	}

	callback(comp)

	for _, sub := range c.Components {
		sub.visit(t, comp, callback)
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
		Transform:  t,
		Components: Components{c},
	}
}

func (c *Component) Clone(n int, dx, dy float64) *Component {
	res := &Component{}
	for i := range n {
		k := float64(i) - float64(n-1)/2
		clone := c.Arrange(transform.Move(k*dx, k*dy))
		res.Components = append(res.Components, clone)
	}

	return res
}

func (c *Component) Size() (float64, float64) {
	var b path.Bounds

	c.Visit(func(c *Component) {
		b.AddPaths(c.Transform, c.Cuts)
		b.AddPaths(c.Transform, c.Pads)
		b.AddPaths(c.Transform, c.Tracks)
		b.AddPaths(c.Transform, c.GroundTracks)
		b.AddPaths(c.Transform, c.Marks)
	})

	return b.Width(), b.Height()
}

func ComponentGrid(cols int, dx, dy float64, comps ...*Component) *Component {
	rows := (len(comps) + cols - 1) / cols

	grid := &Component{}

	for i, comp := range comps {
		c := float64(i%cols) - float64(cols-1)/2
		r := float64(i/cols) - float64(rows-1)/2

		grid.Components = append(grid.Components, comp.Arrange(transform.Move(c*dx, -r*dy)))
	}

	return grid
}
