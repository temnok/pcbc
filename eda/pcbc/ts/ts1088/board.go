package ts1088

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/ts/xunpu"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var Board = &lib.Component{
	Cuts: path.Paths{
		path.RoundRect(9.5, 6.5, 1),
	},
	Components: lib.Components{
		{
			Transform: geom.MoveXY(3, 0).RotateD(-90),
			Components: lib.Components{
				xunpu.SwitchTS1088,
			},
		},
		{
			Transform: geom.MoveXY(-3.25, 0).RotateD(-90),
			Components: lib.Components{
				mph100imp40f.G_V_SP_x2,
			},
		},
		{
			Transform: geom.MoveXY(-0.25, 0),
			Components: lib.Components{
				pcbc.MountHole,
			},
		},
	},
	Marks: path.Strokes{}.Append(
		font.CenterBold("SW").Transform(geom.MoveXY(-0.25, 2.4).ScaleXY(2, 1.5)),
		pcbc.Logo.Transform(geom.MoveXY(-1, -2.1).ScaleK(1)),
		pcbc.TmnkTech.Transform(geom.MoveXY(0.65, -2.2).ScaleK(0.8)),
	),
}

func init() {
	pad := Board.Squash().Pads.Centers()

	Board.Tracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[0]}.XY(pad[2]),
			eda.Track{pad[1]}.XY(pad[3]),
		),
	}

	Board.GroundTracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[4]}.DX(0.7),
			eda.Track{pad[7]}.DX(-0.7),
		),
	}
}
