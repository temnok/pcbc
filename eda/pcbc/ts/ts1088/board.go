package ts1088

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/ts/xunpu"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Board = &lib.Component{
	Cuts: path.Paths{
		path.RoundRect(9.5, 6.5, 1),
	},

	Components: lib.Components{
		xunpu.SwitchTS1088.Arrange(transform.Rotate(-90).Move(3, 0)),
		mph100imp40f.G_V_SP_x2.Arrange(transform.Rotate(-90).Move(-3.25, 0)),
		pcbc.MountHole.Arrange(transform.Move(-0.25, 0)),
	},

	Marks: path.Strokes{}.Append(
		font.CenterBold("SW").Apply(transform.Scale(2, 1.5).Move(-0.25, 2.4)),
		pcbc.Logo.Apply(transform.Move(-1, -2.1)),
		pcbc.TmnkTech.Apply(transform.ScaleK(0.8).Move(0.65, -2.2)),
	),
}

func init() {
	pad := Board.Flatten().Pads.Centers()

	Board.Tracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pad[0]}.XY(pad[2]),
			eda.Track{pad[1]}.XY(pad[3]),
		),
	}
}
