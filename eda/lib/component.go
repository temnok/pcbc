package lib

import (
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

type Components = []*Component

// Component represents vector data for different PCB layers.
type Component struct {
	Description string

	Transform geom.Transform

	Components []*Component

	// FR4: cuts with tabs
	// Mask: dotted cut strokes
	Cuts path.Paths

	// FR4: cuts without tabs
	Holes path.Paths

	// FR4: copper fills
	// Stencil: cuts without tabs
	Pads path.Paths

	// Mask: solid cut strokes
	Openings path.Paths

	// FR4: copper strokes
	Tracks path.Strokes

	// Mask: solid mark strokes
	Marks path.Strokes
}

// Squash merges component tree into the single component.
func (c *Component) Squash() *Component {
	out := &Component{
		Transform: geom.Identity(),
		Tracks:    path.Strokes{},
		Marks:     path.Strokes{},
	}

	c.dump(geom.Identity(), out)

	return out
}

func (c *Component) dump(t geom.Transform, out *Component) {
	if !c.Transform.IsZero() {
		t = t.Multiply(c.Transform)
	}

	for _, sub := range c.Components {
		sub.dump(t, out)
	}

	out.Cuts = append(out.Cuts, c.Cuts.Transform(t)...)
	out.Holes = append(out.Holes, c.Holes.Transform(t)...)
	out.Pads = append(out.Pads, c.Pads.Transform(t)...)
	out.Openings = append(out.Openings, c.Openings.Transform(t)...)

	out.Tracks.Append(c.Tracks.Transform(t))
	out.Marks.Append(c.Marks.Transform(t))
}
