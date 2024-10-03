package qfn

import (
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/path"
)

func QFN16G() *lib.Component {
	pad := path.RoundRect(0.6, 0.25, 0.12)

	col := pad.Clone(4, 0, -0.5).Transform(geom.MoveXY(-1.5, 0))

	pads := path.Paths{}
	for a := 0.0; a < 360; a += 90 {
		pads = append(pads, col.Transform(geom.RotateD(a))...)
	}

	const g = 0.6
	tracks := path.Paths{
		path.Lines([]geom.XY{{-g, -g}, {g, -g}, {g, g}, {-g, g}, {-g, -g}}),
	}

	for x := -g; x <= g; x += g {
		for y := -g; y <= g; y += g {
			pads = append(pads, path.Rect(0.35, 0.35).Transform(geom.MoveXY(x, y)))
		}
	}

	return &lib.Component{
		Pads:     pads,
		Tracks:   path.Strokes{0.2: tracks},
		Openings: pads,
	}
}
