package lib

import (
	"temnok/lab/geom"
	"temnok/lab/path"
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

	// FR4: copper strokes
	Tracks path.Strokes

	// Mask: solid cut strokes
	Openings path.Paths

	// Mask: solid mark strokes
	Marks path.Strokes

	// Stencil: cuts with tabs
	StencilCuts path.Paths

	// Mask: extra-wide solid cut strokes
	// Mask: base cuts without tabs
	MaskBaseHoles path.Paths
}
