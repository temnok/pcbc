package pcbc

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/font"
	"temnok/pcbc/transform"
)

var (
	Logo = &eda.Component{
		Marks: append(
			font.CenteredPaths("pc").Apply(transform.Move(0, 0.33)),
			font.CenteredPaths("bc").Apply(transform.Move(0, -0.33))...,
		).Apply(transform.Rotate(45)),
	}

	Firm = eda.CenteredText(
		"TMNK",
		"TECH",
	)
)
