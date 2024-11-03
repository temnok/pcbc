package usbc

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/connector/yiyuan"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	mountPos     = transform.Move(-5, 1.5)
	connectorPos = transform.Move(2.5, 2)
	headerPos    = transform.Move(0, -2.5)
)

var Board = &eda.Component{
	Cuts: path.Paths{
		path.RoundRect(16, 8, 1),
	},

	MarkStrokes: path.Strokes{}.Append(
		font.CenterBold("YTC-TC8-565").Apply(transform.Scale(0.7, 0.9).Move(-5, 3.4)),
		pcbc.LogoStrokes.Apply(transform.ScaleK(0.8).Move(-7.2, 1.5)),
	),

	Components: eda.Components{
		pcbc.MountHole.Arrange(mountPos),
		yiyuan.YTC_TC8_565.Arrange(connectorPos),
		mph100imp40f.G_V_SP_x6.Arrange(headerPos),
	},
}

func init() {
	pins := yiyuan.YTC_TC8_565.Pads.Apply(connectorPos).Centers(transform.Identity)
	pads := mph100imp40f.G_V_SP_x6.Pads.Apply(headerPos).Centers(transform.Identity)

	Board.Tracks = eda.TrackPaths(
		eda.Track{pins[1]}.Y(1.75).Y(2).X(pins[6].X).Y(1.75).Y(pins[6].Y),
		eda.Track{pins[2]}.Y(-0.3).Y(-0.8).X(pads[1].X).Y(pads[1].Y),
		eda.Track{pins[3]}.Y(-0.8).Y(-1.3).X(pads[2].X).Y(pads[2].Y),
		eda.Track{pads[3]}.X(pins[4].X).Y(pins[4].Y),
		eda.Track{pads[4]}.X(pins[5].X).Y(pins[5].Y),
		eda.Track{pads[5]}.X(pins[6].X).Y(pins[6].Y),
	)

	Board.GroundTracks = eda.TrackPaths(
		eda.Track{pins[0]}.Y(1.75).Y(2.5).X(pins[7].X).Y(1.75).Y(pins[7].Y),
		eda.Track{pins[0]}.Y(0).Y(-0.3).X(pads[0].X).Y(pads[0].Y),
	)

	padNames := []string{
		"GND",
		"CC1",
		"DN2",
		"DP2",
		"SBU1",
		"VBUS",
	}

	const tenth = 2.54

	for i, padName := range padNames {
		Board.MarkStrokes = Board.MarkStrokes.Append(
			font.CenterBold(padName).Apply(transform.Scale(0.9, 1.2).Move(tenth*(float64(i)-2.5), -0.6)),
		)
	}
}
