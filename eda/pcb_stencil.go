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
				Name:     Param{Value: "Apertures"},
				Index:    Param{Value: "0"},
				Priority: Param{Value: "0"},

				Speed:        Param{Value: "200"},
				GlobalRepeat: Param{Value: "20"},

				MaxPower:    Param{Value: "90"},
				QPulseWidth: Param{Value: "200"},
				Frequency:   Param{Value: "20000"},
			},
		},
	}

	pcb.component.Visit(func(component *Component) {
		t := component.Transform.Multiply(center)

		for _, pad := range component.Pads {
			p.Shape = append(p.Shape, lbrn.NewPath(0, t, pad))
		}
	})

	return p.SaveToFile(filename)
}
