package eda

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/font"
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
	width, height float64

	trackWidth float64
	lbrnCenter path.Point

	savePath  string
	component *Component

	fr4, copper, mask, maskB, silk, stencil *bitmap.Bitmap
}

const (
	resolution = 100.0 // pixels per mm

	clearBrushDiameter = 0.5

	extraCopper = 0.05 // compensate copper lost during etching
)

var (
	brush1  = shape.Circle(int(0.1 * resolution))
	brush02 = shape.Circle(int(0.02 * resolution))
)

func GeneratePCB(component *Component) error {
	return NewPCB(component).SaveFiles()
}

func NewPCB(component *Component) *PCB {
	width, height := component.Size()
	width, height = width+1, height+1

	wi, hi := int(width*resolution), int(height*resolution)

	pcb := &PCB{
		width:  width,
		height: height,

		component: component,

		trackWidth: 0.25,

		lbrnCenter: path.Point{X: 55, Y: 55},
		savePath:   "out/",

		fr4:     bitmap.NewBitmap(wi, hi),
		copper:  bitmap.NewBitmap(wi, hi),
		mask:    bitmap.NewBitmap(wi, hi),
		maskB:   bitmap.NewBitmap(wi, hi),
		silk:    bitmap.NewBitmap(wi, hi),
		stencil: bitmap.NewBitmap(wi, hi),
	}

	pcb.copper.Invert()

	pcb.processBoard()
	pcb.processMask()
	pcb.processStencil()

	return pcb
}

func (pcb *PCB) processBoard() {
	pcb.component.Visit(pcb.removeCopper)
	pcb.component.Visit(pcb.addCopper)
	pcb.component.Visit(pcb.cutBoard)
}

func (pcb *PCB) processMask() {
	pcb.component.Visit(pcb.addMarks)
	pcb.component.Visit(pcb.cutOpenings)
}

func (pcb *PCB) processStencil() {
	pcb.component.Visit(pcb.cutStencil)
}

func (pcb *PCB) removeCopper(c *Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Clears
	clears := c.Clears.Apply(t)
	shape.IterateContoursRows(clears, pcb.copper.Set0)

	// Pads
	pads := c.Pads.Apply(t)
	clearBrush := shape.Circle(int(clearBrushDiameter * resolution))
	clearBrush.IterateContours(pads, pcb.copper.Set0)

	// Non-ground tracks
	brushW := c.TrackThickness
	if brushW == 0 {
		brushW = pcb.trackWidth
	}
	brush := shape.Circle(int((brushW + clearBrushDiameter) * resolution))
	brush.IterateContours(c.Tracks.Apply(t), pcb.copper.Set0)

	cutClearBrush := shape.Circle(int((clearBrushDiameter / 2) * resolution))

	// Holes
	holes := c.Holes.Apply(t)
	cutClearBrush.IterateContours(holes, pcb.copper.Set0)

	// Cuts
	cuts := c.Cuts.Apply(t)
	cutClearBrush.IterateContours(cuts, pcb.copper.Set0)
}

func (pcb *PCB) addCopper(c *Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Pads
	pads := c.Pads.Apply(t)
	shape.IterateContoursRows(pads, pcb.copper.Set1)

	extraCopperBrush := shape.Circle(int(extraCopper * resolution))
	extraCopperBrush.IterateContours(pads, pcb.copper.Set1)

	// Tracks
	brushW := c.TrackThickness
	if brushW == 0 {
		brushW = pcb.trackWidth
	}
	brush := shape.Circle(int((brushW + extraCopper) * resolution))
	brush.IterateContours(c.Tracks.Apply(t), pcb.copper.Set1)
	brush.IterateContours(c.GroundTracks.Apply(t), pcb.copper.Set1)
}

func (pcb *PCB) cutBoard(c *Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Holes
	holes := c.Holes.Apply(t)
	brush02.IterateContours(holes, pcb.fr4.Set1)

	// Cuts
	cuts := c.Cuts.Apply(t)
	brush02.IterateContours(cuts, pcb.fr4.Set1)
}

func (pcb *PCB) addMarks(c *Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Marks:
	brushW := font.Bold * font.WeightScale(t)
	brush := shape.Circle(int(brushW))
	brush.IterateContours(c.Marks.Apply(t), pcb.silk.Set1)
}

func (pcb *PCB) cutOpenings(c *Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Pads
	pads := c.Pads.Apply(t)
	brush1.IterateContours(pads, pcb.mask.Set1)

	// Holes
	holes := c.Holes.Apply(t)
	brush1.IterateContours(holes, pcb.mask.Set1)
	brush1.IterateContours(holes, pcb.maskB.Set1)

	// Cuts
	cuts := c.Cuts.Apply(t)
	cuts.Jump(int(0.2*resolution), func(x, y int) {
		brush1.IterateRowsXY(x, y, pcb.mask.Set1)
		brush1.IterateRowsXY(x, y, pcb.maskB.Set1)
	})

	// Openings
	openings := c.Openings.Apply(t)
	shape.IterateContoursRows(openings, pcb.mask.Set0)
	brush1.IterateContours(openings, pcb.mask.Set1)
}

func (pcb *PCB) cutStencil(c *Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Pads
	pads := c.Pads.Apply(t)
	brush02.IterateContours(pads, pcb.stencil.Set1)
}

func (pcb *PCB) bitmapTransform() transform.Transform {
	return transform.Move(pcb.width/2, pcb.height/2).Scale(resolution, resolution)
}

func (pcb *PCB) SaveFiles() error {
	return util.GoAll([]func() error{
		pcb.SaveEtch,
		pcb.SaveMask,
		pcb.SaveMaskBottom,
		pcb.SaveStencil,
		pcb.SaveOverview,
	})
}

func (pcb *PCB) SaveOverview() error {
	filename := pcb.savePath + "overview.png"

	image := bitmap.NewBitmapsImage(
		[]*bitmap.Bitmap{pcb.copper, pcb.fr4, pcb.mask, pcb.silk, pcb.stencil},
		[][2]color.Color{
			{color.RGBA{G: 0x40, B: 0x10, A: 0xFF}, color.RGBA{R: 0xC0, G: 0x60, A: 0xFF}},
			{color.RGBA{}, color.RGBA{G: 0xFF, A: 0xFF}},
			{color.RGBA{}, color.RGBA{R: 0x80, G: 0x80, B: 0xFF, A: 0xA0}},
			{color.RGBA{}, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xA0}},
			{color.RGBA{}, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}},
		},
		true,
	)
	if err := util.SavePng(filename, image); err != nil {
		return err
	}

	return nil
}
