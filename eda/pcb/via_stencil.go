// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/lbrn"
)

var viaStencilSettings = []*lbrn.CutSetting{
	{
		Type:     "Cut",
		Name:     &lbrn.Param{Value: "Al200um Cut"},
		Index:    &lbrn.Param{Value: "0"},
		Priority: &lbrn.Param{Value: "0"},

		MaxPower:    &lbrn.Param{Value: "90"},
		QPulseWidth: &lbrn.Param{Value: "200"},
		Frequency:   &lbrn.Param{Value: "20000"},

		NumPasses:    &lbrn.Param{Value: "1"},
		GlobalRepeat: &lbrn.Param{Value: "50"},
		Speed:        &lbrn.Param{Value: "800"},
	},
}

func SaveViaStencil(config *config.Config, component *eda.Component) error {
	p := &lbrn.LightBurnProject{
		UIPrefs:    lbrn.UIPrefsDefaults,
		CutSetting: viaStencilSettings,
	}

	nonEmpty := renderViaStencil(config, component, p)
	if !nonEmpty {
		return nil
	}

	return p.SaveToFile(config.SavePath + "via-stencil.lbrn")
}

func renderViaStencil(config *config.Config, component *eda.Component, p *lbrn.LightBurnProject) bool {
	hasVias := false

	component.Visit(func(c *eda.Component) {
		hasVias = hasVias || len(c.Vias) > 0

		t := c.Transform.Multiply(config.LbrnCenterMove())

		for _, via := range c.Vias {
			p.Shape = append(p.Shape, lbrn.NewPath(0, t, via))
		}

		if c.CutsOuter {
			for _, cut := range c.Cuts {
				p.Shape = append(p.Shape, lbrn.NewPath(0, t, cut))
			}
		}
	})

	return hasVias
}
