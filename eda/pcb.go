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
	width, height, resolution float64
	trackWidth                float64

	cuts                                    [][]XY
	holes, maskHoles                        [][]XY
	stencilCuts, stencilHoles, stencilMarks [][]XY
	cu, mask, silk                          *bitmap.Bitmap
}

func NewPCB(w, h float64) *PCB {
	const scale = 100

	wi, hi := int(w*scale), int(h*scale)

	return &PCB{
		width:      w,
		height:     h,
		resolution: scale,
		trackWidth: 0.2,
		cu:         bitmap.NewBitmap(wi, hi),
		mask:       bitmap.NewBitmap(wi, hi),
		silk:       bitmap.NewBitmap(wi, hi),
	}
}

func (pcb *PCB) bitmapTransform() geom.Transform {
	return geom.ScaleK(pcb.resolution).MoveXY(pcb.width/2, pcb.height/2)
}

func (pcb *PCB) Cut(contour []XY) {
	pcb.cuts = append(pcb.cuts, contour)

	brush := shape.Circle(int(0.1 * pcb.resolution))

	path.IterateDotted(contour, pcb.bitmapTransform(), int(0.2*pcb.resolution), func(x, y int) {
		brush.IterateRowsXY(x, y, pcb.mask.SetRow1)
	})
}

func (pcb *PCB) StencilCut(contours ...[]XY) {
	pcb.stencilCuts = append(pcb.stencilCuts, contours...)
}

func (pcb *PCB) Hole(hole []XY) {
	pcb.HoleNoStencil(hole)
	pcb.StencilHole(hole)
}

func (pcb *PCB) StencilHole(hole ...[]XY) {
	pcb.stencilHoles = append(pcb.stencilHoles, hole...)
}

func (pcb *PCB) StencilMark(mark ...[]XY) {
	pcb.stencilMarks = append(pcb.stencilMarks, mark...)
}

func (pcb *PCB) HoleNoStencil(hole []XY) {
	pcb.holes = append(pcb.holes, hole)

	//w := contour.Size(hole).X
	//k := (w + 0.2) / w
	//shape.IterateContourRows(hole, pcb.bitmapTransform(geom.Identity()).ScaleK(k), pcb.cu.SetRow0)

	shape.IterateContourRows(hole, pcb.bitmapTransform(), pcb.cu.SetRow0)
}

func (pcb *PCB) Track(points []XY) {
	brush := shape.Circle(int(pcb.trackWidth * pcb.resolution))
	brush.IterateContour(contour.Lines(points), pcb.bitmapTransform(), pcb.cu.SetRow1)
}

func (pcb *PCB) Pad(padContours ...[]XY) {

	pcb.PadNoStencil(padContours...)
	pcb.stencilHoles = append(pcb.stencilHoles, padContours...)
}

func (pcb *PCB) PadNoStencil(padContours ...[]XY) {
	shape.IterateContoursRows(padContours, pcb.bitmapTransform(), pcb.cu.SetRow1)
	pcb.MaskPad(padContours...)
}

func (pcb *PCB) MaskPad(padContours ...[]XY) {
	pcb.MaskContour(0.1, padContours...)
}

func (pcb *PCB) MaskContour(w float64, contour ...[]XY) {
	brush := shape.Circle(int(w * pcb.resolution))
	brush.IterateContours(contour, pcb.bitmapTransform(), pcb.mask.SetRow1)
}

func (pcb *PCB) MaskHole(contour []XY) {
	pcb.MaskContour(0.2, contour)
	pcb.maskHoles = append(pcb.maskHoles, contour)
}

func (pcb *PCB) SilkContour(w float64, contour []XY) {
	brush := shape.Circle(int(w * pcb.resolution))
	brush.IterateContour(contour, pcb.bitmapTransform(), pcb.silk.SetRow1)
}

func (pcb *PCB) SilkText(t geom.Transform, height float64, text string) {
	brush := shape.Circle(int(font.Normal * height * pcb.resolution))

	for i, c := range text {
		if c := int(c); c < len(font.Paths) {
			t := pcb.bitmapTransform().Multiply(t).ScaleK(height).MoveXY(float64(i)*font.Width, 0.4)
			brush.IterateContours(font.Paths[c], t, pcb.silk.SetRow1)
		}
	}
}

func (pcb *PCB) SaveFiles(path string) error {
	//util.SaveTmpPng("cu.png", pcb.cu.ToImage(color.Black, color.White))
	//util.SaveTmpPng("mask.png", pcb.mask.ToImage(color.White, color.Black))
	//util.SaveTmpPng("silk.png", pcb.silk.ToImage(color.White, color.Black))

	pcb.technologicalParts()

	util.SavePng(path+"overview.png", &util.MultiImage{
		Images: []image.Image{
			pcb.cu.ToImage(color.RGBA{0, 0x40, 0x10, 0xFF}, color.RGBA{0xFF, 0x80, 0, 0x7F}),
			pcb.mask.ToImage(color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0x40}),
			pcb.silk.ToImage(color.RGBA{0, 0, 0, 0}, color.RGBA{0xFF, 0xFF, 0xFF, 0x40}),
		},
	})

	if err := pcb.SaveEtch(path + "etch.lbrn"); err != nil {
		return err
	}

	if err := pcb.SaveMask(path + "mask.lbrn"); err != nil {
		return err
	}

	if err := pcb.SaveStencil(path + "stencil.lbrn"); err != nil {
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

		pcb.Hole(t.Points(holder))
		pcb.MaskPad(t.Points(holder))

		pcb.MaskHole(t.Points(maskHole))
	}

	key := []XY{
		{0.25, -0.25},
		{0.2, 0.25},
		{-0.25, -0.2},
		{0.25, -0.25},
	}
	t := geom.Move(XY{-16.3, 21.3})
	pcb.Track(t.Points(key))
	pcb.SilkContour(0.2, t.Points(contour.Lines(key)))
	pcb.StencilMark(t.Points(contour.Lines(key)))
}
