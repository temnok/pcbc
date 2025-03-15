package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/util"
)

func (pcb *PCB) SaveOverview() error {
	filename := pcb.SavePath + "overview.png"

	image := image.New(
		[]*bitmap.Bitmap{
			pcb.copper,
			pcb.mask,
			pcb.silk,

			pcb.overviewCopperbaseCuts,
			pcb.overviewStencilCuts,
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
