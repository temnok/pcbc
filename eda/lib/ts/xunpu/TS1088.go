// Copyright © 2025 Alex Temnok. All rights reserved.

package xunpu

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/path"
)

var SwitchTS1088 = &eda.Component{
	Pads: path.Rect(1.35, 1.8).
		CloneXY(2, 4.15, 0),

	Marks: path.Paths{
		path.Rect(4, 2.9),
	},
}
