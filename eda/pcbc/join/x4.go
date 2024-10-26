package join

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var X4 = &lib.Component{
	Cuts: path.Paths{
		path.RoundRect(14.5, 4, 1),
	},
	Components: lib.Components{
		mph100imp40f.G_V_SP_x4.Arrange(transform.Move(-1.9, -0.5)),
		pcbc.MountHole.Arrange(transform.Move(5.2, 0)),
	},
	Marks: path.Strokes{}.Append(
		path.Strokes{
			0.3: path.Paths{
				path.Lines(path.Points{{-1.5 * 2.54, 0}, {1.5 * 2.54, 0}}),
			}.Append(path.Circle(0.3).Clone(4, 2.54, 0)),
		}.Apply(transform.Move(-1.9, 1.35)),
		pcbc.Logo.Apply(transform.ScaleK(0.6).Move(6.5, 1.4)),
		pcbc.TmnkTech.Apply(transform.ScaleK(0.4).Move(6.6, -1.2)),
	),
}

func init() {
	pad := X4.Flatten().Pads.Centers()

	X4.Tracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[0]}.XY(pad[3]),
		),
	}

	X4.GroundTracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[6]}.DX(-0.4).DY(0.4),
			eda.Track{pad[8]}.DX(-0.4).DY(-0.4),
		),
	}
}
