// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var (
	alignLines = path.Paths{
		eda.LinearTrack(path.Point{-0.5, 0}, path.Point{0.5, 0}),
		eda.LinearTrack(path.Point{0, -0.5}, path.Point{0, 0.5}),
	}

	AlignTopMark = &eda.Component{
		AlignMarks: alignLines,
	}

	AlignBottomMark = &eda.Component{
		Bottom: true,

		AlignMarks: alignLines,
	}

	AlignMark = &eda.Component{
		Nested: eda.Components{
			AlignTopMark,
			AlignBottomMark,
		},
	}
)
