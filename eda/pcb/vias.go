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

func viaStencilSettings() []*lbrn.CutSetting {
	return []*lbrn.CutSetting{
		{
			Type:     "Cut",
			Name:     &lbrn.Param{Value: "Cut Al200um"},
			Index:    &lbrn.Param{Value: "0"},
			Priority: &lbrn.Param{Value: "0"},

			MaxPower:    &lbrn.Param{Value: "90"},
			QPulseWidth: &lbrn.Param{Value: "200"},
			Frequency:   &lbrn.Param{Value: "20000"},

			NumPasses:    &lbrn.Param{Value: "1"},
			GlobalRepeat: &lbrn.Param{Value: "50"},
			Speed:        &lbrn.Param{Value: "800"},
		},
	}
}

var viaAblateSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     &lbrn.Param{Value: "Ablate Vias"},
		Index:    &lbrn.Param{Value: "1"},
		Priority: &lbrn.Param{Value: "1"},

		MaxPower:    &lbrn.Param{Value: "30"},
		QPulseWidth: &lbrn.Param{Value: "30"},
		Frequency:   &lbrn.Param{Value: "3000000"},

		NumPasses: &lbrn.Param{Value: "4"},
		Speed:     &lbrn.Param{Value: "400"},
		Interval:  &lbrn.Param{Value: "0.01"},
		DPI:       &lbrn.Param{Value: "2540"},

		Angle:            &lbrn.Param{Value: "-90"},
		CrossHatch:       &lbrn.Param{Value: "1"},
		UseDotCorrection: &lbrn.Param{Value: "1"},
		DotWidth:         &lbrn.Param{Value: "0.05"},
	},
}

func SaveVias(config *config.Config, component *eda.Component, stencilMode bool) error {
	stencilSettings := viaStencilSettings()

	p := &lbrn.LightBurnProject{
		UIPrefs:    lbrn.UIPrefsDefaults,
		CutSetting: stencilSettings,
	}

	nonEmpty := renderViaStencil(config, component, p)
	if !nonEmpty {
		return nil
	}

	if stencilMode {
		return p.SaveToFile(config.SavePath + "via-stencil.lbrn")
	}

	stencilSettings[0].DoOutput = &lbrn.Param{"0"}

	ablateBitmap := bitmap.New(config.BitmapSizeInPixels())
	renderViaAblate(config, component, ablateBitmap)
	ablateImage := image.NewSingle(ablateBitmap, color.Transparent, color.Black)

	p.CutSettingImg = viaAblateSettings
	p.Shape = append(p.Shape,
		lbrn.NewBitmapShapeFromImage(1, config.LbrnBitmapScale(), ablateImage),
	)

	return p.SaveToFile(config.SavePath + "via-ablate.lbrn")
}

func renderViaStencil(config *config.Config, component *eda.Component, p *lbrn.LightBurnProject) bool {
	hasVias := false

	component.Visit(func(c *eda.Component) {
		hasVias = hasVias || len(c.Vias) > 0

		t := c.Transform.Multiply(config.LbrnCenterMove())

		for _, via := range c.Vias {
			p.Shape = append(p.Shape, lbrn.NewPath(0, t, via))
		}

		if c.CutsOuter {
			for _, cut := range c.Cuts {
				p.Shape = append(p.Shape, lbrn.NewPath(0, t, cut))
			}
		}
	})

	return hasVias
}

func renderViaAblate(config *config.Config, component *eda.Component, ablate *bitmap.Bitmap) {
	bmT := config.BitmapTransform()

	component.Visit(func(c *eda.Component) {
		t := c.Transform.Multiply(bmT)
		shape.ForEachRow(c.Vias, t, ablate.Set1)
	})

	brush := shape.Circle(int(config.ViaAblateWidth * 2 * config.PixelsPerMM))
	component.Visit(func(c *eda.Component) {
		t := c.Transform.Multiply(bmT)
		brush.ForEachPathsPixel(c.Vias, t, ablate.Set1)
	})
}
