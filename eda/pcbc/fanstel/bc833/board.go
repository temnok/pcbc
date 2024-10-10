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
	Board      = board(false)
	ShortBoard = board(true)
)

func board(short bool) *lib.Component {
	header := &lib.Component{
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

	pad := header.Squash().Pads.Centers()

	if short {
		header = &lib.Component{
			Components: lib.Components{
				{
					Transform: geom.MoveXY(-8.9, -2.9-2.54).RotateD(-90),
					Components: lib.Components{
						mph100imp40f.G_V_SP_x6,
					},
				},
				{
					Transform: geom.MoveXY(0, -15.5),
					Components: lib.Components{
						mph100imp40f.G_V_SP_x8,
					},
				},
				{
					Transform: geom.MoveXY(8.9, -2.9-2.54).RotateD(90),
					Components: lib.Components{
						mph100imp40f.G_V_SP_x6,
					},
				},
			},
		}
	}

	pin := fanstel.BC833.Pads.Centers()

	labelShift := geom.XY{2.54 / 0.8, 0}
	labelScale := geom.XY{0.8, 1.3}

	tracks := eda.TrackPaths(
		eda.Track{pin[8]}.DX(1.3).DY(3.5).DY(1.4).DX(-5.4).DY(1.5).YX(pad[0]),
		eda.Track{pin[7]}.DX(0.7).DY(3.3).DX(-1).XY(pad[1]),
		eda.Track{pin[6]}.DY(1.7).DX(-0.5).XY(pad[2]),
		eda.Track{pin[1]}.XY(pad[3]),
		eda.Track{pin[2]}.YX(pad[4]),
		eda.Track{pin[3]}.YX(pad[5]),
		eda.Track{pin[4]}.DX(-3.1).YX(pad[6]),

		eda.Track{pin[21]}.DY(-2.7).DX(-4.6).DX(-1).DY(-7.6).DY(-1.9).XY(pad[8]),
		eda.Track{pin[20]}.DY(-2.3).DX(-4.3).DX(-0.7).DY(-7.4).DY(-2.1).XY(pad[9]),
		eda.Track{pin[9]}.DX(-0.4).YX(pad[10]),
		eda.Track{pin[10]}.DX(-0.8).YX(pad[11]),
		eda.Track{pin[11]}.DX(-1.2).YX(pad[12]),
		eda.Track{pin[12]}.DX(-1.6).YX(pad[13]),
		eda.Track{pin[13]}.DX(-2.0).DY(-8.8).DX(2.1).XY(pad[14]),
		eda.Track{pin[14]}.DX(-2.4).DY(-8.2).DX(1.9).XY(pad[15]),

		eda.Track{pin[19]}.DX(3.5).YX(pad[16]),
		eda.Track{pin[15]}.DX(3.1).YX(pad[17]),
		eda.Track{pin[16]}.YX(pad[18]),
		eda.Track{pin[17]}.YX(pad[19]),
		eda.Track{pin[18]}.XY(pad[20]),
		eda.Track{pin[22]}.DY(0.7).XY(pad[21]),
		eda.Track{pin[23]}.DY(1.3).XY(pad[22]),
		eda.Track{pin[24]}.DY(1.9).DX(6.1).DY(1.5).YX(pad[23]),
	)

	leftLabels := []string{"P031", "P030", "P029", "VDD", "P003", "P002", "P028", "GND"}
	centerLabels := []string{"P020", "P017", "P004", "P005", "P109", "P011", "VDDH", "VBUS"}
	rightLabels := []string{"D+", "D-", "P015", "P018", "SWD", "SWC", "P009", "P010"}

	if short {
		tracks[0] = nil
		tracks[1] = nil
		tracks[21] = nil
		tracks[22] = nil

		leftLabels[0] = ""
		leftLabels[1] = ""
		rightLabels[6] = ""
		rightLabels[7] = ""
	}

	leftMount := geom.XY{-5, -10.5}
	rightMount := geom.XY{5, -10.5}

	shiftedBoard := &lib.Component{
		Components: lib.Components{
			fanstel.BC833,

			header,

			{
				Transform: geom.Move(leftMount),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
			{
				Transform: geom.Move(rightMount),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
		},

		Tracks: path.Strokes{
			0: tracks,
		},

		GroundTracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pin[5]}.DX(-2.5).YX(pad[7]),
				eda.Track{pin[5]}.DX(2).YX(pin[28]).XY(pin[26]).XY(pin[25]).YX(pin[27]).XY(pin[28]),

				eda.Track{leftMount}.YX(pad[7]),
				eda.Track{leftMount}.DX(1.5).DY(6.1).YX(pin[27]),
				eda.Track{rightMount}.DX(-1.3).DY(1.3),
				eda.Track{rightMount}.DX(1.3).DY(-1.3),
			),
		},

		Marks: path.Strokes{}.Append(
			pcbc.Logo.Transform(geom.MoveXY(0, -8.3).ScaleK(1.2)),
			font.CenterBold("BC833").Transform(geom.MoveXY(0, -10.4).ScaleK(2)),
			font.CenterBold("nRF52833").Transform(geom.MoveXY(0, -12.2).ScaleK(1.5)),

			font.CenterBolds(leftLabels, labelShift).
				Transform(geom.MoveXY(-7.1, -2.9).RotateD(-90).Scale(labelScale)),
			font.CenterBolds(centerLabels, labelShift).
				Transform(geom.MoveXY(0, -13.65).Scale(labelScale)),
			font.CenterBolds(rightLabels, labelShift).
				Transform(geom.MoveXY(7.1, -2.9).RotateD(90).Scale(labelScale)),
		),
	}

	boardShift := 4.75
	boardCut := path.RoundRect(21, 24.6, 1)
	boardClears := path.Paths{
		path.Rect(21.5, 5.5).Transform(geom.MoveXY(0, 9.7)),
	}

	if short {
		boardShift += 2.54
		boardCut = path.RoundRect(21, 19.5, 1)
		boardClears = nil
	}

	return &lib.Component{
		Clears: boardClears,

		Cuts: path.Paths{
			boardCut,
		},

		Components: lib.Components{
			{
				Transform: geom.MoveXY(0, boardShift),
				Components: lib.Components{
					shiftedBoard,
				},
			},
		},
	}
}
