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
		Name:     &lbrn.Param{Value: "Stencil"},
		Index:    &lbrn.Param{Value: "0"},
		Priority: &lbrn.Param{Value: "0"},

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

func SaveStencil(config *config.Config, component *eda.Component) (*bitmap.Bitmap, error) {
	stencil := bitmap.New(config.BitmapSizeInPixels())

	renderStencil(config, component, stencil)

	stencilImage := image.NewSingle(stencil, color.Transparent, color.Black)

	p := lbrn.LightBurnProject{
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSettingImg: stencilCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShapeFromImage(0, config.LbrnBitmapScale(), stencilImage),
		},
	}

	return stencil, p.SaveToFile(config.SavePath + "1-stencil.lbrn")
}

func renderStencil(config *config.Config, component *eda.Component, stencil *bitmap.Bitmap) {
	bmT := config.BitmapTransform()

	component.Visit(func(c *eda.Component) {
		brush := shape.Circle(int(c.CutsWidth * config.PixelsPerMM))

		t := c.Transform.Multiply(bmT)
		brush.ForEachPathsPixel(c.Pads, t, stencil.Set1)

		t = c.Transform.Multiply(bmT)
		brush.ForEachPathsPixel(c.AlignCuts, t, stencil.Set1)
	})
}
