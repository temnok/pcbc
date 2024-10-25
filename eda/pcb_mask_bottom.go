package eda

import (
	"image/color"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/transform"
)

func (pcb *PCB) SaveMaskBottom() error {
	filename := pcb.savePath + "mask-bottom.lbrn"

	mask := pcb.maskB.ToImage(color.White, color.Black)

	bitmapTransform := transform.Scale(-1, 1).ScaleK(1/pcb.resolution).
		Move(pcb.lbrnCenter.X, pcb.lbrnCenter.Y)

	p := lbrn.LightBurnProject{
		CutSettingImg: maskCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmap(1, bitmapTransform, mask),
			lbrn.NewBitmap(2, bitmapTransform, mask),
		},
	}

	return p.SaveToFile(filename)
}
