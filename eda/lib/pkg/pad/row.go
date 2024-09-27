package pad

import (
	"temnok/lab/eda"
	"temnok/lab/geom"
	"temnok/lab/path"
)

func Row(pcb *eda.PCB, transform geom.Transform, padContour, silkContour path.Path, n int, d float64) path.Path {
	if n <= 0 || d <= 0 {
		return nil
	}

	centers := make([]geom.XY, n)

	for i := 0; i < n; i++ {
		x := float64(d) * (float64(i) - float64(n-1)/2)
		y := 0.0
		centers[i] = geom.XY{x, y}

		t := transform.MoveXY(x, y)
		pcb.Pad(padContour.Transform(t))
		pcb.SilkContour(0.1, silkContour.Transform(transform.MoveXY(x, y)))
	}

	return centers
}
