package eda

import (
	"image"
	"image/color"
	"temnok/lab/bitmap"
	"temnok/lab/contour"
	"temnok/lab/font"
	"temnok/lab/geom"
	"temnok/lab/path"
	"temnok/lab/shape"
	"temnok/lab/util"
)

type XY = geom.XY

type PCB struct {
	scaleTransform geom.Transform
	Transform      geom.Transform
	TrackWidth     float64

	scale                  float64
	cuts, maskHoles, holes [][]XY
	cu, mask, silk         *bitmap.Bitmap
}

func NewPCB(w, h float64) *PCB {
	const scale = 100

	wi, hi := int(w*scale), int(h*scale)

	return &PCB{
		scaleTransform: geom.ScaleK(scale).MoveXY(w/2, h/2),
		Transform:      geom.Identity(),
		TrackWidth:     0.2,
		scale:          scale,
		cu:             bitmap.NewBitmap(wi, hi),
		mask:           bitmap.NewBitmap(wi, hi),
		silk:           bitmap.NewBitmap(wi, hi),
	}
}

func (pcb *PCB) scaled() geom.Transform {
	return pcb.scaleTransform.Multiply(pcb.Transform)
}

func (pcb *PCB) With(block func()) {
	saved := *pcb
	block()

	cuts, holes, maskHoles := pcb.cuts, pcb.holes, pcb.maskHoles
	*pcb = saved
	pcb.cuts, pcb.holes, pcb.maskHoles = cuts, holes, maskHoles
}

func (pcb *PCB) Track(t geom.Transform, points ...XY) {
	brush := shape.Circle(int(pcb.TrackWidth * pcb.scale))
	brush.IterateContour(contour.Lines(t.Points(points)), pcb.scaled(), pcb.cu.SetRow1)
}

func (pcb *PCB) Pad(t geom.Transform, padContours ...[]XY) {
	shape.IterateContoursRows(padContours, pcb.scaled().Multiply(t), pcb.cu.SetRow1)
	pcb.MaskPad(t, padContours...)
}

func (pcb *PCB) MaskPad(t geom.Transform, padContours ...[]XY) {
	pcb.MaskContour(t, 0.1, padContours...)
}

func (pcb *PCB) MaskContour(t geom.Transform, w float64, contour ...[]XY) {
	brush := shape.Circle(int(w * pcb.scale))
	brush.IterateContours(contour, pcb.scaled().Multiply(t), pcb.mask.SetRow1)
}

func (pcb *PCB) MaskHole(t geom.Transform, contour []XY) {
	pcb.MaskContour(geom.Identity(), 0.2, contour)
	pcb.maskHoles = append(pcb.maskHoles, pcb.Transform.Points(contour))
}

func (pcb *PCB) SilkContour(t geom.Transform, w float64, contour []XY) {
	brush := shape.Circle(int(w * pcb.scale))
	brush.IterateContour(contour, pcb.scaled().Multiply(t), pcb.silk.SetRow1)
}

func (pcb *PCB) SilkText(t geom.Transform, height float64, text string) {
	brush := shape.Circle(int(font.Normal * height * pcb.scale))

	for i, c := range text {
		if c := int(c); c < len(font.Paths) {
			t := pcb.scaled().Multiply(t).ScaleK(height).MoveXY(float64(i)*font.Width, 0.4)
			brush.IterateContours(font.Paths[c], t, pcb.silk.SetRow1)
		}
	}
}

func (pcb *PCB) Cut(t geom.Transform, contour []XY) {
	pcb.cuts = append(pcb.cuts, t.Points(contour))

	brush := shape.Circle(int(0.1 * pcb.scale))

	path.IterateDotted(contour, pcb.scaled().Multiply(t), int(0.2*pcb.scale), func(x, y int) {
		brush.IterateRowsXY(x, y, pcb.mask.SetRow1)
	})
}

func (pcb *PCB) Hole(t geom.Transform, hole []XY) {
	pcb.holes = append(pcb.holes, pcb.Transform.Multiply(t).Points(hole))

	w := contour.Size(hole).X
	k := (w + 0.2) / w
	shape.IterateContourRows(hole, pcb.scaled().Multiply(t).ScaleK(k), pcb.cu.SetRow0)
}

func (pcb *PCB) SaveFiles() error {
	//util.SaveTmpPng("cu.png", pcb.cu.ToImage(color.Black, color.White))
	//util.SaveTmpPng("mask.png", pcb.mask.ToImage(color.White, color.Black))
	//util.SaveTmpPng("silk.png", pcb.silk.ToImage(color.White, color.Black))

	pcb.technologicalParts()

	util.SaveTmpPng("overview.png", &util.MultiImage{
		Images: []image.Image{
			pcb.cu.ToImage(color.RGBA{0, 0x40, 0x10, 0xFF}, color.RGBA{0xFF, 0x80, 0, 0x7F}),
			pcb.mask.ToImage(color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0x40}),
			pcb.silk.ToImage(color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0x40}),
		},
	})

	if err := pcb.SaveEtch("tmp/etch.lbrn"); err != nil {
		return err
	}

	if err := pcb.SaveMask("tmp/mask.lbrn"); err != nil {
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

	holder := contour.Circle(1)
	maskHole := contour.Circle(0.65)

	for _, h := range holders {
		t := geom.Move(h)

		pcb.Hole(t, holder)
		pcb.MaskPad(t, holder)

		pcb.MaskHole(t, maskHole)
	}

	key := []XY{
		{0.2, -0.2},
		{0.15, 0.2},
		{-0.2, -0.15},
		{0.2, -0.2},
	}
	t := geom.Move(XY{-16.3, 21.3})
	pcb.Track(t, key...)
	pcb.SilkContour(t, 0.2, contour.Lines(key))
}
