// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/lbrn"
)

func (pcb *PCB) SaveEtchPI() error {
	filename := pcb.SavePath + "etch-pi.lbrn"
	im := image.NewSingle(pcb.copper, color.White, color.Black)

	p := &lbrn.LightBurnProject{
		CutSettingImg: []*lbrn.CutSetting{
			{
				Type:     "Image",
				Name:     Param{Value: "Remove Adhesive"},
				Index:    Param{Value: "2"},
				Priority: Param{Value: "2"},

				MaxPower:    Param{Value: "25"},
				QPulseWidth: Param{Value: "1"},
				Frequency:   Param{Value: "650000"},

				Speed:            Param{Value: "500"},
				Interval:         Param{Value: "0.01"},
				DPI:              Param{Value: "2540"},
				UseDotCorrection: Param{Value: "1"},
				DotWidth:         Param{Value: "0.05"},

				NumPasses: Param{Value: "5"},

				CrossHatch: Param{Value: "1"},
				Angle:      Param{Value: "90"},

				Negative: Param{Value: "1"},
			},
		},
		CutSetting: []*lbrn.CutSetting{
			{
				Type:     "Scan",
				Name:     Param{Value: "Remove Kapton"},
				Index:    Param{Value: "0"},
				Priority: Param{Value: "0"},

				MaxPower:    Param{Value: "15"},
				QPulseWidth: Param{Value: "80"},
				Frequency:   Param{Value: "200000"},

				Speed:    Param{Value: "500"},
				Interval: Param{Value: "0.02"},
				DPI:      Param{Value: "1270"},

				NumPasses: Param{Value: "4"},

				CrossHatch: Param{Value: "1"},
			},
			{
				Type:     "Cut",
				Name:     Param{Value: "Cut"},
				Index:    Param{Value: "1"},
				Priority: Param{Value: "1"},

				MaxPower:    Param{Value: "80"},
				QPulseWidth: Param{Value: "80"},
				Frequency:   Param{Value: "40000"},

				Speed: Param{Value: "800"},

				NumPasses: Param{Value: "80"},
			},
			{
				Type:     "Scan",
				Name:     Param{Value: "Clean"},
				Index:    Param{Value: "3"},
				Priority: Param{Value: "3"},

				MaxPower:    Param{Value: "20"},
				QPulseWidth: Param{Value: "80"},
				Frequency:   Param{Value: "200000"},

				Speed:    Param{Value: "500"},
				Interval: Param{Value: "0.02"},
				DPI:      Param{Value: "1270"},

				NumPasses: Param{Value: "1"},

				CrossHatch: Param{Value: "1"},
			},
		},
		Shape: []*lbrn.Shape{
			lbrn.NewBitmap(2, pcb.lbrnBitmapScale(), im),
		},
	}

	p.Shape = append(p.Shape, lbrn.NewRect(0, pcb.lbrnCenterMove(), pcb.Width, pcb.Height))
	p.Shape = append(p.Shape, lbrn.NewRect(3, pcb.lbrnCenterMove(), pcb.Width, pcb.Height))

	pcb.component.Visit(func(component *eda.Component) {
		t := component.Transform.Multiply(pcb.lbrnCenterMove())

		for _, cut := range component.Cuts {
			p.Shape = append(p.Shape, lbrn.NewPathWithTabs(1, t, cut))
		}

		for _, hole := range component.Holes {
			p.Shape = append(p.Shape, lbrn.NewPath(1, t, hole))
		}

		for _, perforation := range component.Perforations {
			p.Shape = append(p.Shape, lbrn.NewPath(1, t, perforation))
		}
	})

	return p.SaveToFile(filename)
}
