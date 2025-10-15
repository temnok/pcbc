// Copyright © 2025 Alex Temnok. All rights reserved.

package mph100imp40f

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/path"
)

func G_V_SP(n int) *eda.Component {
	const step = 2.54

	return &eda.Component{
		Pads:  path.Circle(1.8).CloneXY(n, step, 0),
		Marks: path.CutRect(step, step, 0.3).CloneXY(n, step, 0),
	}
}
