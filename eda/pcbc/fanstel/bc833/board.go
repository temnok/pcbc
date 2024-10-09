package bc833

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/fanstel"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var (
	header = &lib.Component{
		Components: lib.Components{
			{
				Transform: geom.MoveXY(-8.9, -2.9).RotateD(-90),
				Components: lib.Components{
					mph100imp40f.G_V_SP_x8,
				},
			},
			{
				Transform: geom.MoveXY(0, -15.5),
				Components: lib.Components{
					mph100imp40f.G_V_SP_x8,
				},
			},
			{
				Transform: geom.MoveXY(8.9, -2.9).RotateD(90),
				Components: lib.Components{
					mph100imp40f.G_V_SP_x8,
				},
			},
		},
	}

	pads = header.Squash().Pads.Centers()

	pins = fanstel.BC833.Pads.Centers()

	labelShift = geom.XY{2.54 / 0.8, 0}
	labelScale = geom.XY{0.8, 1.3}

	shiftedBoard = &lib.Component{
		Components: lib.Components{
			fanstel.BC833,

			header,

			{
				Transform: geom.MoveXY(-5, -10.5),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
			{
				Transform: geom.MoveXY(5, -10.5),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
		},

		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pins[8]}.DX(1.3).DY(3.5).DY(1.4).DX(-5.4).DY(1.5).YX(pads[0]),
				eda.Track{pins[7]}.DX(0.7).DY(3.3).DX(-1).XY(pads[1]),
				eda.Track{pins[6]}.DY(1.7).DX(-0.5).XY(pads[2]),
				eda.Track{pins[1]}.XY(pads[3]),
				eda.Track{pins[2]}.YX(pads[4]),
				eda.Track{pins[3]}.YX(pads[5]),
				eda.Track{pins[4]}.DX(-3.1).YX(pads[6]),

				eda.Track{pins[21]}.DY(-2.7).DX(-4.6).DX(-1).DY(-7.6).DY(-1.9).XY(pads[8]),
				eda.Track{pins[20]}.DY(-2.3).DX(-4.3).DX(-0.7).DY(-7.4).DY(-2.1).XY(pads[9]),
				eda.Track{pins[9]}.DX(-0.4).YX(pads[10]),
				eda.Track{pins[10]}.DX(-0.8).YX(pads[11]),
				eda.Track{pins[11]}.DX(-1.2).YX(pads[12]),
				eda.Track{pins[12]}.DX(-1.6).YX(pads[13]),
				eda.Track{pins[13]}.DX(-2.0).DY(-8.8).DX(2.1).XY(pads[14]),
				eda.Track{pins[14]}.DX(-2.4).DY(-8.2).DX(1.9).XY(pads[15]),

				eda.Track{pins[19]}.DX(3.5).YX(pads[16]),
				eda.Track{pins[15]}.DX(3.1).YX(pads[17]),
				eda.Track{pins[16]}.YX(pads[18]),
				eda.Track{pins[17]}.YX(pads[19]),
				eda.Track{pins[18]}.XY(pads[20]),
				eda.Track{pins[22]}.DY(0.7).XY(pads[21]),
				eda.Track{pins[23]}.DY(1.3).XY(pads[22]),
				eda.Track{pins[24]}.DY(1.9).DX(6.1).DY(1.5).YX(pads[23]),
			),
		},

		GroundTracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pins[5]}.DX(-2.5).YX(pads[7]),
				eda.Track{pins[5]}.DX(2).YX(pins[28]).XY(pins[26]).XY(pins[25]).YX(pins[27]).XY(pins[28]),

				eda.Track{{-5, -10.5}}.YX(pads[7]),
				eda.Track{{5, -10.5}}.DX(1.3).DY(-1.3),
			),
		},

		Marks: path.Strokes{}.Append(
			pcbc.Logo.Transform(geom.MoveXY(0, -8.3).ScaleK(1.2)),
			font.CenterBold("BC833").Transform(geom.MoveXY(0, -10.4).ScaleK(2)),
			font.CenterBold("nRF52833").Transform(geom.MoveXY(0, -12.2).ScaleK(1.5)),

			font.CenterBolds([]string{
				"P031",
				"P030",
				"P029",
				"VDD",
				"P003",
				"P002",
				"P028",
				"GND",
			}, labelShift).Transform(geom.MoveXY(-7.1, -2.9).RotateD(-90).Scale(labelScale)),
			font.CenterBolds([]string{
				"P020",
				"P017",
				"P004",
				"P005",
				"P109",
				"P011",
				"VDDH",
				"VBUS",
			}, labelShift).Transform(geom.MoveXY(0, -13.65).Scale(labelScale)),
			font.CenterBolds([]string{
				"D+",
				"D-",
				"P015",
				"P018",
				"SWD",
				"SWC",
				"P009",
				"P010",
			}, labelShift).Transform(geom.MoveXY(7.1, -2.9).RotateD(90).Scale(labelScale)),
		),
	}

	Board = &lib.Component{
		Clears: path.Paths{
			path.Rect(21.5, 5.5).Transform(geom.MoveXY(0, 9.7)),
		},

		Cuts: path.Paths{
			path.RoundRect(21, 24.6, 1),
		},

		Components: lib.Components{
			{
				Transform: geom.MoveXY(0, 4.75),
				Components: lib.Components{
					shiftedBoard,
				},
			},
		},
	}
)
