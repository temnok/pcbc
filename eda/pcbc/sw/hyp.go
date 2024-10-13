package sw

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/switch/hyp"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var Hyp1TS026A = &lib.Component{
	Cuts: path.Paths{
		path.RoundRect(9, 5.5, 1),
	},
	Components: lib.Components{
		{
			Transform: geom.MoveXY(3, 0).RotateD(90),
			Components: lib.Components{
				hyp.Switch1TS026A,
			},
		},
		{
			Transform: geom.MoveXY(-3, 0).RotateD(-90),
			Components: lib.Components{
				mph100imp40f.G_V_SP_x2,
			},
		},
		{
			Transform: geom.MoveXY(0, 0),
			Components: lib.Components{
				pcbc.MountHole,
			},
		},
	},
	Marks: path.Strokes{}.Append(
		font.CenterBold("SW").Transform(geom.MoveXY(0, 2).ScaleK(1.5)),
		pcbc.Logo.Transform(geom.MoveXY(0, -1.9).ScaleK(0.7)),
	),
}

func init() {
	pad := Hyp1TS026A.Squash().Pads.Centers()

	Hyp1TS026A.Tracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[0]}.XY(pad[1]),
			eda.Track{pad[1]}.DY(0.7).XY(pad[4]),

			eda.Track{pad[2]}.XY(pad[3]),
			eda.Track{pad[3]}.DY(-0.7).XY(pad[5]),
		),
	}

	Hyp1TS026A.GroundTracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[6]}.DX(0.7),
			eda.Track{pad[9]}.DX(-0.7),
		),
	}
}
