// Copyright © 2025 Alex Temnok. All rights reserved.

package smd

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

// https://i.ebayimg.com/images/g/ufQAAOSwazpdFvUO/s-l1600.webp

var (
	I0201 = M0603
	M0603 = &eda.Component{
		Pads: path.Rect(0.25, 0.35).CloneXY(2, 0.55, 0),
		Marks: path.Paths{
			path.Rect(1.0, 0.55),
		},
	}

	I0402 = M1005
	M1005 = &eda.Component{
		Pads: path.Rect(0.5, 0.6).CloneXY(2, 1, 0),
		Marks: path.Paths{
			path.Rect(1.7, 0.8),
		},
	}

	I0603 = M1608
	M1608 = &eda.Component{
		Pads: path.Rect(0.6, 0.9).CloneXY(2, 1.5, 0),
		Marks: path.Paths{
			path.Rect(1.8, 1.0),
		},
	}
)
