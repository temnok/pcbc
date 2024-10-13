package ws2812b_2020

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/worldsemi"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var Board = &lib.Component{
	Cuts: path.Paths{
		path.RoundRect(10, 5.5, 1),
	},
	Components: lib.Components{
		{
			Transform: geom.MoveXY(-3, 0),
			Components: lib.Components{
				worldsemi.WS2812B_2020,
			},
		},
		{
			Transform: geom.MoveXY(3.5, 0).RotateD(-90),
			Components: lib.Components{
				mph100imp40f.G_V_SP_x2,
			},
		},
		{
			Transform: geom.MoveXY(0.55, 0).RotateD(30),
			Components: lib.Components{
				pcbc.MountHole,
			},
		},
	},
	Marks: path.Strokes{}.Append(
		font.CenterBold("WS2812B").Transform(geom.MoveXY(-2.1, 1.8).ScaleXY(1, 1.8)),
		font.CenterBold("LED").Transform(geom.MoveXY(-3, -1.8).ScaleXY(1.5, 1.8)),
		font.CenterBold("DI").Transform(geom.MoveXY(1.6, 2).ScaleXY(0.8, 1.4)),
		font.CenterBold("VDD").Transform(geom.MoveXY(1.3, -2).ScaleXY(0.8, 1.4)),
		pcbc.Logo.Transform(geom.MoveXY(-0.6, -2).ScaleK(0.7)),
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
