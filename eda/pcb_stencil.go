package eda

import (
	"temnok/lab/lbrn"
	"temnok/lab/shape"
)

func (pcb *PCB) SaveStencil(filename string) error {
	p := lbrn.LightBurnProject{
		CutSetting: []lbrn.CutSetting{
			{
				Type:     "Scan",
				Index:    Param{"0"},
				Name:     Param{"Key"},
				Priority: Param{"0"},

				Speed: Param{"200"},

				MaxPower:    Param{"25"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				Interval: Param{"0.02"},
			},
			{
				Type:     "Cut",
				Index:    Param{"1"},
				Name:     Param{"Apertures"},
				Priority: Param{"1"},

				Speed:        Param{"400"},
				GlobalRepeat: Param{"40"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},
			},
			{
				Type:     "Cut",
				Index:    Param{"2"},
				Name:     Param{"Perimeter"},
				Priority: Param{"2"},

				Speed:        Param{"400"},
				GlobalRepeat: Param{"40"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				TabsEnabled: Param{"1"},
				TabSize:     Param{"0.2"},
			},
			{
				Type:     "Scan",
				Index:    Param{"3"},
				Name:     Param{"Clean"},
				Priority: Param{"3"},

				MaxPower:    Param{"50"},
				QPulseWidth: Param{"2"},
				Frequency:   Param{"280000"},

				Speed:      Param{"2000"},
				Interval:   Param{"0.01"},
				CrossHatch: Param{"1"},
			},
		},
		Shape: []*lbrn.Shape{
			lbrn.NewRect(3, lbrnCenter, 36, 46, 0),
		},
	}

	brush := shape.Circle(2)

	for _, mark := range pcb.stencilMarks {
		p.Shape = append(p.Shape, lbrn.NewPath(0, lbrnCenter, mark))
	}

	for _, hole := range pcb.stencilHoles {
		resizedHole := hole.Resize(-0.1)
		p.Shape = append(p.Shape, lbrn.NewPath(1, lbrnCenter, resizedHole))
		brush.IterateContour(resizedHole.Transform(pcb.bitmapTransform()), pcb.stencil.Set1)
	}

	for _, cut := range pcb.stencilCuts {
		p.Shape = append(p.Shape, lbrn.NewPathWithTabs(2, lbrnCenter, cut))
		brush.IterateContour(cut.Transform(pcb.bitmapTransform()), pcb.stencil.Set1)
	}

	return p.SaveToFile(filename)
}
