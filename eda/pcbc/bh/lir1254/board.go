package lir1254

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/battery/holder"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var Board = &lib.Component{
	Cuts: path.Paths{
		path.RoundRect(21, 14, 1),
	},
	Components: lib.Components{
		{
			Transform: geom.MoveXY(0, -5.5),
			Components: lib.Components{
				mph100imp40f.G_V_SP_x8,
			},
			Marks: font.CenterBolds(
				[]string{"3V7", "3V7", "3V7", "3V7", "3V7", "3V7", "3V7", "3V7"},
				geom.XY{2.54, 0},
			).Transform(geom.MoveXY(0, 1.8)),
		},
		{
			Transform: geom.MoveXY(0, 2.75),
			Components: lib.Components{
				holder.LIR1254,
			},
		},
		{
			Transform: geom.MoveXY(-7.5, -1.75),
			Components: lib.Components{
				pcbc.MountHole,
			},
		},
		{
			Transform: geom.MoveXY(7.5, -1.75),
			Components: lib.Components{
				pcbc.MountHole,
			},
		},
	},

	Marks: path.Strokes{}.Append(
		font.CenterBold("LIR1254").Transform(geom.MoveXY(-7.8, 6).ScaleXY(1, 2)),
		font.CenterBold("COIN BAT").Transform(geom.MoveXY(7.8, 6).ScaleXY(0.9, 2)),

		pcbc.Logo.Transform(geom.MoveXY(-5, -2).ScaleK(1.2)),
		pcbc.TmnkTech.Transform(geom.MoveXY(5, -2.3).ScaleK(1)),
	),
}

func init() {
	pad := Board.Squash().Pads.Centers()

	Board.Tracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[0]}.XY(pad[7]),
			eda.Track{pad[0]}.DX(-0.8).YX(pad[8]),
			eda.Track{pad[7]}.DX(0.8).YX(pad[9]),
		),
	}

	Board.GroundTracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[10]}.DX(-2).DY(2),
			eda.Track{pad[10]}.DX(2).DY(2),

			eda.Track{pad[10]}.YX(pad[12]),
			eda.Track{pad[10]}.YX(pad[19]),
			eda.Track{pad[16]}.YX(pad[21]),
		),
	}
}
