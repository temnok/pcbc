package e73

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/ebyte"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	labelShift = path.Point{X: 2.54 / 0.8}
	labelScale = transform.Scale(0.8, 1.8)

	chip = ebyte.E73.Arrange(transform.Move(0, 3.9))

	pin = chip.PadCenters()

	headers = &eda.Component{
		Transform: transform.Move(0, 3.05),

		Components: eda.Components{
			mph100imp40f.G_V_SP_x9.Arrange(transform.Rotate(-90).Move(-10.2, -1)),
			mph100imp40f.G_V_SP_x9.Arrange(transform.Move(0, -15.3)),
			mph100imp40f.G_V_SP_x9.Arrange(transform.Rotate(90).Move(10.2, -1)),
		},

		MarkStrokes: path.Strokes{
			font.Bold: path.Paths{}.Append(
				font.ShiftedAlignedPaths(labelShift, font.AlignCenter,
					"P111", "P110", "P003", "P028", "P002", "P029", "P031", "P030", "P113",
				).Apply(labelScale.Rotate(-90).Move(-8.1, -1)),
				font.ShiftedAlignedPaths(labelShift, font.AlignCenter,
					"GND", "P000", "P001", "P005", "P109", "VDD", "VDDH", "DCCH", "P024",
				).Apply(labelScale.Move(0, -13.2)),
				font.ShiftedAlignedPaths(labelShift, font.AlignCenter,
					"P013", "P018", "VBUS", "D-", "D+", "SWD", "SWC", "P009", "P010",
				).Apply(labelScale.Rotate(90).Move(8.1, -1)),
			),
		},
	}

	pad = headers.PadCenters()

	mountHoles = &eda.Component{
		Components: eda.Components{
			pcbc.MountHole.Arrange(transform.Move(-5, -8)),
			pcbc.MountHole.Arrange(transform.Move(5, -8)),
		},
	}

	Board_nRF52840 = &eda.Component{
		Clears: path.Paths{
			path.Rect(24, 4.6).Apply(transform.Move(0, 11.6)),
		},

		Cuts: path.Paths{
			path.RoundRect(23.5, 27.5, 1),
		},

		Components: eda.Components{
			chip,
			headers,
			mountHoles,
		},

		Tracks: eda.TrackPaths(
			eda.Track{pin[0]}.DX(-2).DY(0.5).YX(pad[0]),
			eda.Track{pin[1]}.XY(pad[1]),
			eda.Track{pin[2]}.XY(pad[2]),
			eda.Track{pin[3]}.XY(pad[3]),

			eda.Track{pin[5]}.DX(1).DX(0.5).YX(pad[8]),
			eda.Track{pin[6]}.XY(pad[4]),
			eda.Track{pin[7]}.XY(pad[5]),
			eda.Track{pin[8]}.XY(pad[6]),
			eda.Track{pin[9]}.YX(pad[7]),

			eda.Track{pin[10]}.DX(-1).XY(pad[10]),
			eda.Track{pin[11]}.DY(-2.1).X(-2.4).YX(pad[11]),
			eda.Track{pin[12]}.DY(-1.8).X(-1.2).YX(pad[12]),
			eda.Track{pin[13]}.DY(-1.5).X(0).YX(pad[13]),
			eda.Track{pin[14]}.DY(-1.2).X(1.2).YX(pad[14]),
			eda.Track{pin[16]}.DY(-2.1).X(2.4).YX(pad[15]),
			eda.Track{pin[17]}.DX(1).XY(pad[16]),

			eda.Track{pin[18]}.YX(pad[19]),
			eda.Track{pin[19]}.XY(pad[20]),
			eda.Track{pin[20]}.XY(pad[21]),
			eda.Track{pin[21]}.XY(pad[22]),
			eda.Track{pin[22]}.DX(-1).DX(-0.5).YX(pad[18]),
			eda.Track{pin[23]}.DX(-1).DX(-1.2).DY(-10.9).DX(4).YX(pad[17]),
			eda.Track{pin[24]}.XY(pad[23]),
			eda.Track{pin[25]}.XY(pad[24]),
			eda.Track{pin[26]}.XY(pad[25]),
			eda.Track{pin[27]}.DX(2).DY(0.5).YX(pad[26]),
		),

		GroundTracks: eda.TrackPaths(
			eda.Track{pin[4]}.DX(-1),
			eda.Track{pin[4]}.DX(1).DX(1.2).DY(-10.9).DX(-4).YX(pad[9]),
			eda.Track{pin[15]}.DY(1),
			eda.Track{pin[15]}.DY(-1),
			eda.Track{pad[9]}.DX(-0.8).DY(-0.8),
		),

		MarkStrokes: path.Strokes{}.Append(
			pcbc.LogoStrokes.Apply(transform.Move(-6.3, -6)),
			font.CenterBold("E73").Apply(transform.Scale(1.2, 2).Move(-2.5, -6.8)),
			font.CenterBold("-2G4M08S1C").Apply(transform.Scale(0.8, 2).Move(1.2, -6.8)),
			font.CenterBold("nRF52840").Apply(transform.Scale(1.4, 2).Move(0, -8.4)),
		),
	}
)
