package eda

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/path"
	"temnok/pcbc/shape"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
)

type (
	Path  = path.Path
	Paths = path.Paths
)

type PCB struct {
	width, height, resolution float64
	trackWidth                float64
	lbrnCenter                path.Point

	component *lib.Component

	fr4, copper, mask, maskB, silk, stencil *bitmap.Bitmap
}

func NewPCB(width, height float64, component *lib.Component) *PCB {
	const scale = 100

	wi, hi := int(width*scale), int(height*scale)

	pcb := &PCB{
		width:      width,
		height:     height,
		resolution: scale,
		trackWidth: 0.25,

		fr4:     bitmap.NewBitmap(wi, hi),
		copper:  bitmap.NewBitmap(wi, hi),
		mask:    bitmap.NewBitmap(wi, hi),
		maskB:   bitmap.NewBitmap(wi, hi),
		silk:    bitmap.NewBitmap(wi, hi),
		stencil: bitmap.NewBitmap(wi, hi),
	}

	pcb.copper.Invert()

	pcb.setComponent(component)

	pcb.lbrnCenter = path.Point{55, 55}

	return pcb
}

func (pcb *PCB) setComponent(c *lib.Component) {
	c = c.Squash()
	pcb.component = c

	bt := transform.Move(pcb.width/2, pcb.height/2).Scale(pcb.resolution, pcb.resolution)

	brush1 := shape.Circle(int(0.1 * pcb.resolution))
	brush02 := shape.Circle(int(0.02 * pcb.resolution))

	const clearBrushW = 0.5
	clearBrush := shape.Circle(int(clearBrushW * pcb.resolution))

	const extraCopper = 0.05 // compensate copper lost during etching
	//extraCopperBrush := shape.Circle(int(extraCopper * pcb.resolution))

	cutClearBrush := shape.Circle(int((clearBrushW/2 - extraCopper) * pcb.resolution))

	// Clears: remove groundfill
	shape.IterateContoursRows(c.Clears.Apply(bt), pcb.copper.Set0)

	// Pads: remove groundfill
	pads := c.Pads.Apply(bt)
	clearBrush.IterateContours(pads, pcb.copper.Set0)

	// Non-ground tracks: remove groundfill
	for brushW, tracks := range c.Tracks {
		if brushW == 0 {
			brushW = pcb.trackWidth
		}
		brush := shape.Circle(int((brushW + clearBrushW) * pcb.resolution))
		brush.IterateContours(tracks.Apply(bt), pcb.copper.Set0)
	}

	// Pads
	shape.IterateContoursRows(pads, pcb.copper.Set1)
	//extraCopperBrush.IterateContours(pads, pcb.copper.Set1)
	brush1.IterateContours(pads, pcb.mask.Set1)

	//resizedPads := c.Pads.Resize(-StencilShrink).Apply(bt)
	brush02.IterateContours(pads, pcb.stencil.Set1)

	// Tracks
	for brushW, tracks := range (path.Strokes{}).Append(c.Tracks, c.GroundTracks) {
		if brushW == 0 {
			brushW = pcb.trackWidth
		}
		brush := shape.Circle(int((brushW + extraCopper) * pcb.resolution))
		brush.IterateContours(tracks.Apply(bt), pcb.copper.Set1)
	}

	// Marks
	for brushW, marks := range c.Marks {
		brush := shape.Circle(int(brushW * pcb.resolution))
		brush.IterateContours(marks.Apply(bt), pcb.silk.Set1)
	}

	// Holes
	holes := c.Holes.Apply(bt)
	brush1.IterateContours(holes, pcb.mask.Set1)
	brush1.IterateContours(holes, pcb.maskB.Set1)
	shape.IterateContoursRows(holes, pcb.copper.Set1)
	cutClearBrush.IterateContours(holes, pcb.copper.Set0)
	brush02.IterateContours(holes, pcb.fr4.Set1)

	// Cuts
	cuts := c.Cuts.Apply(bt)
	cutClearBrush.IterateContours(cuts, pcb.copper.Set0)
	cuts.Jump(int(0.2*pcb.resolution), func(x, y int) {
		brush1.IterateRowsXY(x, y, pcb.mask.Set1)
		brush1.IterateRowsXY(x, y, pcb.maskB.Set1)
	})
	brush02.IterateContours(cuts, pcb.fr4.Set1)

	// Openings
	openings := c.Openings.Apply(bt)
	shape.IterateContoursRows(openings, pcb.mask.Set0)
	brush1.IterateContours(openings, pcb.mask.Set1)
}

func (pcb *PCB) SaveFiles(path string) error {
	return util.GoAll([]func() error{
		func() error { return pcb.SaveEtch(path + "etch.lbrn") },
		func() error { return pcb.SaveMask(path + "mask.lbrn") },
		func() error { return pcb.SaveMaskBottom(path + "mask-bottom.lbrn") },
		func() error { return pcb.SaveStencil(path + "stencil.lbrn") },
		func() error { return pcb.SaveOverview(path + "overview.png") },
	})
}

func (pcb *PCB) SaveOverview(filename string) error {
	image := bitmap.NewBitmapsImage(
		[]*bitmap.Bitmap{pcb.copper, pcb.fr4, pcb.mask, pcb.silk, pcb.stencil},
		[][2]color.Color{
			{color.RGBA{0, 0x40, 0x10, 0xFF}, color.RGBA{0xC0, 0x60, 0, 0xFF}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0, 0xFF, 0, 0xFF}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0x80, 0x80, 0xFF, 0xA0}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0xA0}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}},
		},
		true,
	)
	if err := util.SavePng(filename, image); err != nil {
		return err
	}

	return nil
}
