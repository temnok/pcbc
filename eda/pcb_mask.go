package eda

import (
	"image/color"
	"temnok/lab/lbrn"
)

func (pcb *PCB) SaveMask(filename string) error {
	silk := pcb.silk.ToImage(color.White, color.Black)
	mask := pcb.mask.ToImage(color.Transparent, color.Black)

	k := 1 / pcb.resolution

	p := lbrn.LightBurnProject{
		CutSettingImg: []lbrn.CutSetting{
			{
				Type:     "Image",
				Index:    Param{"0"},
				Name:     Param{"C00"},
				Priority: Param{"0"},

				MaxPower:    Param{"6"},
				QPulseWidth: Param{"80"},
				Frequency:   Param{"2000000"},

				NumPasses:        Param{"8"},
				Speed:            Param{"1000"},
				Interval:         Param{"0.02"},
				DPI:              Param{"1270"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.05"},

				CrossHatch: Param{"1"},
				Angle:      Param{"90"},
			},
			{
				Type:     "Image",
				Index:    Param{"1"},
				Name:     Param{"C01"},
				Priority: Param{"1"},

				MaxPower:    Param{"10"},
				QPulseWidth: Param{"80"},
				Frequency:   Param{"2000000"},

				NumPasses:        Param{"8"},
				Speed:            Param{"500"},
				Interval:         Param{"0.01"},
				DPI:              Param{"2540"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.05"},

				CrossHatch: Param{"1"},
				Angle:      Param{"90"},
			},
			{
				Type:     "Image",
				Index:    Param{"2"},
				Name:     Param{"C02"},
				Priority: Param{"2"},

				MaxPower:    Param{"20"},
				QPulseWidth: Param{"80"},
				Frequency:   Param{"2000000"},

				NumPasses:        Param{"8"},
				Speed:            Param{"500"},
				Interval:         Param{"0.01"},
				DPI:              Param{"2540"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.05"},

				CrossHatch: Param{"1"},
				Angle:      Param{"90"},
			},
		},
		CutSetting: []lbrn.CutSetting{
			{
				Type:     "Cut",
				Index:    Param{"3"},
				Name:     Param{"C03"},
				Priority: Param{"3"},

				Speed:     Param{"200"},
				NumPasses: Param{"15"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},
			},
		},
		Shape: []lbrn.Shape{
			lbrn.NewBitmap(0, lbrnCenter.Scale(XY{k, -k}), silk),
			lbrn.NewBitmap(1, lbrnCenter.Scale(XY{k, -k}), mask),
			lbrn.NewBitmap(2, lbrnCenter.Scale(XY{k, -k}), mask),
		},
	}

	for _, cut := range pcb.maskHoles {
		p.Shape = append(p.Shape, lbrn.NewPath(3, lbrnCenter, cut))
	}

	return p.SaveToFile(filename)
}
