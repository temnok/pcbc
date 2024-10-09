package eda

import (
	"temnok/pcbc/lbrn"
)

func (pcb *PCB) SaveStencil(filename string) error {
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

	for _, hole := range pcb.apertures {
		p.Shape = append(p.Shape, lbrn.NewPath(0, lbrnCenter, hole))
	}

	return p.SaveToFile(filename)
}
