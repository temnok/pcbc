package eda

import (
	"image/color"
	"temnok/lab/bitmap"
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/path"
	"temnok/lab/shape"
	"temnok/lab/util"
)

type (
	XY    = geom.XY
	Path  = path.Path
	Paths = path.Paths
)

type PCB struct {
	width, height, resolution float64
	trackWidth                float64

	cuts, holes, openings, apertures Paths

	fr4, copper, mask, silk, stencil *bitmap.Bitmap
}

func NewPCB(w, h float64) *PCB {
	const scale = 100

	wi, hi := int(w*scale), int(h*scale)

	return &PCB{
		width:      w,
		height:     h,
		resolution: scale,
		trackWidth: 0.2,

		fr4:     bitmap.NewBitmap(wi, hi),
		copper:  bitmap.NewBitmap(wi, hi),
		mask:    bitmap.NewBitmap(wi, hi),
		silk:    bitmap.NewBitmap(wi, hi),
		stencil: bitmap.NewBitmap(wi, hi),
	}
}

func (pcb *PCB) bitmapTransform() geom.Transform {
	return geom.ScaleK(pcb.resolution).MoveXY(pcb.width/2, pcb.height/2)
}

func (pcb *PCB) Component(c *lib.Component) {
	pcb.component(c, geom.Identity())
}

func (pcb *PCB) component(c *lib.Component, t geom.Transform) {
	if !c.Transform.IsZero() {
		t = t.Multiply(c.Transform)
	}
	bt := pcb.bitmapTransform().Multiply(t)

	brush1 := shape.Circle(int(0.1 * pcb.resolution))
	brush2 := shape.Circle(int(0.2 * pcb.resolution))
	brush02 := shape.Circle(int(0.02 * pcb.resolution))

	// Cuts
	pcb.cuts = append(pcb.cuts, c.Cuts.Transform(t)...)
	c.Cuts.Transform(bt).Jump(int(0.2*pcb.resolution), func(x, y int) {
		brush1.IterateRowsXY(x, y, pcb.mask.Set1)
	})
	brush02.IterateContours(c.Cuts.Transform(bt), pcb.fr4.Set1)

	// Pads
	shape.IterateContoursRows(c.Pads.Transform(bt), pcb.copper.Set1)
	pcb.apertures = append(pcb.apertures, c.Pads.Transform(t)...)
	brush02.IterateContours(c.Pads.Transform(bt), pcb.stencil.Set1)

	// Tracks
	for brushW, tracks := range c.Tracks {
		if brushW == 0 {
			brushW = pcb.trackWidth
		}
		brush := shape.Circle(int(brushW * pcb.resolution))
		brush.IterateContours(tracks.Transform(bt), pcb.copper.Set1)
	}

	// Openings
	brush1.IterateContours(c.Openings.Transform(bt), pcb.mask.Set1)

	// Marks
	for brushW, marks := range c.Marks {
		brush := shape.Circle(int(brushW * pcb.resolution))
		brush.IterateContours(marks.Transform(bt), pcb.silk.Set1)
	}

	// Holes
	pcb.holes = append(pcb.holes, c.Holes.Transform(t)...)

	holes := c.Holes.Transform(bt)
	shape.IterateContoursRows(holes, pcb.copper.Set0)
	brush2.IterateContours(holes, pcb.copper.Set0)
	brush02.IterateContours(holes, pcb.fr4.Set1)

	// Sub-components
	for _, sub := range c.Components {
		pcb.component(sub, t)
	}
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

	image := bitmap.NewBitmapsImage(
		[]*bitmap.Bitmap{pcb.fr4, pcb.copper, pcb.mask, pcb.silk, pcb.stencil},
		[][2]color.Color{
			{color.RGBA{0, 0x40, 0x10, 0xFF}, color.RGBA{0, 0xFF, 0, 0xFF}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0x50, 0, 0xFF}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0x80, 0x80, 0xFF, 0xA0}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0x80}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}},
		},
		true,
	)
	if err := util.SavePng(path+"overview.png", image); err != nil {
		return err
	}

	return nil
}
