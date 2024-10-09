package usbc

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/connector/yiyuan"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var (
	mountPos     = geom.MoveXY(-5, 1.5)
	connectorPos = geom.MoveXY(2.5, 2)
	headerPos    = geom.MoveXY(0, -2.5)
)

var Board = &lib.Component{
	Cuts: path.Paths{
		path.RoundRect(16, 8, 1),
	},

	Marks: path.Strokes{}.Append(
		font.CenterBold("YTC-TC8-565").Transform(geom.MoveXY(-5, 3.4).ScaleXY(0.7, 0.9)),
		pcbc.Logo.Transform(geom.MoveXY(-7.2, 1.5).ScaleK(0.8)),
		pcbc.TmnkTech.Transform(geom.MoveXY(7.25, 3.5).ScaleK(0.4)),
	),

	Components: lib.Components{
		{
			Transform: mountPos,
			Components: lib.Components{
				pcbc.MountHole,
			},
		},
		{
			Transform: connectorPos,
			Components: lib.Components{
				yiyuan.YTC_TC8_565,
			},
		},
		{
			Transform: headerPos,
			Components: lib.Components{
				mph100imp40f.G_V_SP_x6,
			},
		},
	},
}

func init() {
	mountPads := pcbc.MountHole.Pads.Transform(mountPos).Centers()
	pins := yiyuan.YTC_TC8_565.Pads.Transform(connectorPos).Centers()
	pads := mph100imp40f.G_V_SP_x6.Pads.Transform(headerPos).Centers()

	Board.Tracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{pins[1]}.Y(1.75).Y(2).X(pins[6].X).Y(1.75).Y(pins[6].Y),
			eda.Track{pins[2]}.Y(-0.3).Y(-0.8).X(pads[1].X).Y(pads[1].Y),
			eda.Track{pins[3]}.Y(-0.8).Y(-1.3).X(pads[2].X).Y(pads[2].Y),
			eda.Track{pads[3]}.X(pins[4].X).Y(pins[4].Y),
			eda.Track{pads[4]}.X(pins[5].X).Y(pins[5].Y),
			eda.Track{pads[5]}.X(pins[6].X).Y(pins[6].Y),
		),
	}

	Board.GroundTracks = path.Strokes{
		0: eda.TrackPaths(
			eda.Track{mountPads[5], {-4.1, -0.3}},
			eda.Track{pins[0]}.Y(1.75).Y(2.5).X(pins[7].X).Y(1.75).Y(pins[7].Y),
			eda.Track{pins[0]}.Y(0).Y(-0.3).X(pads[0].X).Y(pads[0].Y),
		),
	}

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
		Board.Marks = Board.Marks.Append(
			font.CenterBold(padName).Transform(geom.MoveXY(tenth*(float64(i)-2.5), -0.6).ScaleXY(0.9, 1.2)),
		)
	}
}
