package e73

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/ebyte"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var (
	labelShift = geom.XY{2.54 / 0.8, 0}
	labelScale = geom.XY{0.8, 1.8}

	chip = &lib.Component{
		Transform: geom.MoveXY(0, 3.9),
		Components: lib.Components{
			ebyte.E73,
		},
	}

	pin = chip.Squash().Pads.Centers()

	headers = &lib.Component{
		Transform: geom.MoveXY(0, 3.05),

		Components: lib.Components{
			{
				Transform: geom.MoveXY(-10.2, -1).RotateD(-90),
				Components: lib.Components{
					mph100imp40f.G_V_SP_x9,
				},
			},
			{
				Transform: geom.MoveXY(0, -15.3),
				Components: lib.Components{
					mph100imp40f.G_V_SP_x9,
				},
			},
			{
				Transform: geom.MoveXY(10.2, -1).RotateD(90),
				Components: lib.Components{
					mph100imp40f.G_V_SP_x9,
				},
			},
		},
		//Marks: path.Strokes{
		//	font.Bold: path.Paths{}.Append(
		//		font.StringsPaths([]string{
		//			"P113", "P115", "P002", "P029", "P031", "P109", "P012", "GND",
		//		}, font.AlignCenter, labelShift).Transform(geom.MoveXY(-10.6, -1).RotateD(-90).Scale(labelScale)),
		//		font.StringsPaths([]string{
		//			"VDD", "P008", "P006", "P004", "P026", "P024", "P022", "P020", "P018", "P015", "VDDH",
		//		}, font.AlignCenter, labelShift).Transform(geom.MoveXY(0, -11.95).Scale(labelScale)),
		//		font.StringsPaths([]string{
		//			"D-", "D+", "P013", "P100", "SWD", "SWC", "P009", "P010",
		//		}, font.AlignCenter, labelShift).Transform(geom.MoveXY(10.6, -1).RotateD(90).Scale(labelScale)),
		//	),
		//},
	}

	pad = headers.Squash().Pads.Centers()

	mountHoles = &lib.Component{
		Transform: geom.MoveXY(0, 3),
		Components: lib.Components{
			{
				Transform: geom.MoveXY(-5, -11.3),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
			{
				Transform: geom.MoveXY(5, -11.3),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
		},
	}

	mountPad = mountHoles.Squash().Pads.Centers()

	Board_nRF52840 = &lib.Component{
		Clears: path.Paths{
			path.Rect(24, 4.6).Transform(geom.MoveXY(0, 11.6)),
		},

		Cuts: path.Paths{
			path.RoundRect(23.5, 27.5, 1),
		},

		Components: lib.Components{
			chip,
			headers,
			mountHoles,
		},

		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pin[0]}.DX(-2).DY(0.5).YX(pad[0]),
				eda.Track{pin[1]}.XY(pad[1]),
				eda.Track{pin[2]}.XY(pad[2]),
				eda.Track{pin[3]}.XY(pad[3]),

				eda.Track{pin[5]}.XY(pad[4]),
				eda.Track{pin[6]}.XY(pad[5]),
				eda.Track{pin[7]}.DX(1).DX(0.5).YX(pad[8]),
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
				eda.Track{pin[20]}.DX(-1).DX(-0.5).YX(pad[18]),
				eda.Track{pin[21]}.XY(pad[21]),
				eda.Track{pin[22]}.XY(pad[22]),
				eda.Track{pin[23]}.DX(-1).DX(-1.2).DY(-10.9).DX(4).YX(pad[17]),
				eda.Track{pin[24]}.XY(pad[23]),
				eda.Track{pin[25]}.XY(pad[24]),
				eda.Track{pin[26]}.XY(pad[25]),
				eda.Track{pin[27]}.DX(2).DY(0.5).YX(pad[26]),
			),
		},

		GroundTracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pin[4]}.DX(-0.8),
				eda.Track{pin[4]}.DX(1).DX(1.2).DY(-10.9).DX(-4).YX(pad[9]),
				eda.Track{pin[15]}.DY(0.8),
				eda.Track{pin[15]}.DY(-0.8),
				eda.Track{pad[9]}.DX(-0.8).DY(-0.8),
				eda.Track{mountPad[0]}.DX(0.5),
				eda.Track{mountPad[3]}.DX(-0.5),
				eda.Track{mountPad[6]}.DX(0.5),
				eda.Track{mountPad[9]}.DX(-0.5),
			),
		},
		//
		//	Marks: path.Strokes{}.Append(
		//		pcbc.Logo.Transform(geom.MoveXY(-4.8, -6.5).ScaleK(1)),
		//		//pcbc.TmnkTech.Transform(geom.MoveXY(4.8, -6.7).ScaleK(1)),
		//		font.CenterBold("MS88SF21").Transform(geom.MoveXY(1.3, -5.9).ScaleXY(1.8, 1.8)),
		//		font.CenterBold("nRF52840").Transform(geom.MoveXY(1.3, -7.4).ScaleXY(1.8, 1.8)),
		//	),
	}
)
