package eda

import (
	"image/color"
	"math"
	"temnok/lab/bitmap"
	"temnok/lab/eda/lib"
	"temnok/lab/font"
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

	cuts, holes                             Paths
	maskHoles                               Paths
	stencilCuts, stencilHoles, stencilMarks Paths
	placerCuts, placerHoles, placerMarks    Paths

	copper, mask, silk, stencil *bitmap.Bitmap
}

func NewPCB(w, h float64) *PCB {
	const scale = 100

	wi, hi := int(w*scale), int(h*scale)

	return &PCB{
		width:      w,
		height:     h,
		resolution: scale,
		trackWidth: 0.2,
		copper:     bitmap.NewBitmap(wi, hi),
		mask:       bitmap.NewBitmap(wi, hi),
		silk:       bitmap.NewBitmap(wi, hi),
		stencil:    bitmap.NewBitmap(wi, hi),
	}
}

func (pcb *PCB) bitmapTransform() geom.Transform {
	return geom.ScaleK(pcb.resolution).MoveXY(pcb.width/2, pcb.height/2)
}

func (pcb *PCB) Cut(contour Path) {
	pcb.cuts = append(pcb.cuts, contour)

	brush := shape.Circle(int(0.1 * pcb.resolution))

	contour.Transform(pcb.bitmapTransform()).Jump(int(0.2*pcb.resolution), func(x, y int) {
		brush.IterateRowsXY(x, y, pcb.mask.Set1)
	})
}

func (pcb *PCB) StencilCut(contours ...Path) {
	pcb.stencilCuts = append(pcb.stencilCuts, contours...)
}

func (pcb *PCB) PlacerCut(contours ...Path) {
	pcb.placerCuts = append(pcb.placerCuts, contours...)
}

func (pcb *PCB) Hole(hole Path) {
	pcb.HoleNoStencil(hole)
	pcb.StencilHole(hole)
}

func (pcb *PCB) StencilHole(hole ...Path) {
	pcb.stencilHoles = append(pcb.stencilHoles, hole...)
}

func (pcb *PCB) HoleNoStencil(hole Path) {
	pcb.holes = append(pcb.holes, hole)

	shape.IterateContourRows(hole, pcb.bitmapTransform(), pcb.copper.Set0)

	brush := shape.Circle(int(0.2 * pcb.resolution))
	brush.IterateContour(hole, pcb.bitmapTransform(), pcb.copper.Set0)
}

func (pcb *PCB) Track(points []XY) {
	brush := shape.Circle(int(pcb.trackWidth * pcb.resolution))
	brush.IterateContour(path.Lines(points), pcb.bitmapTransform(), pcb.copper.Set1)
}

func (pcb *PCB) Component(c *lib.Component) {
	pcb.Pad(c.Pads...)

	pcb.placerHoles = append(pcb.placerHoles, c.Placer)
}

func (pcb *PCB) Pad(padContours ...Path) {
	pcb.PadNoStencil(padContours...)
	pcb.stencilHoles = append(pcb.stencilHoles, padContours...)
}

func (pcb *PCB) PadNoStencil(padContours ...Path) {
	shape.IterateContoursRows(padContours, pcb.bitmapTransform(), pcb.copper.Set1)
	pcb.MaskPad(padContours...)
}

func (pcb *PCB) MaskPad(padContours ...Path) {
	pcb.MaskContour(0.1, padContours...)
}

func (pcb *PCB) MaskContour(w float64, contour ...Path) {
	brush := shape.Circle(int(w * pcb.resolution))
	brush.IterateContours(contour, pcb.bitmapTransform(), pcb.mask.Set1)
}

func (pcb *PCB) MaskHole(contour Path) {
	pcb.MaskContour(0.2, contour)
	pcb.maskHoles = append(pcb.maskHoles, contour)
}

func (pcb *PCB) SilkContour(w float64, contour Path) {
	brush := shape.Circle(int(w * pcb.resolution))
	brush.IterateContour(contour, pcb.bitmapTransform(), pcb.silk.Set1)
}

func (pcb *PCB) SilkText(t geom.Transform, text string) {
	scale := min(math.Sqrt(t.I.X*t.I.X+t.I.Y*t.I.Y), math.Sqrt(t.J.X*t.J.X+t.J.Y*t.J.Y))
	brush := shape.Circle(int(font.Bold * scale * pcb.resolution))

	t1 := pcb.bitmapTransform().Multiply(t)
	brush.IterateContours(font.StringPaths(text, font.AlignCenter), t1, pcb.silk.Set1)
}

func (pcb *PCB) SaveFiles(path string) error {
	pcb.technologicalParts()

	if err := pcb.SaveEtch(path + "etch.lbrn"); err != nil {
		return err
	}

	if err := pcb.SaveMask(path + "mask.lbrn"); err != nil {
		return err
	}

	if err := pcb.SaveStencil(path + "stencil.lbrn"); err != nil {
		return err
	}

	if err := pcb.SavePlacer(path + "placer.lbrn"); err != nil {
		return err
	}

	image := bitmap.NewBitmapsImage(
		[]*bitmap.Bitmap{pcb.copper, pcb.mask, pcb.silk, pcb.stencil},
		[][2]color.Color{
			{color.RGBA{0, 0x40, 0x10, 0xFF}, color.RGBA{0xFF, 0x50, 0, 0xFF}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0x80, 0x80, 0xFF, 0xA0}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0x60}},
			{color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}},
		},
		true,
	)
	if err := util.SavePng(path+"overview.png", image); err != nil {
		return err
	}

	return nil
}

func (pcb *PCB) technologicalParts() {
	holders = []XY{
		{-15, 20},
		{15, 20},
		{-15, -20},
		{15, -20},
	}

	holder := path.Circle(1)
	maskHole := path.Circle(0.65)

	for _, h := range holders {
		t := geom.Move(h)

		pcb.Hole(holder.Transform(t))
		pcb.MaskPad(holder.Transform(t))

		pcb.MaskHole(maskHole.Transform(t))

		pcb.placerHoles = append(pcb.placerHoles, holder.Transform(t))
	}

	key := path.Points{
		{0.5, -0.5},
		{0.3, 0.5},
		{-0.5, -0.3},
		{0.5, -0.5},
	}
	t := geom.MoveXY(-16.4, 21.4)
	pcb.Track(key.Transform(t))
	pcb.SilkContour(0.2, path.Lines(key).Transform(t))

	mark := path.Lines(key).Transform(t)
	pcb.stencilMarks = append(pcb.stencilMarks, mark)
	pcb.placerMarks = append(pcb.placerMarks, mark)
}
