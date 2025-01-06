// Copyright © 2025 Alex Temnok. All rights reserved.

package pcbc

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	Logo = &eda.Component{
		Marks: path.Join(
			font.Centered("pc").Apply(transform.Move(0, 0.33)),
			font.Centered("bc").Apply(transform.Move(0, -0.33)),
		).Apply(transform.Rotate(45)),
	}

	Firm = eda.CenteredTextColumn(-1,
		"TMNK",
		"TECH",
	)
)
