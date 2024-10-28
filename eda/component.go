package eda

import (
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

type Components = []*Component

// Component represents vector data for different PCB layers.
type Component struct {
	Transform transform.Transform

	// FR4: remove groundfill
	Clears path.Paths

	// FR4: cuts with tabs
	// Mask: dotted cut strokes
	Cuts path.Paths

	// FR4: cuts without tabs
	// Mask: solid cut strokes
	Holes path.Paths

	// FR4: copper fills
	// Mask: solid cut strokes
	// Stencil: cuts without tabs
	Pads path.Paths

	// Mask: solid cut strokes
	Openings path.Paths

	// Mask: solid mark strokes
	Marks path.Strokes

	// FR4: copper strokes with groundfill clean
	Tracks path.Strokes

	// FR4: copper strokes without groundfill clean
	GroundTracks path.Strokes

	Components []*Component
}

// Visit calls provided callback for each subcomponent recursively,
// as if every component is isolated (without subcomponents)
func (c *Component) Visit(callback func(*Component)) {
	c.visit(transform.Identity, callback)
}

func (c *Component) visit(t transform.Transform, callback func(*Component)) {
	if c.Transform != (transform.Transform{}) {
		t = c.Transform.Multiply(t)
	}

	callback(&Component{
		Transform:    t,
		Clears:       c.Clears,
		Cuts:         c.Cuts,
		Holes:        c.Holes,
		Pads:         c.Pads,
		Openings:     c.Openings,
		Marks:        c.Marks,
		Tracks:       c.Tracks,
		GroundTracks: c.GroundTracks,
	})

	for _, sub := range c.Components {
		sub.visit(t, callback)
	}
}

func (c *Component) PadCenters() path.Points {
	var centers path.Points

	c.visit(transform.Identity, func(component *Component) {
		for _, pad := range component.Pads {
			centers = append(centers, pad.Center(component.Transform))
		}
	})

	return centers
}
func (c *Component) Arrange(t transform.Transform) *Component {
	return &Component{
		Transform:  t,
		Components: Components{c},
	}
}

func (c *Component) Clone(n int, dx, dy float64) *Component {
	res := &Component{}
	for i := range n {
		k := (float64(i) - float64(n-1)/2)
		clone := c.Arrange(transform.Move(k*dx, k*dy))
		res.Components = append(res.Components, clone)
	}

	return res
}

func (c *Component) Size() (float64, float64) {
	var b path.Bounds

	c.Visit(func(c *Component) {
		b.AddPaths(c.Transform, c.Clears)
		b.AddPaths(c.Transform, c.Cuts)
		b.AddPaths(c.Transform, c.Holes)
		b.AddPaths(c.Transform, c.Pads)
		b.AddPaths(c.Transform, c.Openings)
		b.AddStrokes(c.Transform, c.Marks)
		b.AddStrokes(c.Transform, c.Tracks)
		b.AddStrokes(c.Transform, c.GroundTracks)
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
