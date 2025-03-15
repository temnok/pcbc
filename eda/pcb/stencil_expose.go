// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/shape"
)

func (pcb *PCB) createStencilExposeBitmap() *bitmap.Bitmap {
	bm := bitmap.New(pcb.bitmapSize())

	brush := shape.Circle(int(pcb.StencilExposeWidth * pcb.PixelsPerMM))

	bitmapTransform := pcb.bitmapTransform()

	pcb.component.Visit(func(c *eda.Component) {
		t := c.Transform.Multiply(bitmapTransform)

		shape.IterateContoursRows(t, c.Pads, bm.Set1)
		brush.IterateContours(t, c.Pads, bm.Set1)

		shape.IterateContoursRows(t, c.Perforations, bm.Set1)
	})

	return bm
}

func (pcb *PCB) SaveStencilExpose() error {
	filename := pcb.SavePath + "stencil-expose.lbrn"
	im := image.NewSingle(pcb.createStencilExposeBitmap(), color.White, color.Black)

	p := lbrn.LightBurnProject{
		CutSettingImg: []*lbrn.CutSetting{
			{
				Type:     "Image",
				Name:     Param{Value: "Stencil-Pad Clean"},
				Index:    Param{Value: "0"},
				Priority: Param{Value: "0"},

				MaxPower:    Param{Value: "25"},
				QPulseWidth: Param{Value: "1"},
				Frequency:   Param{Value: "650000"},

				NumPasses: Param{Value: "5"},
				Speed:     Param{Value: "500"},
				Interval:  Param{Value: "0.01"},
				DPI:       Param{Value: "2540"},

				CrossHatch: Param{Value: "1"},
				Angle:      Param{Value: "90"},
			},
		},
		Shape: []*lbrn.Shape{
			lbrn.NewBitmap(0, pcb.lbrnBitmapScale(), im),
		},
	}

	return p.SaveToFile(filename)
}
