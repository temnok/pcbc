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

	scale          float64
	cuts, holes    [][]XY
	cu, mask, silk *bitmap.Bitmap
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

	cuts, holes := pcb.cuts, pcb.holes
	*pcb = saved
	pcb.cuts, pcb.holes = cuts, holes
}

func (pcb *PCB) Track(points ...XY) {
	brush := shape.Circle(int(pcb.TrackWidth * pcb.scale))
	brush.IterateContour(contour.Lines(points), pcb.scaled(), pcb.cu.SetRow1)
}

func (pcb *PCB) Pad(t geom.Transform, padContours ...[]XY) {
	shape.IterateContoursRows(padContours, pcb.scaled().Multiply(t), pcb.cu.SetRow1)
	pcb.PadMask(t, padContours...)
}

func (pcb *PCB) PadMask(t geom.Transform, padContours ...[]XY) {
	brush := shape.Circle(int(0.1 * pcb.scale))
	brush.IterateContours(padContours, pcb.scaled().Multiply(t), pcb.mask.SetRow1)
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

func (pcb *PCB) Cut(contour []XY) {
	pcb.cuts = append(pcb.cuts, pcb.Transform.Points(contour))

	brush := shape.Circle(int(0.1 * pcb.scale))

	path.IterateDotted(contour, pcb.scaled(), int(0.2*pcb.scale), func(x, y int) {
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
	for _, h := range holders {
		t := geom.Move(h)

		pcb.Hole(t, holder)
		pcb.PadMask(t, holder)
	}

	key := []XY{
		{-0.4, 0.4},
		{0.4, 0.25},
		{-0.25, -0.4},
		{-0.4, 0.4},
	}
	t := geom.Move(XY{-13.5, 18.5})
	pcb.Track(t.Points(key)...)
	pcb.SilkContour(t, 0.2, contour.Lines(key))

	//for r := 1.5; r <= 3; r += 1.5 {
	//	pcb.PadMask(geom.Move(XY{-15.5, 0}), contour.Circle(r))
	//}
}
