package sw

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/switch/xunpu"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var XunpuTS1088 = &lib.Component{
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
		font.CenterBold("SW").Transform(geom.MoveXY(-0.25, 2.4).ScaleK(1.5)),
		pcbc.Logo.Transform(geom.MoveXY(-0.25, -2.2).ScaleK(0.8)),
	),
}

func init() {
	pad := XunpuTS1088.Squash().Pads.Centers()

	XunpuTS1088.Tracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[0]}.XY(pad[2]),
			eda.Track{pad[1]}.XY(pad[3]),
		),
	}

	XunpuTS1088.GroundTracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[4]}.DX(0.7),
			eda.Track{pad[7]}.DX(-0.7),
		),
	}
}
