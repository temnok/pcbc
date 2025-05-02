// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda"
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

	return pcb
}

func (pcb *PCB) bitmapSize() (int, int) {
	return int(pcb.Width * pcb.PixelsPerMM), int(pcb.Height * pcb.PixelsPerMM)
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
	err := util.RunConcurrently(
		pcb.SaveEtch,
		pcb.SaveMask,
		pcb.SaveStencil,
	)
	if err != nil {
		return err
	}

	return pcb.SaveOverview()
}
