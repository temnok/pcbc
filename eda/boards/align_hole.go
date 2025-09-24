// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var AlignHole = &eda.Component{
	CutsFully: true,

	Cuts: path.Paths{
		path.Circle(1.45),
	},
}
