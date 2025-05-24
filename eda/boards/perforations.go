// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	Perforations34x42 *eda.Component
	Perforations72x42 *eda.Component
)

func init() {
	perforation := path.Circle(2.1)

	Perforations34x42 = &eda.Component{
		Perforations: path.Paths{
			perforation.Transform(transform.Move(-17, 19)),
			perforation.Transform(transform.Move(-17, -21)),
			perforation.Transform(transform.Move(17, 21)),
			perforation.Transform(transform.Move(17, -21)),
		},
	}

	Perforations72x42 = &eda.Component{
		Perforations: path.Paths{
			perforation.Transform(transform.Move(-36, 19)),
			perforation.Transform(transform.Move(-36, -21)),
			perforation.Transform(transform.Move(36, 21)),
			perforation.Transform(transform.Move(36, -21)),
		},
	}
}
