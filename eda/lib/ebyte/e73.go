package ebyte

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

// https://www.cdebyte.com/pdf-down.aspx?id=2587

var E73 *eda.Component

func init() {
	const (
		padInW       = 0.7
		padOutW      = 0.7
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

	pad := path.RoundRect(padW, padH, 0.2)

	E73 = &eda.Component{
		Marks: path.Rect(componentW, componentH),

		Pads: path.Join(
			pad.Clone(padRows, 0, padStep).Apply(transform.Rotate(180).Move(-padHShift, -padVShift)),
			pad.Clone(8, 0, padStep).Apply(transform.Rotate(-90).Move(0, -componentH/2)),
			pad.Clone(padRows, 0, padStep).Apply(transform.Move(padHShift, -padVShift)),
		),
	}
}
