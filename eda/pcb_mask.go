// Copyright © 2025 Alex Temnok. All rights reserved.

package eda

import (
	"image/color"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/transform"
)

var maskCutSettings = []lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     Param{Value: "Silk"},
		Index:    Param{Value: "0"},
		Priority: Param{Value: "0"},

		MaxPower:    Param{Value: "6"},
		QPulseWidth: Param{Value: "80"},
		Frequency:   Param{Value: "2000000"},

		NumPasses:        Param{Value: "5"},
		Speed:            Param{Value: "1000"},
		Interval:         Param{Value: "0.02"},
		DPI:              Param{Value: "1270"},
		UseDotCorrection: Param{Value: "1"},
		DotWidth:         Param{Value: "0.05"},

		CrossHatch: Param{Value: "1"},
		Angle:      Param{Value: "90"},
	},
	{
		Type:     "Image",
		Name:     Param{Value: "Mask 1"},
		Index:    Param{Value: "1"},
		Priority: Param{Value: "1"},

		MaxPower:    Param{Value: "10"},
		QPulseWidth: Param{Value: "80"},
		Frequency:   Param{Value: "2000000"},

		NumPasses:        Param{Value: "5"},
		Speed:            Param{Value: "500"},
		Interval:         Param{Value: "0.01"},
		DPI:              Param{Value: "2540"},
		UseDotCorrection: Param{Value: "1"},
		DotWidth:         Param{Value: "0.05"},

		CrossHatch: Param{Value: "1"},
		Angle:      Param{Value: "90"},
	},
	{
		Type:     "Image",
		Name:     Param{Value: "Mask 2"},
		Index:    Param{Value: "2"},
		Priority: Param{Value: "2"},

		MaxPower:    Param{Value: "20"},
		QPulseWidth: Param{Value: "80"},
		Frequency:   Param{Value: "2000000"},

		NumPasses:        Param{Value: "5"},
		Speed:            Param{Value: "500"},
		Interval:         Param{Value: "0.01"},
		DPI:              Param{Value: "2540"},
		UseDotCorrection: Param{Value: "1"},
		DotWidth:         Param{Value: "0.05"},

		CrossHatch: Param{Value: "1"},
		Angle:      Param{Value: "90"},
	},
}

func (pcb *PCB) SaveMask() error {
	filename := pcb.savePath + "mask.lbrn"
	silk := pcb.silk.ToImage(color.White, color.Black)
	mask := pcb.mask.ToImage(color.Transparent, color.Black)

	bitmapTransform := transform.ScaleK(1/resolution).
		Move(pcb.lbrnCenter.X, pcb.lbrnCenter.Y)

	p := lbrn.LightBurnProject{
		CutSettingImg: maskCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmap(0, bitmapTransform, silk),
			lbrn.NewBitmap(1, bitmapTransform, mask),
			lbrn.NewBitmap(2, bitmapTransform, mask),
		},
	}

	return p.SaveToFile(filename)
}
