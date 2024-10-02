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
	// Mask: solid cut strokes
	Holes path.Paths

	// FR4: copper
	// Mask: solid cut strokes
	// Stencil: cuts without tabs
	Pads path.Paths

	// FR4: copper
	Tracks path.Strokes

	// Mask: solid mark strokes
	Marks path.Strokes

	// FR4: cuts with tabs
	// Mask: dotted cut strokes
	// Stencil: cuts with tabs
	GlobalCuts path.Paths

	// FR4: cuts without tabs
	// Mask: solid cut strokes
	// Stencil: cuts without tabs
	GlobalHoles path.Paths

	// FR4: copper
	// Mask: solid mark strokes
	// Stencil: solid mark strokes
	GlobalMarks path.Strokes

	// Mask: extra-wide solid cut strokes
	// Mask: base cuts without tabs
	MaskBaseHoles path.Paths
}

// Merge combines two components into one without changing the originals.
func (c *Component) Merge(d *Component) *Component {
	return &Component{
		Cuts:   c.Cuts.Merge(d.Cuts),
		Holes:  c.Holes.Merge(d.Holes),
		Pads:   c.Pads.Merge(d.Pads),
		Tracks: c.Tracks.Merge(d.Tracks),
		Marks:  c.Marks.Merge(d.Marks),

		GlobalCuts:  c.GlobalCuts.Merge(d.GlobalCuts),
		GlobalHoles: c.GlobalHoles.Merge(d.GlobalHoles),
		GlobalMarks: c.GlobalMarks.Merge(d.GlobalMarks),

		MaskBaseHoles: c.MaskBaseHoles.Merge(d.MaskBaseHoles),
	}
}

// Transform returns transformed copy of a component.
func (c *Component) Transform(t geom.Transform) *Component {
	return &Component{
		Cuts:   c.Cuts.Transform(t),
		Holes:  c.Holes.Transform(t),
		Pads:   c.Pads.Transform(t),
		Tracks: c.Tracks.Transform(t),
		Marks:  c.Marks.Transform(t),

		MaskBaseHoles: c.MaskBaseHoles.Transform(t),
	}
}
