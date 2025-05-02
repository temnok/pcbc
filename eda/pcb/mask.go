// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/font"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/shape"
)

var maskCutSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     param{Value: "Silk"},
		Index:    param{Value: "0"},
		Priority: param{Value: "0"},

		MaxPower:    param{Value: "5"},
		QPulseWidth: param{Value: "200"},
		Frequency:   param{Value: "20000"},

		NumPasses: param{Value: "1"},
		Speed:     param{Value: "800"},
		Interval:  param{Value: "0.02"},
		DPI:       param{Value: "1270"},

		// Making positive default -- negative is much slower!
		//Negative: param{Value: "1"},

		CrossHatch: param{Value: "1"},
		Angle:      param{Value: "-90"},

		UseDotCorrection: param{Value: "1"},
		DotWidth:         param{Value: "0.05"},
	},
	{
		Type:     "Image",
		Name:     param{Value: "Mask 1"},
		Index:    param{Value: "1"},
		Priority: param{Value: "1"},

		MaxPower:    param{Value: "10"},
		QPulseWidth: param{Value: "80"},
		Frequency:   param{Value: "2000000"},

		NumPasses:        param{Value: "5"},
		Speed:            param{Value: "500"},
		Interval:         param{Value: "0.01"},
		DPI:              param{Value: "2540"},
		UseDotCorrection: param{Value: "1"},
		DotWidth:         param{Value: "0.05"},

		CrossHatch: param{Value: "1"},
		Angle:      param{Value: "90"},
	},
	{
		Type:     "Image",
		Name:     param{Value: "Mask 2"},
		Index:    param{Value: "2"},
		Priority: param{Value: "2"},

		MaxPower:    param{Value: "20"},
		QPulseWidth: param{Value: "80"},
		Frequency:   param{Value: "2000000"},

		NumPasses:        param{Value: "5"},
		Speed:            param{Value: "500"},
		Interval:         param{Value: "0.01"},
		DPI:              param{Value: "2540"},
		UseDotCorrection: param{Value: "1"},
		DotWidth:         param{Value: "0.05"},

		CrossHatch: param{Value: "1"},
		Angle:      param{Value: "90"},
	},
}

func (pcb *PCB) SaveMask() error {
	pcb.component.Visit(func(component *eda.Component) {
		pcb.addSilk(component)
		pcb.cutMask1(component)
	})

	pcb.component.Visit(func(component *eda.Component) {
		pcb.cutMask2(component)
	})

	filename := pcb.SavePath + "mask.lbrn"
	silk := image.NewSingle(pcb.silk, color.White, color.Black)
	mask := image.NewSingle(pcb.mask, color.Transparent, color.Black)
	maskBM := lbrn.NewBase64Bitmap(mask)

	p := &lbrn.LightBurnProject{
		CutSettingImg: maskCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShapeFromImage(0, pcb.lbrnBitmapScale(), silk),
			lbrn.NewBitmapShape(1, pcb.lbrnBitmapScale(), maskBM),
			lbrn.NewBitmapShape(2, pcb.lbrnBitmapScale(), maskBM),
		},
	}

	pcb.addMaskPerforations(p)

	return p.SaveToFile(filename)
}

func (pcb *PCB) addMaskPerforations(p *lbrn.LightBurnProject) {
	pcb.component.Visit(func(component *eda.Component) {
		t := component.Transform.Multiply(pcb.lbrnCenterMove())

		for _, hole := range component.Perforations {
			p.Shape = append(p.Shape, lbrn.NewPath(3, t, hole))
		}
	})

	p.CutSetting = []*lbrn.CutSetting{
		{
			Type:     "Cut",
			Name:     param{Value: "Perforation"},
			Index:    param{Value: "3"},
			Priority: param{Value: "3"},

			Speed:        param{Value: "100"},
			GlobalRepeat: param{Value: "30"},

			MaxPower:    param{Value: "90"},
			QPulseWidth: param{Value: "200"},
			Frequency:   param{Value: "20000"},
		},
	}
}

func (pcb *PCB) addSilk(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Marks:
	brushW := font.Bold * font.WeightScale(t)
	brush := shape.Circle(int(brushW))
	brush.ForEachPathsPixel(c.Marks, t, pcb.silk.Set1)
}

func (pcb *PCB) cutMask1(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	brush := shape.Circle(int(pcb.MaskCutWidth * pcb.PixelsPerMM))

	// Pads
	brush.ForEachPathsPixel(c.Pads, t, pcb.mask.Set1)

	// Cuts
	c.Cuts.ForEachPixelDist(t, int(2*pcb.MaskCutWidth*pcb.PixelsPerMM), func(x, y int) {
		brush.ForEachRowWithOffset(x, y, pcb.mask.Set1)
	})

	// Holes
	brush.ForEachPathsPixel(c.Holes, t, pcb.mask.Set1)

	// Perforations
	brush.ForEachPathsPixel(c.Perforations, t, pcb.mask.Set1)
}

func (pcb *PCB) cutMask2(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	brush := shape.Circle(int(pcb.MaskCutWidth * pcb.PixelsPerMM))

	// Openings
	shape.ForEachRow(c.Openings, t, pcb.mask.Set0)
	brush.ForEachPathsPixel(c.Openings, t, pcb.mask.Set1)
}
