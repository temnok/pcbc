// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda"
	"temnok/pcbc/font"
	"temnok/pcbc/shape"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
)

type PCB struct {
	component *eda.Component

	Width, Height float64
	PixelsPerMM   float64

	DefaultTrackWidth float64
	ExtraCopperWidth  float64
	CopperClearWidth  float64
	MaskCutWidth      float64
	OverviewCutWidth  float64

	LbrnCenterX, LbrnCenterY float64

	SavePath string

	copper, mask, maskBottom, silk              *bitmap.Bitmap
	overviewCopperbaseCuts, overviewStencilCuts *bitmap.Bitmap
}

func New(component *eda.Component) *PCB {
	width, height := component.Size()
	width, height = width+1, height+1

	return &PCB{
		component: component,

		Width:       width,
		Height:      height,
		PixelsPerMM: 100,

		DefaultTrackWidth: 0.25,
		ExtraCopperWidth:  0.05,
		CopperClearWidth:  0.25,
		MaskCutWidth:      0.1,
		OverviewCutWidth:  0.02,

		LbrnCenterX: 55,
		LbrnCenterY: 55,

		SavePath: "out/",
	}
}

func Generate(component *eda.Component) error {
	return Process(component).SaveFiles()
}

func Process(component *eda.Component) *PCB {
	board := New(component)
	board.Process()
	return board
}

func (pcb *PCB) Process() {
	wi, hi := int(pcb.Width*pcb.PixelsPerMM), int(pcb.Height*pcb.PixelsPerMM)

	pcb.copper = bitmap.NewBitmap(wi, hi)
	pcb.mask = bitmap.NewBitmap(wi, hi)
	pcb.maskBottom = bitmap.NewBitmap(wi, hi)
	pcb.silk = bitmap.NewBitmap(wi, hi)
	pcb.overviewCopperbaseCuts = bitmap.NewBitmap(wi, hi)
	pcb.overviewStencilCuts = bitmap.NewBitmap(wi, hi)

	pcb.copper.Invert()

	pcb.processPass1()
	pcb.processPass2()
}

func (pcb *PCB) processPass1() {
	pcb.component.Visit(func(component *eda.Component) {
		pcb.removeCopper(component)
		pcb.cutCopperbaseOverview(component)
		pcb.addSilk(component)
		pcb.cutMask1(component)
		pcb.cutStencil(component)
	})
}

func (pcb *PCB) processPass2() {
	pcb.component.Visit(func(component *eda.Component) {
		pcb.addCopper(component)
		pcb.cutMask2(component)
	})
}

func (pcb *PCB) removeCopper(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Clears
	shape.IterateContoursRows(t, c.Clears, pcb.copper.Set0)

	clearWidth := 2 * (pcb.CopperClearWidth - pcb.ExtraCopperWidth)

	// Pads
	clearBrush := shape.Circle(int(clearWidth * pcb.PixelsPerMM))
	clearBrush.IterateContours(t, c.Pads, pcb.copper.Set0)

	// Non-ground tracks
	brushW := c.TrackWidth
	if brushW == 0 {
		brushW = pcb.DefaultTrackWidth
	}
	brush := shape.Circle(int((brushW + clearWidth) * pcb.PixelsPerMM))
	brush.IterateContours(t, c.Tracks, pcb.copper.Set0)

	// TODO: remove the following line
	clearBrush = shape.Circle(int(pcb.CopperClearWidth * pcb.PixelsPerMM))

	// Cuts
	clearBrush.IterateContours(t, c.Cuts, pcb.copper.Set0)

	// Holes
	clearBrush.IterateContours(t, c.Holes, pcb.copper.Set0)

	// Perforations
	clearBrush.IterateContours(t, c.Perforations, pcb.copper.Set0)
}

func (pcb *PCB) addCopper(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Pads
	shape.IterateContoursRows(t, c.Pads, pcb.copper.Set1)

	extraCopperBrush := shape.Circle(int(pcb.ExtraCopperWidth * pcb.PixelsPerMM))
	extraCopperBrush.IterateContours(t, c.Pads, pcb.copper.Set1)

	// Tracks
	brushW := c.TrackWidth
	if brushW == 0 {
		brushW = pcb.DefaultTrackWidth
	}

	brush := shape.Circle(int((brushW + pcb.ExtraCopperWidth) * pcb.PixelsPerMM))
	brush.IterateContours(t, c.Tracks, pcb.copper.Set1)
	brush.IterateContours(t, c.GroundTracks, pcb.copper.Set1)
}

func (pcb *PCB) cutCopperbaseOverview(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	brush := shape.Circle(int(pcb.OverviewCutWidth * pcb.PixelsPerMM))

	// Holes
	brush.IterateContours(t, c.Holes, pcb.overviewCopperbaseCuts.Set1)

	// Cuts
	brush.IterateContours(t, c.Cuts, pcb.overviewCopperbaseCuts.Set1)

	// Perforations
	brush.IterateContours(t, c.Perforations, pcb.overviewCopperbaseCuts.Set1)
}

func (pcb *PCB) addSilk(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Marks:
	brushW := font.Bold * font.WeightScale(t)
	brush := shape.Circle(int(brushW))
	brush.IterateContours(t, c.Marks, pcb.silk.Set1)
}

func (pcb *PCB) cutMask1(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	brush := shape.Circle(int(pcb.MaskCutWidth * pcb.PixelsPerMM))

	// Pads
	brush.IterateContours(t, c.Pads, pcb.mask.Set1)

	// Cuts
	c.Cuts.Jump(t, int(2*pcb.MaskCutWidth*pcb.PixelsPerMM), func(x, y int) {
		brush.IterateRowsXY(x, y, pcb.mask.Set1)
		brush.IterateRowsXY(x, y, pcb.maskBottom.Set1)
	})

	// Holes
	brush.IterateContours(t, c.Holes, pcb.mask.Set1)
	brush.IterateContours(t, c.Holes, pcb.maskBottom.Set1)

	// Perforations
	brush.IterateContours(t, c.Perforations, pcb.mask.Set1)
	brush.IterateContours(t, c.Perforations, pcb.maskBottom.Set1)
}

func (pcb *PCB) cutMask2(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	brush := shape.Circle(int(pcb.MaskCutWidth * pcb.PixelsPerMM))

	// Openings
	shape.IterateContoursRows(t, c.Openings, pcb.mask.Set0)
	brush.IterateContours(t, c.Openings, pcb.mask.Set1)
}

func (pcb *PCB) cutStencil(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	brush := shape.Circle(int(pcb.OverviewCutWidth * pcb.PixelsPerMM))

	// Pads
	brush.IterateContours(t, c.Pads, pcb.overviewStencilCuts.Set1)

	// Perforations
	brush.IterateContours(t, c.Perforations, pcb.overviewStencilCuts.Set1)
}

func (pcb *PCB) bitmapTransform() transform.T {
	return transform.Move(pcb.Width/2, pcb.Height/2).Scale(pcb.PixelsPerMM, pcb.PixelsPerMM)
}

func (pcb *PCB) SaveFiles() error {
	return util.RunConcurrently(
		pcb.SaveEtch,
		pcb.SaveMask,
		pcb.SaveMaskBottom,
		pcb.SaveStencil,
		pcb.SaveOverview,
	)
}

func (pcb *PCB) SaveOverview() error {
	filename := pcb.SavePath + "overview.png"

	image := bitmap.NewBitmapsImage(
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
