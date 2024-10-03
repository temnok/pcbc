package lib

import (
	"temnok/lab/geom"
	"temnok/lab/path"
)

// Component represents vector data for different PCB layers.
type Component struct {
	// FR4: cuts with tabs
	// Mask: dotted cut strokes
	Cuts path.Paths

	// FR4: cuts without tabs
	Holes path.Paths

	// FR4: copper fills
	// Stencil: cuts without tabs
	Pads path.Paths

	// FR4: copper strokes
	Tracks path.Strokes

	// Mask: solid cut strokes
	Openings path.Paths

	// Mask: solid mark strokes
	Marks path.Strokes

	// Mask: extra-wide solid cut strokes
	// Mask: base cuts without tabs
	MaskBaseHoles path.Paths
}

// Merge combines two components into one without changing the originals.
func (c *Component) Merge(d *Component) *Component {
	return &Component{
		Cuts:          c.Cuts.Merge(d.Cuts),
		Holes:         c.Holes.Merge(d.Holes),
		Pads:          c.Pads.Merge(d.Pads),
		Tracks:        c.Tracks.Merge(d.Tracks),
		Openings:      c.Openings.Merge(d.Openings),
		Marks:         c.Marks.Merge(d.Marks),
		MaskBaseHoles: c.MaskBaseHoles.Merge(d.MaskBaseHoles),
	}
}

// Transform returns transformed copy of a component.
func (c *Component) Transform(t geom.Transform) *Component {
	return &Component{
		Cuts:          c.Cuts.Transform(t),
		Holes:         c.Holes.Transform(t),
		Pads:          c.Pads.Transform(t),
		Tracks:        c.Tracks.Transform(t),
		Openings:      c.Openings.Transform(t),
		Marks:         c.Marks.Transform(t),
		MaskBaseHoles: c.MaskBaseHoles.Transform(t),
	}
}
