// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/lbrn"
)

func (pcb *PCB) SaveStencil() error {
	filename := pcb.SavePath + "stencil.lbrn"

	p := lbrn.LightBurnProject{
		CutSetting: []*lbrn.CutSetting{
			{
				Type:     "Cut",
				Name:     param{Value: "Apertures"},
				Index:    param{Value: "0"},
				Priority: param{Value: "0"},

				Speed:        param{Value: "200"},
				GlobalRepeat: param{Value: "20"},

				MaxPower:    param{Value: "90"},
				QPulseWidth: param{Value: "200"},
				Frequency:   param{Value: "20000"},
			},
		},
	}

	pcb.component.Visit(func(component *eda.Component) {
		t := component.Transform.Multiply(pcb.lbrnCenterMove())

		for _, pad := range component.Pads {
			p.Shape = append(p.Shape, lbrn.NewPath(0, t, pad))
		}

		for _, perforation := range component.Perforations {
			p.Shape = append(p.Shape, lbrn.NewPath(0, t, perforation))
		}
	})

	return p.SaveToFile(filename)
}
