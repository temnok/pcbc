// Copyright © 2025 Alex Temnok. All rights reserved.

package ebyte

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

// https://www.cdebyte.com/pdf-down.aspx?id=2587

var E73 *eda.Component

func init() {
	const (
		padInW  = 0.8
		padOutW = 0.8
		padW    = padInW + padOutW
		padH    = 0.8
		padStep = 1.27
		padRows = 10.0

		componentW = 13.0
		componentH = 18.0

		padShift  = (padInW - padOutW) / 2
		padVShift = (3.97 - 2.60) / 2
	)

	pad := path.RoundRect(padW, padH, 0.25)

	E73 = &eda.Component{
		Marks: path.Paths{path.Rect(componentW, componentH)},

		Pads: path.Join(
			pad.CloneXY(padRows, 0, padStep).Transform(transform.RotateDegrees(180).Move(-componentW/2-padShift, -padVShift)),
			pad.CloneXY(8, 0, padStep).Transform(transform.RotateDegrees(-90).Move(0, -componentH/2-padShift)),
			pad.CloneXY(padRows, 0, padStep).Transform(transform.Move(componentW/2+padShift, -padVShift)),
		),
	}
}
