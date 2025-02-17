// Copyright © 2025 Alex Temnok. All rights reserved.

package pcbc

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Base72x42 *eda.Component

func init() {
	holeContour := path.Circle(2.1)

	hole := &eda.Component{
		Pads:          holeContour,
		Holes:         holeContour,
		MaskbaseHoles: holeContour,
	}

	Base72x42 = &eda.Component{
		Components: eda.Components{
			hole.Arrange(transform.Move(-36, 19)),
			hole.Arrange(transform.Move(-36, -21)),
			hole.Arrange(transform.Move(36, 21)),
			hole.Arrange(transform.Move(36, -21)),
		},
	}
}
