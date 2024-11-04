package bc833

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/fanstel"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	Board      = board(false)
	ShortBoard = board(true)
)

func board(short bool) *eda.Component {
	revision := "v1.0"

	header := &eda.Component{
		Components: eda.Components{
			mph100imp40f.G_V_SP(8).Arrange(transform.Rotate(-90).Move(-8.9, -2.9)),
			mph100imp40f.G_V_SP(8).Arrange(transform.Move(0, -15.5)),
			mph100imp40f.G_V_SP(8).Arrange(transform.Rotate(90).Move(8.9, -2.9)),
		},
	}

	pad := header.PadCenters()

	if short {
		header = &eda.Component{
			Components: eda.Components{
				mph100imp40f.G_V_SP(6).Arrange(transform.Rotate(-90).Move(-8.9, -2.9-2.54)),
				mph100imp40f.G_V_SP(8).Arrange(transform.Move(0, -15.5)),
				mph100imp40f.G_V_SP(6).Arrange(transform.Rotate(90).Move(8.9, -2.9-2.54)),
			},
		}
	}

	pin := fanstel.BC833.Pads.Centers(transform.Identity)

	labelShift := path.Point{X: 2.54 / 0.8}
	labelScale := transform.Scale(0.8, 1.3)

	tracks := eda.TrackPaths(
		eda.Track{pad[0]}.DX(1.5).DY(-4.7).DX(1.6).DX(3.7).DY(-1.2).YX(pin[8]),
		eda.Track{pad[1]}.DY(-2.75).DX(5.3).DX(0.9).YX(pin[7]),
		eda.Track{pad[2]}.DY(-0.8).DX(5).XY(pin[6]),
		eda.Track{pad[3]}.YX(pin[1]),
		eda.Track{pad[4]}.XY(pin[2]),
		eda.Track{pad[5]}.XY(pin[3]),
		eda.Track{pad[6]}.DX(1.5).YX(pin[4]),

		eda.Track{pad[8]}.DY(2.1).X(-3.5).X(-2.1).DY(8).DY(1.1).XY(pin[21]),
		eda.Track{pad[9]}.DY(1.5).X(-3.2).X(-1.5).DY(8.2).DY(0.9).XY(pin[20]),
		eda.Track{pad[10]}.X(-0.9).YX(pin[9]),
		eda.Track{pad[11]}.X(-0.3).YX(pin[10]),

		eda.Track{pad[12]}.X(0.3).YX(pin[11]),
		eda.Track{pad[13]}.X(0.9).YX(pin[12]),
		eda.Track{pad[14]}.DY(1.5).X(3.2).X(1.5).YX(pin[13]),
		eda.Track{pad[15]}.DY(2.1).X(3.5).X(2.1).YX(pin[14]),

		eda.Track{pad[16]}.DX(-2.1).YX(pin[19]),
		eda.Track{pad[17]}.DX(-1.5).YX(pin[15]),
		eda.Track{pad[18]}.XY(pin[16]),
		eda.Track{pad[19]}.XY(pin[17]),
		eda.Track{pad[20]}.YX(pin[18]),
		eda.Track{pad[21]}.DY(-0.8).XY(pin[22]),
		eda.Track{pad[22]}.DY(-2.75).XY(pin[23]),
		eda.Track{pad[23]}.DX(-1.5).DY(-4.7).DX(-1.5).XY(pin[24]),
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

		revision += "s"
	}

	shiftedBoard := &eda.Component{
		Components: eda.Components{
			fanstel.BC833,

			header,

			pcbc.MountHole.Arrange(transform.Rotate(45).Move(-5, -10.5)),

			pcbc.MountHole.Arrange(transform.Rotate(-45).Move(5, -10.5)),

			pcbc.Logo.Arrange(transform.ScaleK(1.2).Move(0, -8.3)),

			eda.CenteredText(revision).Arrange(transform.Scale(0.75, 1).Move(-5, -8.1)),

			eda.CenteredText("BC833").Arrange(transform.ScaleK(2).Move(0, -10.4)),

			eda.CenteredText("nRF52833").Arrange(transform.ScaleK(1.5).Move(0, -12.2)),
		},

		Tracks: tracks,

		GroundTracks: eda.TrackPaths(
			eda.Track{pad[7]}.DX(2.1).YX(pin[5]),
			eda.Track{pin[5]}.DX(3).YX(pin[27]).XY(pin[25]).XY(pin[26]).YX(pin[28]).XY(pin[27]),
		),

		Marks: path.Paths{}.Append(
			font.ShiftedCenteredPaths(labelShift, leftLabels...).
				Apply(labelScale.Rotate(-90).Move(-7, -2.9)),
			font.ShiftedCenteredPaths(labelShift, centerLabels...).
				Apply(labelScale.Move(0, -13.65)),
			font.ShiftedCenteredPaths(labelShift, rightLabels...).
				Apply(labelScale.Rotate(90).Move(7, -2.9)),
		),
	}

	boardShift := 4.75
	boardCut := path.RoundRect(21, 24.6, 1)
	boardClears := path.Paths{
		path.Rect(21.5, 5.5).Apply(transform.Move(0, 9.7)),
	}

	if short {
		boardShift += 2.54
		boardCut = path.RoundRect(21, 19.5, 1)
		boardClears = nil
	}

	return &eda.Component{
		Clears: boardClears,

		Cuts: path.Paths{
			boardCut,
		},

		Components: eda.Components{
			shiftedBoard.Arrange(transform.Move(0, boardShift)),
		},
	}
}
