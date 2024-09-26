package qfn16

import (
	"temnok/lab/contour"
	"temnok/lab/eda"
	"temnok/lab/geom"
	"temnok/lab/path"
)

var PadCenters = []geom.XY{
	1:  {-1.5, 0.75},
	2:  {-1.5, 0.25},
	3:  {-1.5, -0.25},
	4:  {-1.5, -0.75},
	5:  {-0.75, -1.5},
	6:  {-0.25, -1.5},
	7:  {0.25, -1.5},
	8:  {0.75, -1.5},
	9:  {1.5, -0.75},
	10: {1.5, -0.25},
	11: {1.5, 0.25},
	12: {1.5, 0.75},
	13: {0.75, 1.5},
	14: {0.25, 1.5},
	15: {-0.25, 1.5},
	16: {-0.75, 1.5},
	17: {0, 0},
}[1:]

var PadContours = path.Paths{
	1:  hPad,
	2:  hPad,
	3:  hPad,
	4:  hPad,
	5:  vPad,
	6:  vPad,
	7:  vPad,
	8:  vPad,
	9:  hPad,
	10: hPad,
	11: hPad,
	12: hPad,
	13: vPad,
	14: vPad,
	15: vPad,
	16: vPad,
	//17: contour.Rect(1.6, 1.6),
	17: keyedRect(1.6, 1.6, 0.35),
}[1:]

var (
	hPad = contour.RoundRect(0.6, 0.25, 0.12)
	vPad = contour.RoundRect(0.25, 0.6, 0.12)
)

func init() {
	for i, xy := range PadCenters {
		PadContours[i] = PadContours[i].Transform(geom.Move(xy))
	}
}

func keyedRect(w, h, k float64) []geom.XY {
	x, y := w/2, h/2
	return contour.Lines([]geom.XY{{-x + k, y}, {x, y}, {x, -y}, {-x, -y}, {-x, y - k}, {-x + k, y}})
}

func Add(pcb *eda.PCB, t geom.Transform) path.Path {
	pcb.Pad(PadContours.Transform(t)...)
	//pcb.SilkContour(t, 0.1, contour.Rect(3, 3))
	pcb.SilkText(t.MoveXY(-2.3, 0.8), 0.6, "1")

	return PadCenters
}
