// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda"
	"temnok/pcbc/font"
	"temnok/pcbc/shape"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
)

type PCB struct {
	component *eda.Component

	Width, Height float64
	PixelsPerMM   float64

	ExtraCopperWidth float64
	CopperClearWidth float64
	MaskCutWidth     float64
	OverviewCutWidth float64

	LbrnCenterX, LbrnCenterY float64

	SavePath string

	copper, mask, silk *bitmap.Bitmap
}

func New(component *eda.Component) *PCB {
	width, height := component.Size()
	width, height = width+1, height+1

	return &PCB{
		component: &eda.Component{
			TrackWidth: 0.25,
			Components: eda.Components{component},
		},

		Width:       width,
		Height:      height,
		PixelsPerMM: 100,

		ExtraCopperWidth: 0.05,
		CopperClearWidth: 0.25,
		MaskCutWidth:     0.1,
		OverviewCutWidth: 0.02,

		LbrnCenterX: 55,
		LbrnCenterY: 55,

		SavePath: "out/",
	}
}

func Generate(component *eda.Component) error {
	return Process(component).SaveFiles()
}

func Process(component *eda.Component) *PCB {
	return New(component).Process()
}

func (pcb *PCB) Process() *PCB {
	wi, hi := pcb.bitmapSize()

	pcb.copper = bitmap.New(wi, hi)
	pcb.mask = bitmap.New(wi, hi)
	pcb.silk = bitmap.New(wi, hi)

	pcb.processPass1()
	pcb.processPass2()

	return pcb
}

func (pcb *PCB) bitmapSize() (int, int) {
	return int(pcb.Width * pcb.PixelsPerMM), int(pcb.Height * pcb.PixelsPerMM)
}

func (pcb *PCB) processPass1() {
	pcb.component.Visit(func(component *eda.Component) {
		pcb.removeCopper(component)
		pcb.addSilk(component)
		pcb.cutMask1(component)
	})
}

func (pcb *PCB) processPass2() {
	pcb.component.Visit(func(component *eda.Component) {
		pcb.addCopper(component)
		pcb.cutMask2(component)
	})
}

func (pcb *PCB) removeCopper(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Clears
	shape.ForEachRow(c.Clears, t, pcb.copper.Set1)

	clearWidth := 2 * (pcb.CopperClearWidth - pcb.ExtraCopperWidth)

	// Pads
	clearBrush := shape.Circle(int(clearWidth * pcb.PixelsPerMM))
	clearBrush.ForEachPathsPixel(c.Pads, t, pcb.copper.Set1)

	// Non-ground tracks
	brush := shape.Circle(int((c.TrackWidth + clearWidth) * pcb.PixelsPerMM))
	brush.ForEachPathsPixel(c.Tracks, t, pcb.copper.Set1)

	clearBrush = shape.Circle(int(pcb.CopperClearWidth * pcb.PixelsPerMM))
	clearBrush.ForEachPathsPixel(c.Cuts, t, pcb.copper.Set1)
	clearBrush.ForEachPathsPixel(c.Holes, t, pcb.copper.Set1)
	clearBrush.ForEachPathsPixel(c.Perforations, t, pcb.copper.Set1)
}

func (pcb *PCB) addCopper(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Pads
	shape.ForEachRow(c.Pads, t, pcb.copper.Set0)

	extraCopperBrush := shape.Circle(int(pcb.ExtraCopperWidth * pcb.PixelsPerMM))
	extraCopperBrush.ForEachPathsPixel(c.Pads, t, pcb.copper.Set0)

	// Tracks
	brush := shape.Circle(int((c.TrackWidth + pcb.ExtraCopperWidth) * pcb.PixelsPerMM))
	brush.ForEachPathsPixel(c.Tracks, t, pcb.copper.Set0)
	brush.ForEachPathsPixel(c.GroundTracks, t, pcb.copper.Set0)
}

func (pcb *PCB) addSilk(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	// Marks:
	brushW := font.Bold * font.WeightScale(t)
	brush := shape.Circle(int(brushW))
	brush.ForEachPathsPixel(c.Marks, t, pcb.silk.Set1)
}

func (pcb *PCB) cutMask1(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	brush := shape.Circle(int(pcb.MaskCutWidth * pcb.PixelsPerMM))

	// Pads
	brush.ForEachPathsPixel(c.Pads, t, pcb.mask.Set1)

	// Cuts
	c.Cuts.ForEachPixelDist(t, int(2*pcb.MaskCutWidth*pcb.PixelsPerMM), func(x, y int) {
		brush.ForEachRowWithOffset(x, y, pcb.mask.Set1)
	})

	// Holes
	brush.ForEachPathsPixel(c.Holes, t, pcb.mask.Set1)

	// Perforations
	brush.ForEachPathsPixel(c.Perforations, t, pcb.mask.Set1)
}

func (pcb *PCB) cutMask2(c *eda.Component) {
	t := c.Transform.Multiply(pcb.bitmapTransform())

	brush := shape.Circle(int(pcb.MaskCutWidth * pcb.PixelsPerMM))

	// Openings
	shape.ForEachRow(c.Openings, t, pcb.mask.Set0)
	brush.ForEachPathsPixel(c.Openings, t, pcb.mask.Set1)
}

func (pcb *PCB) bitmapTransform() transform.T {
	return transform.Move(pcb.Width/2, pcb.Height/2).ScaleUniformly(pcb.PixelsPerMM)
}

func (pcb *PCB) lbrnCenterMove() transform.T {
	return transform.Move(pcb.LbrnCenterX, pcb.LbrnCenterY)
}

func (pcb *PCB) lbrnBitmapScale() transform.T {
	scale := 1.0 / pcb.PixelsPerMM
	return transform.Scale(scale, -scale).Multiply(pcb.lbrnCenterMove())
}

func (pcb *PCB) SaveFiles() error {
	return util.RunConcurrently(
		pcb.SaveEtch,
		pcb.SaveMask,
		pcb.SaveStencil,
		pcb.SaveOverview,
	)
}
