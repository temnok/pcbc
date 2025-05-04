// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/lbrn"
)

func SaveStencil(config *config.Config, component *eda.Component) error {
	filename := config.SavePath + "stencil.lbrn"

	p := lbrn.LightBurnProject{
		CutSetting: []*lbrn.CutSetting{
			{
				Type:     "Cut",
				Name:     lbrn.Param{Value: "Apertures"},
				Index:    lbrn.Param{Value: "0"},
				Priority: lbrn.Param{Value: "0"},

				Speed:        lbrn.Param{Value: "200"},
				GlobalRepeat: lbrn.Param{Value: "20"},

				MaxPower:    lbrn.Param{Value: "90"},
				QPulseWidth: lbrn.Param{Value: "200"},
				Frequency:   lbrn.Param{Value: "20000"},
			},
		},
	}

	component.Visit(func(component *eda.Component) {
		t := component.Transform.Multiply(config.LbrnCenterMove())

		for _, pad := range component.Pads {
			p.Shape = append(p.Shape, lbrn.NewPath(0, t, pad))
		}

		for _, perforation := range component.Perforations {
			p.Shape = append(p.Shape, lbrn.NewPath(0, t, perforation))
		}
	})

	return p.SaveToFile(filename)
}
