// Copyright © 2025 Alex Temnok. All rights reserved.

package hyp

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Switch1TS026A = &eda.Component{
	Pads: path.Rect(0.5, 0.55).
		Clone(2, 1.3-0.5, 0).
		Clone(2, 0, -(3.2 - 0.55)).
		Transform(transform.Scale(1, -1).RotateDegrees(90)),

	Marks: path.Paths{
		path.Rect(2.6, 1.6),
	},
}
