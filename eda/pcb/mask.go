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
		Name:     lbrn.Param{Value: "Silk"},
		Index:    lbrn.Param{Value: "0"},
		Priority: lbrn.Param{Value: "0"},

		MaxPower:    lbrn.Param{Value: "5"},
		QPulseWidth: lbrn.Param{Value: "200"},
		Frequency:   lbrn.Param{Value: "20000"},

		NumPasses: lbrn.Param{Value: "1"},
		Speed:     lbrn.Param{Value: "800"},
		Interval:  lbrn.Param{Value: "0.02"},
		DPI:       lbrn.Param{Value: "1270"},

		// Making positive default -- negative is much slower!
		//Negative: lbrn.Param{Value: "1"},

		CrossHatch: lbrn.Param{Value: "1"},
		Angle:      lbrn.Param{Value: "-90"},

		UseDotCorrection: lbrn.Param{Value: "1"},
		DotWidth:         lbrn.Param{Value: "0.05"},
	},
	{
		Type:     "Image",
		Name:     lbrn.Param{Value: "Mask 1"},
		Index:    lbrn.Param{Value: "1"},
		Priority: lbrn.Param{Value: "1"},

		MaxPower:    lbrn.Param{Value: "10"},
		QPulseWidth: lbrn.Param{Value: "80"},
		Frequency:   lbrn.Param{Value: "2000000"},

		NumPasses:        lbrn.Param{Value: "5"},
		Speed:            lbrn.Param{Value: "500"},
		Interval:         lbrn.Param{Value: "0.01"},
		DPI:              lbrn.Param{Value: "2540"},
		UseDotCorrection: lbrn.Param{Value: "1"},
		DotWidth:         lbrn.Param{Value: "0.05"},

		CrossHatch: lbrn.Param{Value: "1"},
		Angle:      lbrn.Param{Value: "90"},
	},
	{
		Type:     "Image",
		Name:     lbrn.Param{Value: "Mask 2"},
		Index:    lbrn.Param{Value: "2"},
		Priority: lbrn.Param{Value: "2"},

		MaxPower:    lbrn.Param{Value: "20"},
		QPulseWidth: lbrn.Param{Value: "80"},
		Frequency:   lbrn.Param{Value: "2000000"},

		NumPasses:        lbrn.Param{Value: "5"},
		Speed:            lbrn.Param{Value: "500"},
		Interval:         lbrn.Param{Value: "0.01"},
		DPI:              lbrn.Param{Value: "2540"},
		UseDotCorrection: lbrn.Param{Value: "1"},
		DotWidth:         lbrn.Param{Value: "0.05"},

		CrossHatch: lbrn.Param{Value: "1"},
		Angle:      lbrn.Param{Value: "90"},
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
			Name:     lbrn.Param{Value: "Perforation"},
			Index:    lbrn.Param{Value: "3"},
			Priority: lbrn.Param{Value: "3"},

			Speed:        lbrn.Param{Value: "100"},
			GlobalRepeat: lbrn.Param{Value: "30"},

			MaxPower:    lbrn.Param{Value: "90"},
			QPulseWidth: lbrn.Param{Value: "200"},
			Frequency:   lbrn.Param{Value: "20000"},
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
