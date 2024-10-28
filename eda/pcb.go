package eda

import (
	"image/color"
	"temnok/pcbc/bitmap"
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

	savePath         string
	component        *Component
	nonflatComponent *Component

	fr4, copper, mask, maskB, silk, stencil *bitmap.Bitmap
}

const (
	resolution         = 100.0 // pixels per mm
	clearBrushDiameter = 0.5
	extraCopper        = 0.05 // compensate copper lost during etching
)

var (
	brush1  = shape.Circle(int(0.1 * resolution))
	brush02 = shape.Circle(int(0.02 * resolution))
)

func GeneratePCB(component *Component) error {
	return NewPCB(component).SaveFiles()
}

func NewPCB(component *Component) *PCB {
	flatComponent := component.Flatten()
	width, height := flatComponent.Size()
	width, height = width+1, height+1

	wi, hi := int(width*resolution), int(height*resolution)

	pcb := &PCB{
		width:  width,
		height: height,

		nonflatComponent: component,

		trackWidth: 0.25,

		lbrnCenter: path.Point{55, 55},
		savePath:   "pcb/",

		fr4:     bitmap.NewBitmap(wi, hi),
		copper:  bitmap.NewBitmap(wi, hi),
		mask:    bitmap.NewBitmap(wi, hi),
		maskB:   bitmap.NewBitmap(wi, hi),
		silk:    bitmap.NewBitmap(wi, hi),
		stencil: bitmap.NewBitmap(wi, hi),
	}

	pcb.copper.Invert()

	pcb.setComponent(flatComponent)

	return pcb
}

func (pcb *PCB) setComponent(c *Component) {
	pcb.component = c

	pcb.processBoard()
	pcb.processMask()
	pcb.processStencil()
}

func (pcb *PCB) processBoard() {
	pcb.removeCopper()
	pcb.addCopper()
	pcb.cutBoard()
}

func (pcb *PCB) processMask() {
	pcb.addMarks()
	pcb.addOpenings()
}

func (pcb *PCB) removeCopper() {
	bt := pcb.bitmapTransform()

	// Clears
	clears := pcb.component.Clears.Apply(bt)
	shape.IterateContoursRows(clears, pcb.copper.Set0)

	// Pads
	pads := pcb.component.Pads.Apply(bt)
	clearBrush := shape.Circle(int(clearBrushDiameter * resolution))
	clearBrush.IterateContours(pads, pcb.copper.Set0)

	// Non-ground tracks
	for brushW, tracks := range pcb.component.Tracks {
		if brushW == 0 {
			brushW = pcb.trackWidth
		}
		brush := shape.Circle(int((brushW + clearBrushDiameter) * resolution))
		brush.IterateContours(tracks.Apply(bt), pcb.copper.Set0)
	}

	cutClearBrush := shape.Circle(int((clearBrushDiameter/2 - extraCopper) * resolution))

	// Holes
	holes := pcb.component.Holes.Apply(bt)
	cutClearBrush.IterateContours(holes, pcb.copper.Set0)

	// Cuts
	cuts := pcb.component.Cuts.Apply(bt)
	cutClearBrush.IterateContours(cuts, pcb.copper.Set0)
}

func (pcb *PCB) addCopper() {
	bt := pcb.bitmapTransform()

	extraCopperBrush := shape.Circle(int(extraCopper * resolution))

	// Pads
	pads := pcb.component.Pads.Apply(bt)
	shape.IterateContoursRows(pads, pcb.copper.Set1)
	extraCopperBrush.IterateContours(pads, pcb.copper.Set1)

	// Tracks
	for brushW, tracks := range (path.Strokes{}).Append(pcb.component.Tracks, pcb.component.GroundTracks) {
		if brushW == 0 {
			brushW = pcb.trackWidth
		}
		brush := shape.Circle(int((brushW + extraCopper) * resolution))
		brush.IterateContours(tracks.Apply(bt), pcb.copper.Set1)
	}
}

func (pcb *PCB) cutBoard() {
	bt := pcb.bitmapTransform()

	// Holes
	holes := pcb.component.Holes.Apply(bt)
	brush02.IterateContours(holes, pcb.fr4.Set1)

	// Cuts
	cuts := pcb.component.Cuts.Apply(bt)
	brush02.IterateContours(cuts, pcb.fr4.Set1)
}

func (pcb *PCB) addMarks() {
	bt := pcb.bitmapTransform()

	// Marks
	for brushW, marks := range pcb.component.Marks {
		brush := shape.Circle(int(brushW * resolution))
		brush.IterateContours(marks.Apply(bt), pcb.silk.Set1)
	}
}

func (pcb *PCB) addOpenings() {
	bt := pcb.bitmapTransform()

	// Pads
	pads := pcb.component.Pads.Apply(bt)
	brush1.IterateContours(pads, pcb.mask.Set1)

	// Holes
	holes := pcb.component.Holes.Apply(bt)
	brush1.IterateContours(holes, pcb.mask.Set1)
	brush1.IterateContours(holes, pcb.maskB.Set1)

	// Cuts
	cuts := pcb.component.Cuts.Apply(bt)
	cuts.Jump(int(0.2*resolution), func(x, y int) {
		brush1.IterateRowsXY(x, y, pcb.mask.Set1)
		brush1.IterateRowsXY(x, y, pcb.maskB.Set1)
	})

	// Openings
	openings := pcb.component.Openings.Apply(bt)
	shape.IterateContoursRows(openings, pcb.mask.Set0)
	brush1.IterateContours(openings, pcb.mask.Set1)
}

func (pcb *PCB) processStencil() {
	bt := pcb.bitmapTransform()

	// Pads
	pads := pcb.component.Pads.Apply(bt)
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
