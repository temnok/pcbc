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
				Name:     Param{Value: "Apertures"},
				Index:    Param{Value: "0"},
				Priority: Param{Value: "0"},

				Speed:        Param{Value: "200"},
				GlobalRepeat: Param{Value: "20"},

				MaxPower:    Param{Value: "90"},
				QPulseWidth: Param{Value: "200"},
				Frequency:   Param{Value: "20000"},
			},
		},
	}

	pcb.component.Visit(func(component *eda.Component) {
		t := component.Transform.Multiply(pcb.LbrnCenter)

		for _, pad := range component.Pads {
			p.Shape = append(p.Shape, lbrn.NewPath(0, t, pad))
		}

		for _, perforation := range component.Perforations {
			p.Shape = append(p.Shape, lbrn.NewPath(0, t, perforation))
		}
	})

	return p.SaveToFile(filename)
}
