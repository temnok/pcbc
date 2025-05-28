// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/shape"
)

var stencilCutSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     lbrn.Param{Value: "Stencil 1"},
		Index:    lbrn.Param{Value: "0"},
		Priority: lbrn.Param{Value: "0"},

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
		Name:     lbrn.Param{Value: "Stencil 2"},
		Index:    lbrn.Param{Value: "1"},
		Priority: lbrn.Param{Value: "1"},

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

func SaveStencil(config *config.Config, component *eda.Component) (*bitmap.Bitmap, error) {
	stencil := bitmap.New(config.BitmapSizeInPixels())

	brush := shape.Circle(int(config.MaskCutWidth * config.PixelsPerMM))
	component.Visit(func(c *eda.Component) {
		t := c.Transform.Multiply(config.BitmapTransform())

		brush.ForEachPathsPixel(c.Pads, t, stencil.Set1)

		if c.OuterCut {
			brush.ForEachPathsPixel(c.Cuts, t, stencil.Set1)
		}
	})

	stencilImage := image.NewSingle(stencil, color.White, color.Black)
	stencilBitmap := lbrn.NewBase64Bitmap(stencilImage)

	p := lbrn.LightBurnProject{
		CutSettingImg: stencilCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShape(0, config.LbrnBitmapScale(), stencilBitmap),
			lbrn.NewBitmapShape(1, config.LbrnBitmapScale(), stencilBitmap),
		},
	}

	return stencil, p.SaveToFile(config.SavePath + "stencil.lbrn")
}
