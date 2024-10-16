package ms88sf2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/minewsemi"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var (
	labelShift = geom.XY{2.54 / 0.8, 0}
	labelScale = geom.XY{0.8, 1.8}

	chip = &lib.Component{
		Transform: geom.MoveXY(0, 6.6),
		Components: lib.Components{
			minewsemi.MS88SF2,
		},
	}

	pin = chip.Squash().Pads.Centers()

	headers = &lib.Component{
		Transform: geom.MoveXY(0, 3.05),

		Components: lib.Components{
			{
				Transform: geom.MoveXY(-12.7, -1).RotateD(-90),
				Components: lib.Components{
					mph100imp40f.G_V_SP_x8,
				},
			},
			{
				Transform: geom.MoveXY(0, -14),
				Components: lib.Components{
					mph100imp40f.G_V_SP_x11,
				},
			},
			{
				Transform: geom.MoveXY(12.7, -1).RotateD(90),
				Components: lib.Components{
					mph100imp40f.G_V_SP_x8,
				},
			},
		},
		Marks: path.Strokes{
			font.Bold: path.Paths{}.Append(
				font.StringsPaths([]string{
					"P113", "P115", "P002", "P029", "P031", "P109", "P012", "GND",
				}, font.AlignCenter, labelShift).Transform(geom.MoveXY(-10.6, -1).RotateD(-90).Scale(labelScale)),
				font.StringsPaths([]string{
					"VDD", "P008", "P006", "P004", "P026", "P024", "P022", "P020", "P018", "P015", "VDDH",
				}, font.AlignCenter, labelShift).Transform(geom.MoveXY(0, -11.95).Scale(labelScale)),
				font.StringsPaths([]string{
					"D-", "D+", "P013", "P100", "SWD", "SWC", "P009", "P010",
				}, font.AlignCenter, labelShift).Transform(geom.MoveXY(10.6, -1).RotateD(90).Scale(labelScale)),
			),
		},
	}

	pad = append(path.Paths{nil}, headers.Squash().Pads...).Centers()

	mountHoles = &lib.Component{
		Transform: geom.MoveXY(0, 3),
		Components: lib.Components{
			{
				Transform: geom.MoveXY(-7.5, -9.7),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
			{
				Transform: geom.MoveXY(7.5, -9.7),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
		},
	}

	mountPad = mountHoles.Squash().Pads.Centers()

	Board_nRF52840 = &lib.Component{
		Cuts: path.Paths{
			path.RoundRect(28.4, 24.9, 1),
		},

		Components: lib.Components{
			chip,
			headers,
			mountHoles,
		},

		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pad[1]}.YX(pin[2]),
				eda.Track{pad[2]}.YX(pin[3]),
				eda.Track{pad[3]}.YX(pin[4]),
				eda.Track{pad[4]}.DX(1.7).Y(pin[5].Y-1.2).YX(pin[5]),
				eda.Track{pad[5]}.DX(2.4).Y(pin[6].Y-0.4).YX(pin[6]),
				eda.Track{pad[6]}.YX(pin[11]),
				eda.Track{pad[7]}.YX(pin[12]),
				eda.Track{pad[9]}.DX(2.4).Y(pin[14].Y-0.6).YX(pin[14]),

				eda.Track{pad[10]}.DY(1.7).X(-4.8).DY(2.1).Y(pin[19].Y-1).YX(pin[10]),
				eda.Track{pad[11]}.X(-3.6).Y(pin[20].Y-1.5).YX(pin[9]),
				eda.Track{pad[12]}.X(-2.4).Y(pin[21].Y-2).YX(pin[8]),
				eda.Track{pad[13]}.X(-1.2).Y(pin[22].Y-2.5).YX(pin[7]),

				eda.Track{pad[14]}.Y(pin[23].Y-3).YX(pin[23]),

				eda.Track{pad[15]}.X(1.2).Y(pin[22].Y-2.5).YX(pin[22]),
				eda.Track{pad[16]}.X(2.4).Y(pin[21].Y-2).YX(pin[21]),
				eda.Track{pad[17]}.X(3.6).Y(pin[20].Y-1.5).YX(pin[20]),
				eda.Track{pad[18]}.DY(1.7).X(4.8).DY(2.1).Y(pin[19].Y-1).YX(pin[19]),

				eda.Track{pad[19]}.DX(-2.3).Y(pin[15].Y-0.6).YX(pin[15]),
				eda.Track{pad[20]}.DX(-1.6).Y(pin[16].Y-1.4).YX(pin[16]),
				eda.Track{pad[21]}.YX(pin[17]),
				eda.Track{pad[22]}.YX(pin[18]),
				eda.Track{pad[23]}.DX(-2.2).Y(pin[24].Y-0.4).YX(pin[24]),
				eda.Track{pad[24]}.DX(-1.5).Y(pin[25].Y-1.2).YX(pin[25]),
				eda.Track{pad[25]}.YX(pin[26]),
				eda.Track{pad[26]}.YX(pin[27]),
				eda.Track{pad[27]}.YX(pin[28]),
			),
		},

		GroundTracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pin[1]}.DX(1.5),
				eda.Track{pin[1]}.DX(-1.5),
				eda.Track{pad[8]}.DX(1.7).Y(pin[13].Y-1.4).YX(pin[13]),
				eda.Track{mountPad[1]}.DX(0.5).Y(pin[13].Y-0.5).YX(pin[13]),
				eda.Track{mountPad[4]}.DX(-0.5).DY(-0.5),
				eda.Track{mountPad[8]}.DX(-0.5).DY(0.5),
				eda.Track{mountPad[11]}.DX(0.5).DY(-0.5),
			),
		},

		Marks: path.Strokes{}.Append(
			pcbc.Logo.Transform(geom.MoveXY(-4.8, -6.5).ScaleK(1)),
			//pcbc.TmnkTech.Transform(geom.MoveXY(4.8, -6.7).ScaleK(1)),
			font.CenterBold("MS88SF21").Transform(geom.MoveXY(1.3, -5.9).ScaleXY(1.8, 1.8)),
			font.CenterBold("nRF52840").Transform(geom.MoveXY(1.3, -7.4).ScaleXY(1.8, 1.8)),
		),
	}
)
