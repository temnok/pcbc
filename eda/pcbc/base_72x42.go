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
		Pads:  holeContour,
		Holes: holeContour,
	}

	key := path.Circle(0.6).Apply(transform.Move(-37.25, 22.25))

	Base72x42 = &eda.Component{
		Pads: key,

		Components: eda.Components{
			hole.Clone(2, 72, 0).Clone(2, 0, 42),
		},
	}
}
