// Copyright © 2025 Alex Temnok. All rights reserved.

package hyp

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Switch1TS026A = &eda.Component{
	Pads: path.RoundRect(0.5, 0.55, 0.15).
		CloneXY(2, 1.3-0.5, 0).
		Clone(2, 0, -(3.2 - 0.55)).
		Transform(transform.MirrorY.RotateDegrees(90)),

	Marks: path.Paths{
		path.Rect(3.5, 1.6),
	},
}

func init() {
	pads := Switch1TS026A.PadCenters()

	Switch1TS026A.Tracks = path.Paths{
		eda.LinearTrack(pads[0], pads[1]),
		eda.LinearTrack(pads[2], pads[3]),
	}
}
