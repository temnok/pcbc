// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	Guides34x42 *eda.Component
	Guides72x42 *eda.Component
)

func init() {
	hole := path.Circle(2.1)

	Guides34x42 = &eda.Component{
		CutsFully: true,

		Cuts: path.Paths{
			hole.Transform(transform.Move(-17, 19)),
			hole.Transform(transform.Move(-17, -21)),
			hole.Transform(transform.Move(17, 21)),
			hole.Transform(transform.Move(17, -21)),
		},
	}

	Guides72x42 = &eda.Component{
		CutsFully: true,

		Cuts: path.Paths{
			hole.Transform(transform.Move(-36, 19)),
			hole.Transform(transform.Move(-36, -21)),
			hole.Transform(transform.Move(36, 21)),
			hole.Transform(transform.Move(36, -21)),
		},
	}
}
