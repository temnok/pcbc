// Copyright © 2025 Alex Temnok. All rights reserved.

package smd

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

// https://i.ebayimg.com/images/g/ufQAAOSwazpdFvUO/s-l1600.webp

var (
	M0603 = &eda.Component{
		Pads: path.Rect(0.3, 0.3).Clone(2, 0.6, 0),
	}

	M1005 = &eda.Component{
		Pads: path.Rect(0.5, 0.6).Clone(2, 1, 0),
	}

	I0603 = M1608
	M1608 = &eda.Component{
		Pads: path.Rect(0.6, 0.9).Clone(2, 1.5, 0),
	}

	M2012 = &eda.Component{
		Pads: path.Rect(0.7, 1.3).Clone(2, 1.9, 0),
	}
)
