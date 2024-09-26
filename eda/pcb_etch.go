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
)

func (pcb *PCB) SaveEtch(filename string) error {
	im := pcb.copper.ToImage(color.White, color.Black)

	bitmapTransform := lbrnCenter.ScaleK(1 / pcb.resolution)

	p := lbrn.LightBurnProject{
		CutSettingImg: []lbrn.CutSetting{
			// Etch
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

				Negative: Param{"1"},
			},
			// Clean Pass 1
			{
				Type:     "Image",
				Index:    Param{"1"},
				Name:     Param{"C01"},
				Priority: Param{"1"},

				MaxPower:    Param{"50"},
				QPulseWidth: Param{"2"},
				Frequency:   Param{"280000"},

				Speed:            Param{"2000"},
				Interval:         Param{"0.01"},
				DPI:              Param{"2540"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.05"},
			},
			// Clean Pass 2
			{
				Type:     "Image",
				Index:    Param{"3"},
				Name:     Param{"C03"},
				Priority: Param{"3"},

				MaxPower:    Param{"50"},
				QPulseWidth: Param{"2"},
				Frequency:   Param{"280000"},

				Speed:            Param{"2000"},
				Interval:         Param{"0.01"},
				DPI:              Param{"2540"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.05"},

				Angle: Param{"90"},
			},
		},
		CutSetting: []lbrn.CutSetting{
			{
				Type:     "Cut",
				Index:    Param{"2"},
				Name:     Param{"C02"},
				Priority: Param{"2"},

				Speed:        Param{"100"},
				GlobalRepeat: Param{"80"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				TabsEnabled: Param{"1"},
				TabSize:     Param{"0.2"},
			},
		},
		Shape: []lbrn.Shape{
			lbrn.NewBitmap(0, bitmapTransform, im),
			lbrn.NewBitmap(1, bitmapTransform, im),
			lbrn.NewBitmap(3, bitmapTransform, im),
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
