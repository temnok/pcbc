// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/shape"
	"temnok/pcbc/util"
)

func SaveOverview(config *config.Config, component *eda.Component, copper, mask, silk *bitmap.Bitmap) error {
	filename := config.SavePath + "overview.png"

	substrateCuts := bitmap.New(config.BitmapSizeInPixels())
	stencilCuts := bitmap.New(config.BitmapSizeInPixels())
	renderCutsOverview(config, component, substrateCuts, stencilCuts)

	im := image.New(
		[]*bitmap.Bitmap{
			copper,
			mask,
			silk,

			substrateCuts,
			stencilCuts,
		},
		[][2]color.Color{
			{color.RGBA{R: 0xC0, G: 0x60, A: 0xFF}, color.RGBA{G: 0x40, B: 0x10, A: 0xFF}},
			{color.RGBA{}, color.RGBA{R: 0x80, G: 0x80, B: 0xFF, A: 0xC0}},
			{color.RGBA{}, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xA0}},

			{color.RGBA{}, color.RGBA{G: 0xFF, B: 0xFF, A: 0xFF}},
			{color.RGBA{}, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}},
		},
	)
	if err := util.SavePNG(filename, im); err != nil {
		return err
	}

	return nil
}

func renderCutsOverview(config *config.Config, component *eda.Component, substrateCuts, stencilCuts *bitmap.Bitmap) {
	brush := shape.Circle(int(config.OverviewCutWidth * config.PixelsPerMM))

	component.Visit(func(c *eda.Component) {
		t := c.Transform.Multiply(config.BitmapTransform())

		// Holes
		brush.ForEachPathsPixel(c.Holes, t, substrateCuts.Set1)

		// Cuts
		brush.ForEachPathsPixel(c.Cuts, t, substrateCuts.Set1)

		// Pads
		brush.ForEachPathsPixel(c.Pads, t, stencilCuts.Set1)

		// Perforations
		brush.ForEachPathsPixel(c.Perforations, t, substrateCuts.Set1)
		brush.ForEachPathsPixel(c.Perforations, t, stencilCuts.Set1)
	})
}
