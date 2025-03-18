// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"fmt"
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

				MaxPower:    Param{Value: "8"},
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
		Shape: []*lbrn.Shape{
			lbrn.NewRect(0, pcb.lbrnCenterMove(), pcb.Width, pcb.Height),
		},
	}

	for pass := 0; pass < 4; pass++ {
		i := 2 + pass*2
		a := Param{Value: fmt.Sprint((pass - 1) * 90)}

		p.CutSettingImg = append(p.CutSettingImg, &lbrn.CutSetting{
			Type:     "Image",
			Name:     Param{Value: fmt.Sprintf("Pass %v - Remove Adhesive", pass+1)},
			Index:    Param{Value: fmt.Sprint(i)},
			Priority: Param{Value: fmt.Sprint(i)},

			MaxPower:    Param{Value: "80"},
			QPulseWidth: Param{Value: "2"},
			Frequency:   Param{Value: "3000000"},

			Speed:            Param{Value: "400"},
			Interval:         Param{Value: "0.01"},
			DPI:              Param{Value: "2540"},
			UseDotCorrection: Param{Value: "1"},
			DotWidth:         Param{Value: "0.05"},

			Angle:     a,
			NumPasses: Param{Value: "1"},
			Negative:  Param{Value: "1"},
		})

		p.CutSetting = append(p.CutSetting, &lbrn.CutSetting{
			Type:     "Scan",
			Name:     Param{Value: fmt.Sprintf("Pass %v - Clean", pass+1)},
			Index:    Param{Value: fmt.Sprint(i + 1)},
			Priority: Param{Value: fmt.Sprint(i + 1)},

			MaxPower:    Param{Value: "8"},
			QPulseWidth: Param{Value: "200"},
			Frequency:   Param{Value: "20000"},

			Speed:    Param{Value: "400"},
			Interval: Param{Value: "0.02"},
			DPI:      Param{Value: "1270"},

			Angle:     a,
			NumPasses: Param{Value: "1"},
		})

		p.Shape = append(p.Shape, lbrn.NewBitmapShape(i, pcb.lbrnBitmapScale(), bm))
		p.Shape = append(p.Shape, lbrn.NewRect(i+1, pcb.lbrnCenterMove(), pcb.Width, pcb.Height))
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
