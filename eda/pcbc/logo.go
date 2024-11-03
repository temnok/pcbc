package pcbc

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	LogoPaths = append(
		font.CenteredPaths("pc").Apply(transform.Move(0, 0.33)),
		font.CenteredPaths("bc").Apply(transform.Move(0, -0.33))...,
	).Apply(transform.Rotate(45))

	LogoStrokes = path.Strokes{
		font.Bold: LogoPaths,
	}

	Logo = &eda.Component{
		Marks: LogoPaths,
	}

	FirmStrokes = append(
		font.CenteredPaths("TMNK").Apply(transform.Move(0, 0.5)),
		font.CenteredPaths("TECH").Apply(transform.Move(0, -0.5))...,
	)

	Firm = &eda.Component{
		Marks: FirmStrokes,
	}

	TmnkTech = path.Strokes{
		font.Bold: FirmStrokes,
	}
)
