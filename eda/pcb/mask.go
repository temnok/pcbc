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

func SaveMask(config *config.Config, component *eda.Component, back bool) (*bitmap.Bitmap, *bitmap.Bitmap, error) {
	mask := bitmap.New(config.BitmapSizeInPixels())
	silk := bitmap.New(config.BitmapSizeInPixels())

	nonEmpty := false

	component.Visit(func(c *eda.Component) {
		cutMask(config, c, back, mask)

		addSilk(config, c, back, silk)

		nonEmpty = nonEmpty || len(c.Pads) > 0 || len(c.Marks) > 0
	})

	filename := config.SavePath + fileNamePrefix[back] + "mask.lbrn"
	silkImage := image.NewSingle(silk, color.Transparent, color.Black)
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

func cutMask(config *config.Config, c *eda.Component, back bool, mask *bitmap.Bitmap) {
	t := c.Transform.Multiply(config.BitmapTransform())

	brush := shape.Circle(int(config.MaskCutWidth * config.PixelsPerMM))

	// Pads
	if !back {
		brush.ForEachPathsPixel(c.Pads, t, mask.Set1)
	}

	// Cuts
	if c.CutsOuter && c.CutsInner {
		brush.ForEachPathsPixel(c.Cuts, t, mask.Set1)
	} else if !c.CutsInner {
		c.Cuts.RasterizeIntermittently(t, config.MaskPerforationStep*config.PixelsPerMM, func(x, y int) {
			brush.ForEachRowWithOffset(x, y, mask.Set1)
		})
	}
}

func addSilk(config *config.Config, c *eda.Component, back bool, silk *bitmap.Bitmap) {
	if c.Back != back {
		return
	}

	t := c.Transform.Multiply(config.BitmapTransform())

	// Marks:
	brushW := font.Bold * font.WeightScale(t)
	brush := shape.Circle(int(brushW))
	brush.ForEachPathsPixel(c.Marks, t, silk.Set1)
}
