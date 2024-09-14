package pkg

import (
	"temnok/lab/contour"
	"temnok/lab/geom"
)

func padRow(n int, d, w, h, r float64) [][]geom.XY {
	row := make([][]geom.XY, n)

	totalD := float64(n-1) * d

	for i := range n {
		x := d*float64(i) - totalD*0.5
		row[i] = geom.Move(geom.XY{x, 0}).Points(contour.RoundRect(w, h, r))
	}

	return row
}
