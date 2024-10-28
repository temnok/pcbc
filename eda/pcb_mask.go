package eda

import (
	"image/color"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/transform"
)

var maskCutSettings = []lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     Param{"Silk"},
		Index:    Param{"0"},
		Priority: Param{"0"},

		MaxPower:    Param{"6"},
		QPulseWidth: Param{"80"},
		Frequency:   Param{"2000000"},

		NumPasses:        Param{"5"},
		Speed:            Param{"1000"},
		Interval:         Param{"0.02"},
		DPI:              Param{"1270"},
		UseDotCorrection: Param{"1"},
		DotWidth:         Param{"0.05"},

		CrossHatch: Param{"1"},
		Angle:      Param{"90"},
	},
	{
		Type:     "Image",
		Name:     Param{"Mask 1"},
		Index:    Param{"1"},
		Priority: Param{"1"},

		MaxPower:    Param{"10"},
		QPulseWidth: Param{"80"},
		Frequency:   Param{"2000000"},

		NumPasses:        Param{"5"},
		Speed:            Param{"500"},
		Interval:         Param{"0.01"},
		DPI:              Param{"2540"},
		UseDotCorrection: Param{"1"},
		DotWidth:         Param{"0.05"},

		CrossHatch: Param{"1"},
		Angle:      Param{"90"},
	},
	{
		Type:     "Image",
		Name:     Param{"Mask 2"},
		Index:    Param{"2"},
		Priority: Param{"2"},

		MaxPower:    Param{"20"},
		QPulseWidth: Param{"80"},
		Frequency:   Param{"2000000"},

		NumPasses:        Param{"5"},
		Speed:            Param{"500"},
		Interval:         Param{"0.01"},
		DPI:              Param{"2540"},
		UseDotCorrection: Param{"1"},
		DotWidth:         Param{"0.05"},

		CrossHatch: Param{"1"},
		Angle:      Param{"90"},
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
