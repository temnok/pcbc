package pcbc

import (
	"temnok/lab/font"
	"temnok/lab/geom"
	"temnok/lab/path"
)

var Logo = path.Strokes{
	font.Bold: append(
		font.StringPaths("pc", font.AlignCenter).Transform(geom.MoveXY(0, 0.33)),
		font.StringPaths("bc", font.AlignCenter).Transform(geom.MoveXY(0, -0.33))...,
	).Transform(geom.RotateD(45)),
}
