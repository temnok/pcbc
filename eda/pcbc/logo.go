package pcbc

import (
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	LogoPaths = append(
		font.CenteredStringPaths("pc").Apply(transform.Move(0, 0.33)),
		font.CenteredStringPaths("bc").Apply(transform.Move(0, -0.33))...,
	).Apply(transform.Rotate(45))

	Logo = path.Strokes{
		font.Bold: LogoPaths,
	}
)
