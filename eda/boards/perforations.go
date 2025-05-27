// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	Holes34x42 *eda.Component
	Holes72x42 *eda.Component
)

func init() {
	hole := path.Circle(2.1)

	Holes34x42 = &eda.Component{
		Cuts: path.Paths{
			hole.Transform(transform.Move(-17, 19)),
			hole.Transform(transform.Move(-17, -21)),
			hole.Transform(transform.Move(17, 21)),
			hole.Transform(transform.Move(17, -21)),
		},
	}

	Holes72x42 = &eda.Component{
		Cuts: path.Paths{
			hole.Transform(transform.Move(-36, 19)),
			hole.Transform(transform.Move(-36, -21)),
			hole.Transform(transform.Move(36, 21)),
			hole.Transform(transform.Move(36, -21)),
		},
	}
}
