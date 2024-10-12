package smd

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/path"
)

// https://i.ebayimg.com/images/g/ufQAAOSwazpdFvUO/s-l1600.webp

var (
	I0201 = M0603
	M0603 = &lib.Component{
		Pads: path.Rect(0.3, 0.3).Clone(2, 0.6, 0),
	}

	I0402 = M1005
	M1005 = &lib.Component{
		Pads: path.Rect(0.5, 0.6).Clone(2, 1, 0),
	}

	I0603 = M1608
	M1608 = &lib.Component{
		Pads: path.Rect(0.6, 0.9).Clone(2, 1.5, 0),
	}

	I0805 = M2012
	M2012 = &lib.Component{
		Pads: path.Rect(0.7, 1.3).Clone(2, 1.9, 0),
	}
)
