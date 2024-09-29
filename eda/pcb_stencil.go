package eda

import (
	"temnok/lab/lbrn"
	"temnok/lab/shape"
)

func (pcb *PCB) SaveStencil(filename string) error {
	p := lbrn.LightBurnProject{
		CutSetting: []lbrn.CutSetting{
			{
				Type:     "Cut",
				Index:    Param{"0"},
				Name:     Param{"C00"},
				Priority: Param{"1"},

				Speed:        Param{"400"},
				GlobalRepeat: Param{"30"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				TabsEnabled: Param{"1"},
				TabSize:     Param{"0.2"},
			},
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
		},
	}

	brush := shape.Circle(2)

	for _, mark := range pcb.stencilMarks {
		p.Shape = append(p.Shape, lbrn.NewPath(1, lbrnCenter, mark))
	}

	for _, hole := range pcb.stencilHoles {
		resizedHole := hole.Resize(-0.1)
		p.Shape = append(p.Shape, lbrn.NewPath(0, lbrnCenter, resizedHole).SetCutOrder(1))
		brush.IterateContour(resizedHole, pcb.bitmapTransform(), pcb.stencil.Set1)
	}

	for _, cut := range pcb.stencilCuts {
		p.Shape = append(p.Shape, lbrn.NewPathWithTabs(0, lbrnCenter, cut).SetCutOrder(2))
		brush.IterateContour(cut, pcb.bitmapTransform(), pcb.stencil.Set1)
	}

	return p.SaveToFile(filename)
}
