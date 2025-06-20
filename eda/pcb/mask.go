// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/font"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/shape"
)

var maskCutSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     &lbrn.Param{Value: "Silk"},
		Index:    &lbrn.Param{Value: "0"},
		Priority: &lbrn.Param{Value: "0"},

		MaxPower:    &lbrn.Param{Value: "5"},
		QPulseWidth: &lbrn.Param{Value: "200"},
		Frequency:   &lbrn.Param{Value: "20000"},

		NumPasses: &lbrn.Param{Value: "1"},
		Speed:     &lbrn.Param{Value: "800"},
		Interval:  &lbrn.Param{Value: "0.02"},
		DPI:       &lbrn.Param{Value: "1270"},

		// Making positive default -- negative is much slower!
		//Negative: &lbrn.Param{Value: "1"},

		CrossHatch: &lbrn.Param{Value: "1"},
		Angle:      &lbrn.Param{Value: "-90"},

		UseDotCorrection: &lbrn.Param{Value: "1"},
		DotWidth:         &lbrn.Param{Value: "0.05"},
	},
	{
		Type:     "Image",
		Name:     &lbrn.Param{Value: "Mask"},
		Index:    &lbrn.Param{Value: "1"},
		Priority: &lbrn.Param{Value: "1"},

		MaxPower:    &lbrn.Param{Value: "10"},
		QPulseWidth: &lbrn.Param{Value: "80"},
		Frequency:   &lbrn.Param{Value: "2000000"},

		NumPasses:        &lbrn.Param{Value: "8"},
		Speed:            &lbrn.Param{Value: "500"},
		Interval:         &lbrn.Param{Value: "0.01"},
		DPI:              &lbrn.Param{Value: "2540"},
		UseDotCorrection: &lbrn.Param{Value: "1"},
		DotWidth:         &lbrn.Param{Value: "0.05"},

		CrossHatch: &lbrn.Param{Value: "1"},
		Angle:      &lbrn.Param{Value: "90"},
	},
}

func SaveMask(config *config.Config, component *eda.Component) (*bitmap.Bitmap, *bitmap.Bitmap, error) {
	mask := bitmap.New(config.BitmapSizeInPixels())
	silk := bitmap.New(config.BitmapSizeInPixels())

	component.Visit(func(c *eda.Component) {
		cutMask1(config, c, mask)
		addSilk(config, c, silk)
	})

	filename := config.SavePath + "mask.lbrn"
	silkImage := image.NewSingle(silk, color.White, color.Black)
	maskImage := image.NewSingle(mask, color.Transparent, color.Black)

	p := &lbrn.LightBurnProject{
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSettingImg: maskCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShapeFromImage(0, config.LbrnBitmapScale(), silkImage),
			lbrn.NewBitmapShapeFromImage(1, config.LbrnBitmapScale(), maskImage),
		},
	}

	return mask, silk, p.SaveToFile(filename)
}

func addSilk(config *config.Config, c *eda.Component, silk *bitmap.Bitmap) {
	t := c.Transform.Multiply(config.BitmapTransform())

	// Marks:
	brushW := font.Bold * font.WeightScale(t)
	brush := shape.Circle(int(brushW))
	brush.ForEachPathsPixel(c.Marks, t, silk.Set1)
}

func cutMask1(config *config.Config, c *eda.Component, mask *bitmap.Bitmap) {
	if c.NoOpening {
		return
	}

	t := c.Transform.Multiply(config.BitmapTransform())

	brush := shape.Circle(int(config.MaskCutWidth * config.PixelsPerMM))

	// Pads
	brush.ForEachPathsPixel(c.Pads, t, mask.Set1)

	// Cuts
	if c.OuterCut {
		brush.ForEachPathsPixel(c.Cuts, t, mask.Set1)
	} else {
		c.Cuts.RasterizeIntermittently(t, config.MaskPerforationStep*config.PixelsPerMM, func(x, y int) {
			brush.ForEachRowWithOffset(x, y, mask.Set1)
		})
	}
}
