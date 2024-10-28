package ts026a

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/ts/hyp"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Board = &eda.Component{
	Cuts: path.Paths{
		path.RoundRect(9, 5.5, 1),
	},

	Components: eda.Components{
		hyp.Switch1TS026A.Arrange(transform.Rotate(90).Move(3, 0)),
		mph100imp40f.G_V_SP_x2.Arrange(transform.Rotate(-90).Move(-3, 0)),
		pcbc.MountHole,
	},

	Marks: path.Strokes{}.Append(
		font.CenterBold("SW").Apply(transform.ScaleK(1.5).Move(0, 2)),
		pcbc.Logo.Apply(transform.ScaleK(0.7).Move(0, -1.9)),
	),
}

func init() {
	pad := Board.PadCenters()

	Board.Tracks = eda.TrackPaths(
		eda.Track{pad[0]}.XY(pad[1]),
		eda.Track{pad[1]}.DY(0.7).XY(pad[4]),

		eda.Track{pad[2]}.XY(pad[3]),
		eda.Track{pad[3]}.DY(-0.7).XY(pad[5]),
	)
}
