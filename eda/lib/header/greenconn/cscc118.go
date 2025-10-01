// Copyright © 2025 Alex Temnok. All rights reserved.

package greenconn

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

func CSCC118(n int, flip bool, labels []string) *eda.Component {
	const padW = 1.95
	pad := path.RoundRect(padW, 0.5, 0.2)

	shift := float64((n+1)%2) / 2

	sign, align0, align1 := 1.0, font.AlignRight, font.AlignLeft
	if flip {
		sign, align0, align1 = -sign, align1, align0
	}

	pads := pad.Clone(n, 0, -1)
	labels0, labels1 := make([]string, (n+1)/2), make([]string, n/2)
	for i := range n {
		if i%2 == 0 {
			if i < len(labels) {
				labels0[i/2] = labels[i]
			}
			pads[i] = pads[i].Transform(transform.Move(sign*padW/2, 0))
		} else {
			if i < len(labels) {
				labels1[i/2] = labels[i]
			}
			pads[i] = pads[i].Transform(transform.Move(-sign*padW/2, 0))
		}
	}

	return &eda.Component{
		Pads: pads,

		Nested: eda.Components{},

		Marks: path.Join(
			path.Paths{path.Rect(1.5, float64(n)+0.5)},

			font.AlignedColumn(align0, -2/1.6, labels0...).
				Transform(transform.Scale(0.8, 1.6).Move(-sign*0.9, shift)),
			font.AlignedColumn(align1, -2/1.6, labels1...).
				Transform(transform.Scale(0.8, 1.6).Move(sign*0.9, -shift)),
		),
	}
}
