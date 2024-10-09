package pcbc

import (
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var Logo = path.Strokes{
	font.Bold: append(
		font.StringPaths("pc", font.AlignCenter).Transform(geom.MoveXY(0, 0.33)),
		font.StringPaths("bc", font.AlignCenter).Transform(geom.MoveXY(0, -0.33))...,
	).Transform(geom.RotateD(45)),
}

var TmnkTech = path.Strokes{}.Append(
	font.CenterBold("TMNK").Transform(geom.MoveXY(0, 0.3).ScaleXY(0.75, 0.5)),
	font.CenterBold("TECH").Transform(geom.MoveXY(0, -0.3).ScaleXY(0.75, 0.5)),
)
