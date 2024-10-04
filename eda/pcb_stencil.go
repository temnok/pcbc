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
				Name:     Param{"Apertures"},
				Index:    Param{"0"},
				Priority: Param{"0"},

				Speed:        Param{"400"},
				GlobalRepeat: Param{"40"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},
			},
			{
				Type:     "Scan",
				Name:     Param{"Clean"},
				Index:    Param{"2"},
				Priority: Param{"2"},

				MaxPower:    Param{"50"},
				QPulseWidth: Param{"2"},
				Frequency:   Param{"280000"},

				Speed:      Param{"2000"},
				Interval:   Param{"0.01"},
				CrossHatch: Param{"1"},
			},
		},
		Shape: []*lbrn.Shape{
			lbrn.NewRect(1, lbrnCenter, 36, 46, 0),
		},
	}

	brush := shape.Circle(2)
	for _, hole := range pcb.stencilHoles {
		p.Shape = append(p.Shape, lbrn.NewPath(0, lbrnCenter, hole))
		brush.IterateContour(hole.Transform(pcb.bitmapTransform()), pcb.stencil.Set1)
	}

	return p.SaveToFile(filename)
}
