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
}

// Squash merges component tree into the single component.
func (c *Component) Squash() *Component {
	out := &Component{
		Transform:    geom.Identity(),
		Marks:        path.Strokes{},
		Tracks:       path.Strokes{},
		GroundTracks: path.Strokes{},
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

	out.Clears = append(out.Clears, c.Clears.Transform(t)...)
	out.Cuts = append(out.Cuts, c.Cuts.Transform(t)...)
	out.Holes = append(out.Holes, c.Holes.Transform(t)...)
	out.Pads = append(out.Pads, c.Pads.Transform(t)...)
	out.Openings = append(out.Openings, c.Openings.Transform(t)...)

	out.Marks.Append(c.Marks.Transform(t))
	out.Tracks.Append(c.Tracks.Transform(t))
	out.GroundTracks.Append(c.GroundTracks.Transform(t))
}

func ComponentsGrid(cols, rows int, dx, dy float64, comps ...*Component) Components {
	x0, y0 := -0.5*float64(cols-1)*dx, 0.5*float64(rows-1)*dy

	grid := Components{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			k := i*cols + j
			if k >= len(comps) {
				break
			}

			grid = append(grid, &Component{
				Transform: geom.MoveXY(x0+float64(j)*dx, y0-float64(i)*dy),
				Components: Components{
					comps[k],
				},
			})
		}
	}

	return grid
}
