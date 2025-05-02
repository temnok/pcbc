// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/shape"
)

type param = lbrn.Param

func (pcb *PCB) SaveEtch() (*bitmap.Bitmap, error) {
	copper := bitmap.New(pcb.bitmapSize())

	pcb.component.Visit(func(component *eda.Component) {
		removeEtchCopper(pcb, component, copper)
	})

	pcb.component.Visit(func(component *eda.Component) {
		addEtchCopper(pcb, component, copper)
	})

	filename := pcb.SavePath + "etch.lbrn"
	im := image.NewSingle(copper, color.Black, color.White)
	bm := lbrn.NewBase64Bitmap(im)

	p := &lbrn.LightBurnProject{
		CutSettingImg: []*lbrn.CutSetting{
			{
				Type:     "Image",
				Name:     param{Value: "Etch"},
				Index:    param{Value: "0"},
				Priority: param{Value: "0"},

				MaxPower:    param{Value: "20"},
				QPulseWidth: param{Value: "200"},
				Frequency:   param{Value: "20000"},

				Speed:            param{Value: "600"},
				Interval:         param{Value: "0.01"},
				DPI:              param{Value: "2540"},
				UseDotCorrection: param{Value: "1"},
				DotWidth:         param{Value: "0.05"},

				Negative: param{Value: "1"},
			},
			{
				Type:     "Image",
				Name:     param{Value: "Clean 1"},
				Index:    param{Value: "1"},
				Priority: param{Value: "1"},

				MaxPower:    param{Value: "50"},
				QPulseWidth: param{Value: "2"},
				Frequency:   param{Value: "280000"},

				Speed:            param{Value: "2000"},
				Interval:         param{Value: "0.01"},
				DPI:              param{Value: "2540"},
				UseDotCorrection: param{Value: "1"},
				DotWidth:         param{Value: "0.15"},

				Negative: param{Value: "1"},
			},
			{
				Type:     "Image",
				Name:     param{Value: "Clean 2"},
				Index:    param{Value: "3"},
				Priority: param{Value: "3"},

				MaxPower:    param{Value: "50"},
				QPulseWidth: param{Value: "2"},
				Frequency:   param{Value: "280000"},

				Speed:            param{Value: "2000"},
				Interval:         param{Value: "0.01"},
				DPI:              param{Value: "2540"},
				UseDotCorrection: param{Value: "1"},
				DotWidth:         param{Value: "0.15"},

				Angle:    param{Value: "90"},
				Negative: param{Value: "1"},
			},
		},
		CutSetting: []*lbrn.CutSetting{
			{
				Type:     "Cut",
				Name:     param{Value: "FR4 Cut"},
				Index:    param{Value: "2"},
				Priority: param{Value: "2"},

				Speed:        param{Value: "100"},
				GlobalRepeat: param{Value: "50"},

				MaxPower:    param{Value: "90"},
				QPulseWidth: param{Value: "200"},
				Frequency:   param{Value: "20000"},

				TabsEnabled: param{Value: "1"},
				TabSize:     param{Value: "0.1"},
			},
		},
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShape(0, pcb.lbrnBitmapScale(), bm),
			lbrn.NewBitmapShape(1, pcb.lbrnBitmapScale(), bm),
			lbrn.NewBitmapShape(3, pcb.lbrnBitmapScale(), bm),
		},
	}

	pcb.addEtchCuts(p)

	return copper, p.SaveToFile(filename)
}

func removeEtchCopper(config *PCB, component *eda.Component, copper *bitmap.Bitmap) {
	t := component.Transform.Multiply(config.bitmapTransform())

	// Clears
	shape.ForEachRow(component.Clears, t, copper.Set1)

	clearWidth := 2 * (config.CopperClearWidth - config.ExtraCopperWidth)

	// Pads
	clearBrush := shape.Circle(int(clearWidth * config.PixelsPerMM))
	clearBrush.ForEachPathsPixel(component.Pads, t, copper.Set1)

	// Non-ground tracks
	brush := shape.Circle(int((component.TrackWidth + clearWidth) * config.PixelsPerMM))
	brush.ForEachPathsPixel(component.Tracks, t, copper.Set1)

	clearBrush = shape.Circle(int(config.CopperClearWidth * config.PixelsPerMM))
	clearBrush.ForEachPathsPixel(component.Cuts, t, copper.Set1)
	clearBrush.ForEachPathsPixel(component.Holes, t, copper.Set1)
	clearBrush.ForEachPathsPixel(component.Perforations, t, copper.Set1)
}

func addEtchCopper(config *PCB, component *eda.Component, copper *bitmap.Bitmap) {
	t := component.Transform.Multiply(config.bitmapTransform())

	// Pads
	shape.ForEachRow(component.Pads, t, copper.Set0)

	extraCopperBrush := shape.Circle(int(config.ExtraCopperWidth * config.PixelsPerMM))
	extraCopperBrush.ForEachPathsPixel(component.Pads, t, copper.Set0)

	// Tracks
	brush := shape.Circle(int((component.TrackWidth + config.ExtraCopperWidth) * config.PixelsPerMM))
	brush.ForEachPathsPixel(component.Tracks, t, copper.Set0)
	brush.ForEachPathsPixel(component.GroundTracks, t, copper.Set0)
}

func (pcb *PCB) addEtchCuts(p *lbrn.LightBurnProject) {
	pcb.component.Visit(func(component *eda.Component) {
		t := component.Transform.Multiply(pcb.lbrnCenterMove())

		for _, cut := range component.Cuts {
			p.Shape = append(p.Shape, lbrn.NewPathWithTabs(2, t, cut))
		}

		for _, hole := range component.Holes {
			p.Shape = append(p.Shape, lbrn.NewPath(2, t, hole))
		}

		for _, perforation := range component.Perforations {
			p.Shape = append(p.Shape, lbrn.NewPath(2, t, perforation))
		}
	})
}
