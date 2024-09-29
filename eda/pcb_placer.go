package eda

import (
	"temnok/lab/lbrn"
)

func (pcb *PCB) SavePlacer(filename string) error {
	p := lbrn.LightBurnProject{
		CutSetting: []lbrn.CutSetting{
			{
				Type:     "Scan",
				Index:    Param{"1"},
				Name:     Param{"C01"},
				Priority: Param{"0"},

				Speed: Param{"200"},

				MaxPower:    Param{"25"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				Interval: Param{"0.02"},
			},
			{
				Type:     "Cut",
				Index:    Param{"0"},
				Name:     Param{"C00"},
				Priority: Param{"1"},

				Speed:        Param{"200"},
				GlobalRepeat: Param{"15"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				TabsEnabled: Param{"1"},
				TabSize:     Param{"0.1"},
			},
		},
	}

	for _, mark := range pcb.placerMarks {
		p.Shape = append(p.Shape, lbrn.NewPath(1, lbrnCenter, mark))
	}

	for _, hole := range pcb.placerHoles {
		p.Shape = append(p.Shape, lbrn.NewPath(0, lbrnCenter, hole).SetCutOrder(1))
	}

	for _, cut := range pcb.placerCuts {
		p.Shape = append(p.Shape, lbrn.NewPathWithTabs(0, lbrnCenter, cut).SetCutOrder(2))
	}

	return p.SaveToFile(filename)
}
