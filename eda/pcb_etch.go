package eda

import (
	"image/color"
	"temnok/lab/geom"
	"temnok/lab/lbrn"
)

type Param = lbrn.Param

var (
	lbrnCenter = geom.Move(XY{55, 55})

	holders = []XY{
		{-14.5, 19.5},
		{14, 19},
		{-14, -19},
		{14, -19},
	}

	holdersR = 1.0
)

func (pcb *PCB) SaveEtch(filename string) error {
	im := pcb.cu.ToImage(color.Black, color.White)

	k := 1 / pcb.scale

	p := lbrn.LightBurnProject{
		CutSettingImg: []lbrn.CutSetting{
			{
				Type:     "Image",
				Index:    Param{"0"},
				Name:     Param{"C00"},
				Priority: Param{"0"},

				MaxPower:    Param{"20"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				Speed:            Param{"600"},
				Interval:         Param{"0.01"},
				DPI:              Param{"2540"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.05"},
			},
			{
				Type:     "Image",
				Index:    Param{"1"},
				Name:     Param{"C01"},
				Priority: Param{"2"},

				MaxPower:    Param{"50"},
				QPulseWidth: Param{"2"},
				Frequency:   Param{"280000"},

				Speed:            Param{"2000"},
				Interval:         Param{"0.01"},
				DPI:              Param{"2540"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.05"},
			},
		},
		CutSetting: []lbrn.CutSetting{
			{
				Type:     "Cut",
				Index:    Param{"2"},
				Name:     Param{"C02"},
				Priority: Param{"1"},

				Speed:        Param{"100"},
				GlobalRepeat: Param{"100"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				TabsEnabled: Param{"1"},
				TabSize:     Param{"0.2"},
			},
		},
		Shape: []lbrn.Shape{
			lbrn.NewBitmap(0, lbrnCenter.Scale(XY{k, -k}), im),
			lbrn.NewBitmap(1, lbrnCenter.Scale(XY{k, -k}), im),
			//			lbrn.NewPathWithTabs(2, lbrnCenter, contour.RoundRect(35, 45, 2.5)),
		},
	}

	for _, cut := range pcb.cuts {
		p.Shape = append(p.Shape, lbrn.NewPathWithTabs(2, lbrnCenter, cut))
	}

	for _, hole := range pcb.holes {
		p.Shape = append(p.Shape, lbrn.NewPath(2, lbrnCenter, hole))
	}

	return p.SaveToFile(filename)
}
