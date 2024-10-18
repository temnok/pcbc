package pcbc

import (
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Logo = path.Strokes{
	font.Bold: append(
		font.StringPaths("pc", font.AlignCenter).Apply(transform.Move(0, 0.33)),
		font.StringPaths("bc", font.AlignCenter).Apply(transform.Move(0, -0.33))...,
	).Apply(transform.Rotate(45)),
}

var TmnkTech = path.Strokes{}.Append(
	font.CenterBold("TMNK").Apply(transform.Scale(0.75, 0.5).Move(0, 0.3)),
	font.CenterBold("TECH").Apply(transform.Scale(0.75, 0.5).Move(0, -0.3)),
)
