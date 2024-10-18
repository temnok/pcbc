package ws2812b_2020

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/worldsemi"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Board = &lib.Component{
	Cuts: path.Paths{
		path.RoundRect(10, 5.5, 1),
	},
	Components: lib.Components{
		worldsemi.WS2812B_2020.Arrange(transform.Move(-3, 0)),
		mph100imp40f.G_V_SP_x2.Arrange(transform.Rotate(-90).Move(3.5, 0)),
		pcbc.MountHole.Arrange(transform.Rotate(30).Move(0.55, 0)),
	},
	Marks: path.Strokes{}.Append(
		font.CenterBold("WS2812B").Apply(transform.Scale(1, 1.8).Move(-2.1, 1.8)),
		font.CenterBold("LED").Apply(transform.Scale(1.5, 1.8).Move(-3, -1.8)),
		font.CenterBold("DI").Apply(transform.Scale(0.8, 1.4).Move(1.6, 2)),
		font.CenterBold("VDD").Apply(transform.Scale(0.8, 1.4).Move(1.3, -2)),
		pcbc.Logo.Apply(transform.ScaleK(0.7).Move(-0.6, -2)),
	),
}

func init() {
	pad := Board.Squash().Pads.Centers()

	Board.Tracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[3]}.DY(1.5).XY(pad[5]),
			eda.Track{pad[4]}.DY(-1.5).XY(pad[6]),
		),
	}

	Board.GroundTracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[2]}.DY(0.5).XY(pad[9]),
			eda.Track{pad[7]}.DX(0.5).DY(0.5),
			eda.Track{pad[12]}.DX(0.5).DY(-0.5),
			eda.Track{pad[10]}.DX(-0.5).DY(-0.5),
		),
	}
}
