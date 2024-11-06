package minewsemi

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

// https://store.minewsemi.com/wp-content/uploads/2024/03/MS88SF2-nRF52840_Datasheet_K_EN-1.pdf

var MS88SF2 *eda.Component

func init() {
	const (
		padInW       = 1.3
		padOutW      = 0.7
		padW         = padInW + padOutW
		padH         = 0.8
		padGap       = 0.3
		padStep      = padH + padGap
		padRows      = 14
		padBottomGap = 0.91

		componentW = 17.4
		componentH = 23.2

		padHShift = (componentW - (padInW - padOutW)) / 2
		padVShift = (componentH-(padRows*padStep-padGap))/2 - padBottomGap
	)

	pad := path.RoundRect(padW, padH, 0.3)

	MS88SF2 = &eda.Component{
		Marks: path.Rect(componentW, componentH),

		Pads: path.Join(
			path.Paths{nil}, // skip pad #0
			pad.Clone(padRows, 0, padStep).Apply(transform.Rotate(180).Move(-padHShift, -padVShift)),
			pad.Clone(padRows, 0, padStep).Apply(transform.Move(padHShift, -padVShift)),
		),
	}
}
