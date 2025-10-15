// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/path"
)

var (
	alignLines = path.Paths{
		eda.LinearTrack(path.Point{-0.5, 0}, path.Point{0.5, 0}),
		eda.LinearTrack(path.Point{0, -0.5}, path.Point{0, 0.5}),
	}

	alignMark = &eda.Component{
		MarksWidth: 0.13,

		Nested: eda.Components{
			{
				AlignMarks: alignLines,
			},
			{
				Bottom: true,

				AlignMarks: alignLines,
			},
		},
	}
)
