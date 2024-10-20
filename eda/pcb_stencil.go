package eda

import (
	"temnok/pcbc/lbrn"
	"temnok/pcbc/transform"
)

const StencilShrink = 0.1

func (pcb *PCB) SaveStencil(filename string) error {
	center := transform.Move(pcb.lbrnCenter.X, pcb.lbrnCenter.Y)

	p := lbrn.LightBurnProject{
		CutSetting: []lbrn.CutSetting{
			{
				Type:     "Cut",
				Name:     Param{"Apertures"},
				Index:    Param{"0"},
				Priority: Param{"0"},

				Speed:        Param{"200"},
				GlobalRepeat: Param{"20"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},
			},
		},
	}

	pads := pcb.component.Pads.Resize(-StencilShrink)
	for _, hole := range pads {
		p.Shape = append(p.Shape, lbrn.NewPath(0, center, hole))
	}

	return p.SaveToFile(filename)
}
