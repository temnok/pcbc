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
	bm := lbrn.NewBase64Bitmap(im)

	p := &lbrn.LightBurnProject{
		CutSetting: []*lbrn.CutSetting{
			{
				Type:     "Scan",
				Name:     Param{Value: "Remove BPI"},
				Index:    Param{Value: "0"},
				Priority: Param{Value: "0"},

				MaxPower:    Param{Value: "10"},
				QPulseWidth: Param{Value: "200"},
				Frequency:   Param{Value: "20000"},

				Speed:    Param{Value: "400"},
				Interval: Param{Value: "0.02"},
				DPI:      Param{Value: "1270"},

				NumPasses: Param{Value: "3"},

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
		},
		CutSettingImg: []*lbrn.CutSetting{
			{
				Type:     "Image",
				Name:     Param{Value: "Remove Adhesive"},
				Index:    Param{Value: "2"},
				Priority: Param{Value: "2"},

				MaxPower:    Param{Value: "60"},
				QPulseWidth: Param{Value: "80"},
				Frequency:   Param{Value: "2000000"},

				Speed:            Param{Value: "500"},
				Interval:         Param{Value: "0.01"},
				DPI:              Param{Value: "2540"},
				UseDotCorrection: Param{Value: "1"},
				DotWidth:         Param{Value: "0.05"},

				//CrossHatch: Param{Value: "1"},
				Angle: Param{Value: "90"},

				NumPasses: Param{Value: "10"},
				Negative:  Param{Value: "1"},

				DitherMode:  Param{Value: "3dslice"},
				CleanupPass: &Param{Value: "1"},

				SubLayer: &lbrn.SubLayer{
					Type:      "Scan",
					Index:     "1",
					IsCleanup: Param{Value: "1"},
					FloodFill: Param{Value: "1"},

					Speed:       Param{Value: "400"},
					MaxPower:    Param{Value: "10"},
					Frequency:   Param{Value: "20000"},
					QPulseWidth: Param{Value: "200"},

					Interval: Param{Value: "0.02"},

					CrossHatch:   Param{Value: "1"},
					Angle:        Param{Value: "90"},
					AnglePerPass: Param{Value: "90"},
				},
			},
		},
		Shape: []*lbrn.Shape{
			lbrn.NewRect(0, pcb.lbrnCenterMove(), pcb.Width, pcb.Height),
			lbrn.NewBitmapShape(2, pcb.lbrnBitmapScale(), bm),
		},
	}

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
