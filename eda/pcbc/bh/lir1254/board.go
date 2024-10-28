package lir1254

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/battery/holder"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Board = &eda.Component{
	Cuts: path.Paths{
		path.RoundRect(21, 14, 1),
	},

	Components: eda.Components{
		{
			Transform: transform.Move(0, -5.5),
			Components: eda.Components{
				mph100imp40f.G_V_SP_x8,
			},
			Marks: font.CenterBolds(
				[]string{"3V7", "3V7", "3V7", "3V7", "3V7", "3V7", "3V7", "3V7"},
				path.Point{X: 2.54},
			).Apply(transform.Move(0, 1.8)),
		},
		holder.LIR1254.Arrange(transform.Move(0, 2.75)),
		pcbc.MountHole.Arrange(transform.Rotate(-45).Move(-7.5, -1.75)),
		pcbc.MountHole.Arrange(transform.Rotate(45).Move(7.5, -1.75)),
	},

	Marks: path.Strokes{}.Append(
		font.CenterBold("LIR1254").Apply(transform.Scale(1, 2).Move(-7.8, 6)),
		font.CenterBold("COIN BAT").Apply(transform.Scale(0.9, 2).Move(7.8, 6)),

		pcbc.Logo.Apply(transform.Scale(1.2, 1.2).Move(-5, -2)),
		pcbc.TmnkTech.Apply(transform.Move(5, -2.3)),
	),
}

func init() {
	pad := Board.PadCenters()

	Board.Tracks = eda.TrackPaths(
		eda.Track{pad[0]}.XY(pad[7]),
		eda.Track{pad[0]}.DX(-0.8).YX(pad[8]),
		eda.Track{pad[7]}.DX(0.8).YX(pad[9]),
	)
}
