package eda

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
	"temnok/pcbc/shape"
	"temnok/pcbc/util"
)

type (
	Path  = path.Path
	Paths = path.Paths
)

type PCB struct {
	width, height, resolution float64
	trackWidth                float64

	component *lib.Component

	fr4, copper, mask, silk, stencil *bitmap.Bitmap
}

func NewPCB(w, h float64) *PCB {
	const scale = 100

	wi, hi := int(w*scale), int(h*scale)

	pcb := &PCB{
		width:      w,
		height:     h,
		resolution: scale,
		trackWidth: 0.25,

		fr4:     bitmap.NewBitmap(wi, hi),
		copper:  bitmap.NewBitmap(wi, hi),
		mask:    bitmap.NewBitmap(wi, hi),
		silk:    bitmap.NewBitmap(wi, hi),
		stencil: bitmap.NewBitmap(wi, hi),
	}

	pcb.copper.Invert()

	return pcb
}

func (pcb *PCB) bitmapTransform() geom.Transform {
	return geom.ScaleK(pcb.resolution).MoveXY(pcb.width/2, pcb.height/2)
}

func (pcb *PCB) Component(c *lib.Component) {

	c = c.Squash()
	pcb.component = c

	bt := pcb.bitmapTransform()

	brush1 := shape.Circle(int(0.1 * pcb.resolution))
	brush25 := shape.Circle(int(0.25 * pcb.resolution))
	brush02 := shape.Circle(int(0.02 * pcb.resolution))
	brush5 := shape.Circle(int(0.5 * pcb.resolution))

	// Clears: remove groundfill
	shape.IterateContoursRows(c.Clears.Transform(bt), pcb.copper.Set0)

	// Pads: remove groundfill
	pads := c.Pads.Transform(bt)
	brush5.IterateContours(pads, pcb.copper.Set0)

	// Non-ground tracks: remove groundfill
	for brushW, tracks := range c.Tracks {
		if brushW == 0 {
			brushW = pcb.trackWidth
		}
		brush := shape.Circle(int((brushW + 0.5) * pcb.resolution))
		brush.IterateContours(tracks.Transform(bt), pcb.copper.Set0)
	}

	// Pads
	shape.IterateContoursRows(pads, pcb.copper.Set1)
	brush1.IterateContours(pads, pcb.mask.Set1)
	brush02.IterateContours(pads, pcb.stencil.Set1)

	// Tracks
	for brushW, tracks := range (path.Strokes{}).Append(c.Tracks, c.GroundTracks) {
		if brushW == 0 {
			brushW = pcb.trackWidth
		}
		brush := shape.Circle(int(brushW * pcb.resolution))
		brush.IterateContours(tracks.Transform(bt), pcb.copper.Set1)
	}

	// Marks
	for brushW, marks := range c.Marks {
		brush := shape.Circle(int(brushW * pcb.resolution))
		brush.IterateContours(marks.Transform(bt), pcb.silk.Set1)
	}

	// Holes
	holes := c.Holes.Transform(bt)
	brush1.IterateContours(holes, pcb.mask.Set1)
	shape.IterateContoursRows(holes, pcb.copper.Set1)
	brush25.IterateContours(holes, pcb.copper.Set0)
	brush02.IterateContours(holes, pcb.fr4.Set1)

	// Cuts
	cuts := c.Cuts.Transform(bt)
	brush25.IterateContours(cuts, pcb.copper.Set0)
	cuts.Jump(int(0.2*pcb.resolution), func(x, y int) {
		brush1.IterateRowsXY(x, y, pcb.mask.Set1)
	})
	brush02.IterateContours(cuts, pcb.fr4.Set1)

	// Openings
	openings := c.Openings.Transform(bt)
	shape.IterateContoursRows(openings, pcb.mask.Set0)
	brush1.IterateContours(openings, pcb.mask.Set1)
}

func (pcb *PCB) SaveFiles(path string) error {
	if err := pcb.SaveEtch(path + "etch.lbrn"); err != nil {
		return err
	}

	if err := pcb.SaveMask(path + "mask.lbrn"); err != nil {
		return err
	}

	if err := pcb.SaveStencil(path + "stencil.lbrn"); err != nil {
		return err
	}

	if err := pcb.SaveOverview(path + "overview.png"); err != nil {
		return err
	}

	return nil
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
