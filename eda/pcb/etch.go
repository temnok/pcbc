// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/lbrn"
)

type Param = lbrn.Param

func (pcb *PCB) SaveEtch() error {
	if pcb.SaveEtchOverride != nil {
		return pcb.SaveEtchOverride()
	}

	return pcb.SaveEtchFR4()
}

func (pcb *PCB) SaveEtchFR4() error {
	filename := pcb.SavePath + "etch.lbrn"
	im := image.NewSingle(pcb.copper, color.White, color.Black)
	bm := lbrn.NewBase64Bitmap(im)

	p := &lbrn.LightBurnProject{
		CutSettingImg: []*lbrn.CutSetting{
			{
				Type:     "Image",
				Name:     Param{Value: "Etch"},
				Index:    Param{Value: "0"},
				Priority: Param{Value: "0"},

				MaxPower:    Param{Value: "20"},
				QPulseWidth: Param{Value: "200"},
				Frequency:   Param{Value: "20000"},

				Speed:            Param{Value: "600"},
				Interval:         Param{Value: "0.01"},
				DPI:              Param{Value: "2540"},
				UseDotCorrection: Param{Value: "1"},
				DotWidth:         Param{Value: "0.05"},

				Negative: Param{Value: "1"},
			},
			{
				Type:     "Image",
				Name:     Param{Value: "Clean 1"},
				Index:    Param{Value: "1"},
				Priority: Param{Value: "1"},

				MaxPower:    Param{Value: "50"},
				QPulseWidth: Param{Value: "2"},
				Frequency:   Param{Value: "280000"},

				Speed:            Param{Value: "2000"},
				Interval:         Param{Value: "0.01"},
				DPI:              Param{Value: "2540"},
				UseDotCorrection: Param{Value: "1"},
				DotWidth:         Param{Value: "0.15"},

				Negative: Param{Value: "1"},
			},
			{
				Type:     "Image",
				Name:     Param{Value: "Clean 2"},
				Index:    Param{Value: "3"},
				Priority: Param{Value: "3"},

				MaxPower:    Param{Value: "50"},
				QPulseWidth: Param{Value: "2"},
				Frequency:   Param{Value: "280000"},

				Speed:            Param{Value: "2000"},
				Interval:         Param{Value: "0.01"},
				DPI:              Param{Value: "2540"},
				UseDotCorrection: Param{Value: "1"},
				DotWidth:         Param{Value: "0.15"},

				Angle:    Param{Value: "90"},
				Negative: Param{Value: "1"},
			},
		},
		CutSetting: []*lbrn.CutSetting{
			{
				Type:     "Cut",
				Name:     Param{Value: "FR4 Cut"},
				Index:    Param{Value: "2"},
				Priority: Param{Value: "2"},

				Speed:        Param{Value: "100"},
				GlobalRepeat: Param{Value: "50"},

				MaxPower:    Param{Value: "90"},
				QPulseWidth: Param{Value: "200"},
				Frequency:   Param{Value: "20000"},

				TabsEnabled: Param{Value: "1"},
				TabSize:     Param{Value: "0.1"},
			},
		},
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShape(0, pcb.lbrnBitmapScale(), bm),
			lbrn.NewBitmapShape(1, pcb.lbrnBitmapScale(), bm),
			lbrn.NewBitmapShape(3, pcb.lbrnBitmapScale(), bm),
		},
	}

	pcb.component.Visit(func(component *eda.Component) {
		t := component.Transform.Multiply(pcb.lbrnCenterMove())

		for _, cut := range component.Cuts {
			p.Shape = append(p.Shape, lbrn.NewPathWithTabs(2, t, cut))
		}

		for _, hole := range component.Holes {
			p.Shape = append(p.Shape, lbrn.NewPath(2, t, hole))
		}

		for _, perforation := range component.Perforations {
			p.Shape = append(p.Shape, lbrn.NewPath(2, t, perforation))
		}
	})

	return p.SaveToFile(filename)
}
