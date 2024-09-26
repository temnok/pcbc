package pad

import (
	"temnok/lab/eda"
	"temnok/lab/geom"
	"temnok/lab/path"
)

func Row(pcb *eda.PCB, transform geom.Transform, padContour path.Path, n int, d, jump float64) path.Path {
	if n <= 0 || d <= 0 {
		return nil
	}

	centers := make([]geom.XY, n)

	for i := 0; i < n; i++ {
		x := float64(d) * (float64(i) - float64(n-1)/2)

		y := jump
		if i%2 != 0 {
			y = -jump
		}
		centers[i] = geom.XY{x, y}

		pcb.Pad(padContour.Transform(transform.MoveXY(x, y)))
	}

	return centers
}
