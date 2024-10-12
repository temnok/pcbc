package eda

import (
	"image/color"
	"temnok/pcbc/geom"
	"temnok/pcbc/lbrn"
)

type Param = lbrn.Param

var (
	lbrnCenter = geom.MoveXY(55, 55)
)

func (pcb *PCB) SaveEtch(filename string) error {
	im := pcb.copper.ToImage(color.White, color.Black)

	bitmapTransform := lbrnCenter.ScaleK(1 / pcb.resolution)

	p := lbrn.LightBurnProject{
		CutSettingImg: []lbrn.CutSetting{
			{
				Type:     "Image",
				Name:     Param{"Etch"},
				Index:    Param{"0"},
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
			{
				Type:     "Image",
				Name:     Param{"Clean 1"},
				Index:    Param{"1"},
				Priority: Param{"1"},

				MaxPower:    Param{"50"},
				QPulseWidth: Param{"2"},
				Frequency:   Param{"280000"},

				Speed:            Param{"2000"},
				Interval:         Param{"0.01"},
				DPI:              Param{"2540"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.15"},

				Negative: Param{"1"},
			},
			{
				Type:     "Image",
				Name:     Param{"Clean 2"},
				Index:    Param{"3"},
				Priority: Param{"3"},

				MaxPower:    Param{"50"},
				QPulseWidth: Param{"2"},
				Frequency:   Param{"280000"},

				Speed:            Param{"2000"},
				Interval:         Param{"0.01"},
				DPI:              Param{"2540"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.15"},

				Angle:    Param{"90"},
				Negative: Param{"1"},
			},
		},
		CutSetting: []lbrn.CutSetting{
			{
				Type:     "Cut",
				Name:     Param{"FR4 Cut"},
				Index:    Param{"2"},
				Priority: Param{"2"},

				Speed:        Param{"100"},
				GlobalRepeat: Param{"75"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				TabsEnabled: Param{"1"},
				TabSize:     Param{"0.2"},
			},
		},
		Shape: []*lbrn.Shape{
			lbrn.NewBitmap(0, bitmapTransform, im),
			lbrn.NewBitmap(1, bitmapTransform, im),
			lbrn.NewBitmap(3, bitmapTransform, im),
		},
	}

	for _, cut := range pcb.component.Cuts {
		p.Shape = append(p.Shape, lbrn.NewPath(2, lbrnCenter, cut)) // Experiment: no tabs
	}

	for _, hole := range pcb.component.Holes {
		p.Shape = append(p.Shape, lbrn.NewPath(2, lbrnCenter, hole))
	}

	return p.SaveToFile(filename)
}
