// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
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

func SaveMask(config *PCB, component *eda.Component) (*bitmap.Bitmap, error) {
	mask := bitmap.New(config.bitmapSize())

	component.Visit(func(c *eda.Component) {
		addSilk(config, c)
		cutMask1(config, c, mask)
	})

	component.Visit(func(c *eda.Component) {
		cutMask2(config, c, mask)
	})

	filename := config.SavePath + "mask.lbrn"
	silk := image.NewSingle(config.silk, color.White, color.Black)
	maskImage := image.NewSingle(mask, color.Transparent, color.Black)
	maskBM := lbrn.NewBase64Bitmap(maskImage)

	p := &lbrn.LightBurnProject{
		CutSettingImg: maskCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShapeFromImage(0, config.lbrnBitmapScale(), silk),
			lbrn.NewBitmapShape(1, config.lbrnBitmapScale(), maskBM),
			lbrn.NewBitmapShape(2, config.lbrnBitmapScale(), maskBM),
		},
	}

	addMaskPerforations(config, p)

	return mask, p.SaveToFile(filename)
}

func addMaskPerforations(config *PCB, p *lbrn.LightBurnProject) {
	config.component.Visit(func(component *eda.Component) {
		t := component.Transform.Multiply(config.lbrnCenterMove())

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

func addSilk(config *PCB, c *eda.Component) {
	t := c.Transform.Multiply(config.bitmapTransform())

	// Marks:
	brushW := font.Bold * font.WeightScale(t)
	brush := shape.Circle(int(brushW))
	brush.ForEachPathsPixel(c.Marks, t, config.silk.Set1)
}

func cutMask1(config *PCB, c *eda.Component, mask *bitmap.Bitmap) {
	t := c.Transform.Multiply(config.bitmapTransform())

	brush := shape.Circle(int(config.MaskCutWidth * config.PixelsPerMM))

	// Pads
	brush.ForEachPathsPixel(c.Pads, t, mask.Set1)

	// Cuts
	c.Cuts.ForEachPixelDist(t, int(2*config.MaskCutWidth*config.PixelsPerMM), func(x, y int) {
		brush.ForEachRowWithOffset(x, y, mask.Set1)
	})

	// Holes
	brush.ForEachPathsPixel(c.Holes, t, mask.Set1)

	// Perforations
	brush.ForEachPathsPixel(c.Perforations, t, mask.Set1)
}

func cutMask2(config *PCB, c *eda.Component, mask *bitmap.Bitmap) {
	t := c.Transform.Multiply(config.bitmapTransform())

	brush := shape.Circle(int(config.MaskCutWidth * config.PixelsPerMM))

	// Openings
	shape.ForEachRow(c.Openings, t, mask.Set0)
	brush.ForEachPathsPixel(c.Openings, t, mask.Set1)
}
