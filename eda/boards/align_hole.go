// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

/*

1x1: 76x44
1x2: 36x44
1x3: 23x44
2x2: 36x20
2x3: 23x20
2x4: 16x20
3x3: 23x12
3x4: 16x12
3x5: 12x12

*/

var AlignHole = &eda.Component{
	Nested: eda.Components{
		{
			AlignCuts: path.Paths{
				path.Circle(1.43),
			},
		},

		{
			Bottom: true,

			AlignCuts: path.Paths{
				path.Circle(2.4),
			},
		},
	},
}
