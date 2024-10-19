package greenconn

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

func CSCC118(n int, labels []string) *lib.Component {
	const padW = 1.95
	pad := path.Rect(padW, 0.5)

	shift := float64((n+1)%2) / 2

	pads := pad.Clone(n, 0, -1)
	labels0, labels1 := make([]string, (n+1)/2), make([]string, n/2)
	for i := range n {
		if i%2 == 0 {
			if i < len(labels) {
				labels0[i/2] = labels[i]
			}
			pads[i] = pads[i].Apply(transform.Move(padW/2, 0))
		} else {
			if i < len(labels) {
				labels1[i/2] = labels[i]
			}
			pads[i] = pads[i].Apply(transform.Move(-padW/2, 0))
		}
	}

	return &lib.Component{
		Pads: pads,

		Marks: path.Strokes{
			0.1: path.Paths{path.Rect(1.5, float64(n)+0.5)},
		}.Append(
			path.Strokes{
				0.15: path.Paths{}.Append(
					font.StringsPaths(labels0, font.AlignRight, geom.XY{0, -2 / 1.6}).
						Apply(transform.Scale(0.8, 1.6).Move(-0.9, shift)),
					font.StringsPaths(labels1, font.AlignLeft, geom.XY{0, -2 / 1.6}).
						Apply(transform.Scale(0.8, 1.6).Move(0.9, -shift)),
				),
			},
		),
	}
}
