package ebyte

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

// https://www.cdebyte.com/pdf-down.aspx?id=2587

var E73 *lib.Component

func init() {
	const (
		padInW       = 0.5
		padOutW      = 0.5
		padW         = padInW + padOutW
		padH         = 0.8
		padStep      = 1.27
		padGap       = padStep - padH
		padRows      = 10
		padBottomGap = 2.6 - padH/2

		componentW = 13
		componentH = 18

		padHShift = (componentW - (padInW - padOutW)) / 2
		padVShift = (componentH-(padRows*padStep-padGap))/2 - padBottomGap
	)

	pad := path.RoundRect(padW, padH, 0.1)

	E73 = &lib.Component{
		Marks: path.Strokes{
			0.1: path.Paths{path.Rect(componentW, componentH)},
		},

		Pads: path.Paths.Append(
			pad.Clone(padRows, 0, padStep).Transform(geom.MoveXY(-padHShift, -padVShift).RotateD(180)),
			pad.Clone(8, 0, padStep).Transform(geom.MoveXY(0, -componentH/2).RotateD(-90)),
			pad.Clone(padRows, 0, padStep).Transform(geom.MoveXY(padHShift, -padVShift)),
		),
	}
}
