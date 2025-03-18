// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/lbrn"
)

var maskCutSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     Param{Value: "Silk"},
		Index:    Param{Value: "0"},
		Priority: Param{Value: "0"},

		MaxPower:    Param{Value: "5"},
		QPulseWidth: Param{Value: "200"},
		Frequency:   Param{Value: "20000"},

		NumPasses: Param{Value: "1"},
		Speed:     Param{Value: "800"},
		Interval:  Param{Value: "0.02"},
		DPI:       Param{Value: "1270"},

		// Making positive default -- negative is much slower!
		//Negative: Param{Value: "1"},

		CrossHatch: Param{Value: "1"},
		Angle:      Param{Value: "-90"},

		UseDotCorrection: Param{Value: "1"},
		DotWidth:         Param{Value: "0.05"},
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
	filename := pcb.SavePath + "mask.lbrn"
	silk := image.NewSingle(pcb.silk, color.White, color.Black)
	mask := image.NewSingle(pcb.mask, color.Transparent, color.Black)
	maskBM := lbrn.NewBase64Bitmap(mask)

	p := &lbrn.LightBurnProject{
		CutSettingImg: maskCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShapeFromImage(0, pcb.lbrnBitmapScale(), silk),
			lbrn.NewBitmapShape(1, pcb.lbrnBitmapScale(), maskBM),
			lbrn.NewBitmapShape(2, pcb.lbrnBitmapScale(), maskBM),
		},
	}

	pcb.addMaskPerforations(p)

	return p.SaveToFile(filename)
}

func (pcb *PCB) addMaskPerforations(p *lbrn.LightBurnProject) {
	pcb.component.Visit(func(component *eda.Component) {
		t := component.Transform.Multiply(pcb.lbrnCenterMove())

		for _, hole := range component.Perforations {
			p.Shape = append(p.Shape, lbrn.NewPath(3, t, hole))
		}
	})

	p.CutSetting = []*lbrn.CutSetting{
		{
			Type:     "Cut",
			Name:     Param{Value: "Perforation"},
			Index:    Param{Value: "3"},
			Priority: Param{Value: "3"},

			Speed:        Param{Value: "100"},
			GlobalRepeat: Param{Value: "30"},

			MaxPower:    Param{Value: "90"},
			QPulseWidth: Param{Value: "200"},
			Frequency:   Param{Value: "20000"},
		},
	}
}
