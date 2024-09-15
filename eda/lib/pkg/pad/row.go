package pad

import (
	"temnok/lab/eda"
	"temnok/lab/geom"
)

func Row(pcb *eda.PCB, transform geom.Transform, padContour []geom.XY, n int, d float64) []geom.XY {
	if n <= 0 || d <= 0 {
		return nil
	}

	centers := make([]geom.XY, n)

	for i := 0; i < n; i++ {
		x := float64(d) * (float64(i) - float64(n-1)/2)
		centers[i] = transform.PointXY(x, 0)

		pcb.Pad(transform.MoveXY(x, 0), padContour)
	}

	return centers
}
