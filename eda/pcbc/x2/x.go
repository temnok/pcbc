package x2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/pkg/smd"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

func X0402(title, upperText, lowerText string) *lib.Component {
	return x(smd.I0402, title, upperText, lowerText, false)
}

func X0402WithGround(title, upperText, lowerText string) *lib.Component {
	return x(smd.I0402, title, upperText, lowerText, true)
}

func X0603(title, upperText, lowerText string) *lib.Component {
	return x(smd.I0603, title, upperText, lowerText, false)
}

func X0603WithGround(title, upperText, lowerText string) *lib.Component {
	return x(smd.I0603, title, upperText, lowerText, true)
}

func x(x *lib.Component, title, upperText, lowerText string, ground bool) *lib.Component {
	groundLabel := ""
	if ground {
		groundLabel = "GND"
	}

	comp := (&lib.Component{
		Cuts: path.Paths{
			path.RoundRect(9.75, 4.75, 1),
		},
		Marks: path.Strokes{
			font.Bold: path.Paths{}.Append(
				font.StringPaths(title, font.AlignCenter).
					Transform(geom.MoveXY(-3.1, 1.3).ScaleXY(1, 2.2)),
			),

			font.Normal: path.Paths{}.Append(
				font.StringPaths(groundLabel, font.AlignCenter).
					Transform(geom.MoveXY(0.8, -0.9).RotateD(-90).ScaleXY(1, 0.75)),
				font.StringPaths(upperText, font.AlignCenter).
					Transform(geom.MoveXY(2.5, 1.8).ScaleXY(0.75, 1)),
				font.StringPaths(lowerText, font.AlignCenter).
					Transform(geom.MoveXY(2.5, -1.8).ScaleXY(0.75, 1)),
			),
		}.Append(
			pcbc.Logo.Transform(geom.MoveXY(4.3, 0).ScaleK(0.5)),
			pcbc.TmnkTech.Transform(geom.MoveXY(4.3, -1).ScaleK(0.4)),
		),
		Components: lib.Components{
			{
				Transform: geom.MoveXY(-0.5, 1.3),
				Components: lib.Components{
					x,
				},
			},
			{
				Transform: geom.MoveXY(-2.1, -0.85),
				Components: lib.Components{
					mph100imp40f.G_V_SP_x2,
				},
			},
			{
				Transform: geom.MoveXY(2.5, 0),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
		},
	}).Squash()

	pad := comp.Pads.Centers()

	tracks := eda.TrackPaths(
		eda.Track{pad[0]}.DX(-1).XY(pad[2]),
	)

	groundTracks := eda.TrackPaths(
		eda.Track{pad[1]}.YX(pad[3]),
		eda.Track{pad[1]}.XY(pad[6]),
		eda.Track{pad[3]}.DY(-0.75).XY(pad[8]),

		eda.Track{pad[6]}.DX(-0.5).DY(0.5),
		eda.Track{pad[8]}.DX(-0.5).DY(-0.5),
	)

	if !ground {
		tracks = append(tracks, groundTracks[0])
		groundTracks = groundTracks[3:]
	}

	comp.Tracks.Append(path.Strokes{
		0: tracks,
	})
	comp.GroundTracks.Append(path.Strokes{
		0: groundTracks,
	})

	return comp
}
