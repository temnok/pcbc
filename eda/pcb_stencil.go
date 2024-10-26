package eda

import (
	"temnok/pcbc/lbrn"
	"temnok/pcbc/transform"
)

func (pcb *PCB) SaveStencil() error {
	filename := pcb.savePath + "stencil.lbrn"

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

	for _, hole := range pcb.component.Pads {
		p.Shape = append(p.Shape, lbrn.NewPath(0, center, hole))
	}

	return p.SaveToFile(filename)
}
