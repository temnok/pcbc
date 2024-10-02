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

	// Mask: solid cut strokes; base cuts without tabs
	ExtraMaskHoles path.Paths

	// Stencil: cuts with tabs
	ExtraStencilCuts path.Paths

	// Stencil: cuts without tabs
	ExtraStencilHoles path.Paths
}

// Merge combines two components into one without changing the originals.
func (c *Component) Merge(d *Component) *Component {
	return &Component{
		Cuts:   c.Cuts.Merge(d.Cuts),
		Holes:  c.Holes.Merge(d.Holes),
		Pads:   c.Pads.Merge(d.Pads),
		Tracks: c.Tracks.Merge(d.Tracks),
		Marks:  c.Marks.Merge(d.Marks),

		ExtraMaskHoles:    c.ExtraMaskHoles.Merge(d.ExtraMaskHoles),
		ExtraStencilCuts:  c.ExtraStencilCuts.Merge(d.ExtraStencilCuts),
		ExtraStencilHoles: c.ExtraStencilHoles.Merge(d.ExtraStencilHoles),
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

		ExtraMaskHoles:    c.ExtraMaskHoles.Transform(t),
		ExtraStencilCuts:  c.ExtraStencilCuts.Transform(t),
		ExtraStencilHoles: c.ExtraStencilHoles.Transform(t),
	}
}
