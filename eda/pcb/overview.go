package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/shape"
	"temnok/pcbc/util"
)

func (pcb *PCB) SaveOverview() error {
	filename := pcb.SavePath + "overview.png"

	substrateCuts := bitmap.New(pcb.bitmapSize())
	stencilCuts := bitmap.New(pcb.bitmapSize())
	pcb.renderCutsOverview(substrateCuts, stencilCuts)

	image := image.New(
		[]*bitmap.Bitmap{
			pcb.copper,
			pcb.mask,
			pcb.silk,

			substrateCuts,
			stencilCuts,
		},
		[][2]color.Color{
			{color.RGBA{G: 0x40, B: 0x10, A: 0xFF}, color.RGBA{R: 0xC0, G: 0x60, A: 0xFF}},
			{color.RGBA{}, color.RGBA{R: 0x80, G: 0x80, B: 0xFF, A: 0xC0}},
			{color.RGBA{}, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xA0}},

			{color.RGBA{}, color.RGBA{G: 0xFF, B: 0xFF, A: 0xFF}},
			{color.RGBA{}, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}},
		},
		true,
	)
	if err := util.SavePNG(filename, image); err != nil {
		return err
	}

	return nil
}

func (pcb *PCB) renderCutsOverview(substrateCuts, stencilCuts *bitmap.Bitmap) {
	brush := shape.Circle(int(pcb.OverviewCutWidth * pcb.PixelsPerMM))

	pcb.component.Visit(func(c *eda.Component) {
		t := c.Transform.Multiply(pcb.bitmapTransform())

		// Holes
		brush.IterateContours(t, c.Holes, substrateCuts.Set1)

		// Cuts
		brush.IterateContours(t, c.Cuts, substrateCuts.Set1)

		// Pads
		brush.IterateContours(t, c.Pads, stencilCuts.Set1)

		// Perforations
		brush.IterateContours(t, c.Perforations, substrateCuts.Set1)
		brush.IterateContours(t, c.Perforations, stencilCuts.Set1)
	})
}
