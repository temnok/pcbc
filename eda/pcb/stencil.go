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

		NumPasses:        &lbrn.Param{Value: "10"},
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

	bmT := config.BitmapTransform()

	// Pass 1: draw pads
	component.Visit(func(c *eda.Component) {
		if c.OuterCut {
			return
		}

		t := c.Transform.Multiply(bmT)
		shape.ForEachRow(c.Pads, t, stencil.Set1)
	})

	// Pass 2
	if config.StencilPadDist > 0 {
		component.Visit(func(c *eda.Component) {
			if c.OuterCut {
				return
			}

			brushD := 2 * config.StencilPadDist
			brush := shape.Circle(int(brushD * config.PixelsPerMM))

			t := c.Transform.Multiply(bmT)
			brush.ForEachPathsPixel(c.Pads, t, stencil.Set0)
		})
	}

	savedBitmap := stencil.Clone()

	// Pass 3
	component.Visit(func(c *eda.Component) {
		if c.OuterCut {
			return
		}

		clearWidth := 2 * (config.StencilPadDist + config.MaskCutWidth)
		brush := shape.Circle(int(clearWidth * config.PixelsPerMM))

		t := c.Transform.Multiply(bmT)
		brush.ForEachPathsPixel(c.Pads, t, stencil.Set0)
	})

	// Pass 4
	stencil.Xor(savedBitmap)

	// Pass 5
	outerCutBrush := shape.Circle(int(config.MaskCutWidth * config.PixelsPerMM))
	component.Visit(func(c *eda.Component) {
		if !c.OuterCut {
			return
		}

		t := c.Transform.Multiply(bmT)
		outerCutBrush.ForEachPathsPixel(c.Cuts, t, stencil.Set1)
	})

	stencilImage := image.NewSingle(stencil, color.White, color.Black)

	p := lbrn.LightBurnProject{
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSettingImg: stencilCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShapeFromImage(0, config.LbrnBitmapScale(), stencilImage),
		},
	}

	return stencil, p.SaveToFile(config.SavePath + "stencil.lbrn")
}
