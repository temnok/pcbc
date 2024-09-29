package qfn

import (
	"temnok/lab/geom"
	"temnok/lab/path"
)

func QFN16G() path.Paths {
	pad := path.RoundRect(0.6, 0.25, 0.12)

	col := pad.Clone(4, 0, -0.5).Transform(geom.MoveXY(-1.5, 0))

	pads := path.Paths{}
	for a := 0.0; a < 360; a += 90 {
		pads = append(pads, col.Transform(geom.RotateD(a))...)
	}

	return append(pads, keyedRect(1.6, 1.6, 0.35))
}

func keyedRect(w, h, k float64) []geom.XY {
	x, y := w/2, h/2
	return path.Lines([]geom.XY{{-x + k, y}, {x, y}, {x, -y}, {-x, -y}, {-x, y - k}, {-x + k, y}})
}
