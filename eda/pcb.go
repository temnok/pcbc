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

type PCB struct {
	Transform  geom.Transform
	TrackWidth float64

	scale          float64
	cu, mask, silk *bitmap.Bitmap
}

func NewPCB(w, h float64) *PCB {
	const scale = 100

	wi, hi := int(w*scale), int(h*scale)

	return &PCB{
		Transform:  geom.ScaleK(scale).MoveXY(w/2, h/2),
		TrackWidth: 0.2,
		scale:      scale,
		cu:         bitmap.NewBitmap(wi, hi),
		mask:       bitmap.NewBitmap(wi, hi),
		silk:       bitmap.NewBitmap(wi, hi),
	}
}

func (pcb *PCB) With(block func()) {
	saved := *pcb
	block()
	*pcb = saved
}

func (pcb *PCB) Track(points ...geom.XY) {
	brush := shape.Circle(int(pcb.TrackWidth * pcb.scale))
	brush.IterateContour(contour.Lines(points), pcb.Transform, pcb.cu.SetRow1)
}

func (pcb *PCB) Pad(t geom.Transform, padContours ...[]geom.XY) {
	t = pcb.Transform.Multiply(t)
	shape.IterateContoursRows(padContours, t, pcb.cu.SetRow1)

	brush := shape.Circle(int(0.1 * pcb.scale))
	brush.IterateContours(padContours, t, pcb.mask.SetRow1)
}

func (pcb *PCB) SilkContour(t geom.Transform, w float64, contour []geom.XY) {
	brush := shape.Circle(int(w * pcb.scale))
	brush.IterateContour(contour, pcb.Transform.Multiply(t), pcb.silk.SetRow1)
}

func (pcb *PCB) SilkText(t geom.Transform, height float64, text string) {
	brush := shape.Circle(int(font.Normal * height * pcb.scale))

	for i, c := range text {
		if c := int(c); c < len(font.Paths) {
			t := pcb.Transform.Multiply(t).ScaleK(height).MoveXY(float64(i)*font.Width, 0.4)
			brush.IterateContours(font.Paths[c], t, pcb.silk.SetRow1)
		}
	}
}

func (pcb *PCB) Hole(t geom.Transform, contour []geom.XY) {
	shape.IterateContourRows(contour, pcb.Transform.Multiply(t), pcb.cu.SetRow0)
}

func (pcb *PCB) Cut(contour []geom.XY) {
	brush := shape.Circle(int(0.1 * pcb.scale))

	path.IterateDotted(contour, pcb.Transform, int(0.2*pcb.scale), func(x, y int) {
		brush.IterateRowsXY(x, y, pcb.mask.SetRow1)
	})
}

func (pcb *PCB) SaveFiles() {
	util.SaveTmpPng("cu.png", pcb.cu.ToImage(color.Black, color.White))
	util.SaveTmpPng("mask.png", pcb.mask.ToImage(color.White, color.Black))
	util.SaveTmpPng("silk.png", pcb.silk.ToImage(color.White, color.Black))

	util.SaveTmpPng("overview.png", &util.MultiImage{
		Images: []image.Image{
			pcb.cu.ToImage(color.RGBA{0, 0x40, 0x10, 0xFF}, color.RGBA{0xFF, 0x80, 0, 0x7F}),
			pcb.mask.ToImage(color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0x40}),
			pcb.silk.ToImage(color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0x40}),
		},
	})
}
