// Copyright © 2025 Alex Temnok. All rights reserved.

package e73

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/ebyte"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	chip = ebyte.E73.Arrange(transform.RotateDegrees(-90).Move(5.2, 0))

	pin = chip.PadCenters()

	headers = &eda.Component{
		Nested: eda.Components{
			greenconn.CSCC118(27, true, []string{
				"P111", "P110", "P003", "P028", "GND", "P113", "P002", "P029", "P031", "P030",
				"P000", "P001", "P005", "P109", "VDD", "VDDH" /*"GND",*/, "DCCH",
				"P018", "VBUS", "D-", "D+", "P013", "P024", "SWD", "SWC", "P009", "P010",
			}).Arrange(transform.Move(-8.2, 0)),
		},
	}

	pad = headers.PadCenters()

	mountHoles = &eda.Component{
		Nested: eda.Components{
			boards.MountHole.Arrange(transform.RotateDegrees(-45).Move(7.5, 10)),
			boards.MountHole.Arrange(transform.RotateDegrees(45).Move(7.5, -10)),
		},
	}

	Board_nRF52840 = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(22.4, 28, 1),
		},

		Nested: eda.Components{
			chip,
			headers,
			mountHoles,
		},

		Tracks: eda.Tracks(
			eda.Track{pin[0]}.DY(5).YX(pad[0]),
			eda.Track{pin[1]}.DY(1.2).DX(0.6).DY(3.5).YX(pad[1]),
			eda.Track{pin[2]}.DY(0.5).YX(pad[2]),
			eda.Track{pin[3]}.DY(0.5).YX(pad[3]),
			eda.Track{pin[4]}.DY(0.5).Y(pad[4].Y+0.3).DX(-10.9).YX(pad[4]),
			eda.Track{pin[5]}.DY(0.5).Y(pad[5].Y+0.6).DX(-9.7).YX(pad[5]),
			eda.Track{pin[6]}.DY(0.5).Y(pad[6].Y+0.9).DX(-8.5).YX(pad[6]),
			eda.Track{pin[7]}.DY(-1).DY(-1).DX(-4.1).DY(0.8).DX(-3.3).YX(pad[9]),
			eda.Track{pin[8]}.DY(-1).DY(-0.4).DX(-2.5).DY(0.8).DX(-3.5).YX(pad[8]),
			eda.Track{pin[9]}.DX(-4.2).DY(0).YX(pad[7]),

			eda.Track{pin[10]}.DX(-1).DY(0).YX(pad[10]),
			eda.Track{pin[11]}.DX(-1).DY(0).YX(pad[11]),
			eda.Track{pin[12]}.DX(-1).DY(0).YX(pad[12]),
			eda.Track{pin[13]}.DX(-1).DY(0).YX(pad[13]),
			eda.Track{pin[14]}.DX(-1).DY(0).YX(pad[14]),
			eda.Track{pin[16]}.DX(-1).DY(0).YX(pad[15]),
			eda.Track{pin[17]}.DX(-1).DY(0).YX(pad[16]),

			eda.Track{pin[18]}.DX(-4.2).DY(0).YX(pad[19]),
			eda.Track{pin[19]}.DY(1).DY(0.4).DX(-2.5).DY(-0.8).DX(-3.5).YX(pad[18]),
			eda.Track{pin[20]}.DY(1).DY(1).DX(-4.1).DY(-0.8).DX(-3.3).YX(pad[17]),
			eda.Track{pin[21]}.DY(-0.5).Y(pad[20].Y-0.9).DX(-8.5).YX(pad[20]),
			eda.Track{pin[22]}.DY(-0.5).Y(pad[21].Y-0.6).DX(-9.7).YX(pad[21]),
			eda.Track{pin[23]}.DY(-0.5).Y(pad[22].Y-0.3).DX(-10.9).YX(pad[22]),
			eda.Track{pin[24]}.DY(-0.5).YX(pad[23]),
			eda.Track{pin[25]}.DY(-0.5).YX(pad[24]),
			eda.Track{pin[26]}.DY(-1.2).DX(0.6).DY(-3.5).YX(pad[25]),
			eda.Track{pin[27]}.DY(-5).YX(pad[26]),
		),
	}
)
