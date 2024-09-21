package eda

import (
	"temnok/lab/lbrn"
)

func (pcb *PCB) SaveStencil(filename string) error {
	p := lbrn.LightBurnProject{
		CutSetting: []lbrn.CutSetting{
			{
				Type:     "Cut",
				Index:    Param{"0"},
				Name:     Param{"C00"},
				Priority: Param{"0"},

				Speed:        Param{"400"},
				GlobalRepeat: Param{"10"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				TabsEnabled: Param{"1"},
				TabSize:     Param{"0.2"},
			},
		},
	}

	for _, cut := range pcb.stencilCuts {
		p.Shape = append(p.Shape, lbrn.NewPathWithTabs(0, lbrnCenter, cut))
	}

	for _, hole := range pcb.stencilHoles {
		p.Shape = append(p.Shape, lbrn.NewPathWithTabs(0, lbrnCenter, hole))
	}

	return p.SaveToFile(filename)
}
