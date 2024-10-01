package qfn

import (
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/path"
)

func QFN16G() *lib.Component {
	pad := path.RoundRect(0.7, 0.3, 0.12)

	col := pad.Clone(4, 0, -0.5).Transform(geom.MoveXY(-1.5, 0))

	pads := path.Paths{}
	for a := 0.0; a < 360; a += 90 {
		pads = append(pads, col.Transform(geom.RotateD(a))...)
	}

	for x := -0.4; x <= 0.4; x += 0.8 {
		for y := -0.4; y <= 0.4; y += 0.8 {
			pads = append(pads, path.Rect(0.6, 0.6).Transform(geom.MoveXY(x, y)))
		}
	}

	//pads = append(pads, keyedRect(1.6, 1.6, 0.35))

	return &lib.Component{
		Pads:   pads,
		Placer: path.Rect(3.1, 3.1),
	}
}

func keyedRect(w, h, k float64) []geom.XY {
	x, y := w/2, h/2
	return path.Lines([]geom.XY{{-x + k, y}, {x, y}, {x, -y}, {-x, -y}, {-x, y - k}, {-x + k, y}})
}
