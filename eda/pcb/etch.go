// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/shape"
)

type Param = lbrn.Param

func (pcb *PCB) SaveEtch() error {
	pcb.component.Visit(func(component *eda.Component) {
		pcb.removeCopper(component)
	})

	pcb.component.Visit(func(component *eda.Component) {
		pcb.addCopper(component)
	})

	filename := pcb.SavePath + "etch.lbrn"
	im := image.NewSingle(pcb.copper, color.Black, color.White)
	bm := lbrn.NewBase64Bitmap(im)

	p := &lbrn.LightBurnProject{
		CutSettingImg: []*lbrn.CutSetting{
			{
				Type:     "Image",
				Name:     Param{Value: "Etch"},
				Index:    Param{Value: "0"},
				Priority: Param{Value: "0"},

				MaxPower:    Param{Value: "20"},
				QPulseWidth: Param{Value: "200"},
				Frequency:   Param{Value: "20000"},

				Speed:            Param{Value: "600"},
				Interval:         Param{Value: "0.01"},
				DPI:              Param{Value: "2540"},
				UseDotCorrection: Param{Value: "1"},
				DotWidth:         Param{Value: "0.05"},

				Negative: Param{Value: "1"},
			},
			{
				Type:     "Image",
				Name:     Param{Value: "Clean 1"},
				Index:    Param{Value: "1"},
				Priority: Param{Value: "1"},

				MaxPower:    Param{Value: "50"},
				QPulseWidth: Param{Value: "2"},
				Frequency:   Param{Value: "280000"},

				Speed:            Param{Value: "2000"},
				Interval:         Param{Value: "0.01"},
				DPI:              Param{Value: "2540"},
				UseDotCorrection: Param{Value: "1"},
				DotWidth:         Param{Value: "0.15"},

				Negative: Param{Value: "1"},
			},
			{
				Type:     "Image",
				Name:     Param{Value: "Clean 2"},
				Index:    Param{Value: "3"},
				Priority: Param{Value: "3"},

				MaxPower:    Param{Value: "50"},
				QPulseWidth: Param{Value: "2"},
				Frequency:   Param{Value: "280000"},

				Speed:            Param{Value: "2000"},
				Interval:         Param{Value: "0.01"},
				DPI:              Param{Value: "2540"},
				UseDotCorrection: Param{Value: "1"},
				DotWidth:         Param{Value: "0.15"},

				Angle:    Param{Value: "90"},
				Negative: Param{Value: "1"},
			},
		},
		CutSetting: []*lbrn.CutSetting{
			{
				Type:     "Cut",
				Name:     Param{Value: "FR4 Cut"},
				Index:    Param{Value: "2"},
				Priority: Param{Value: "2"},

				Speed:        Param{Value: "100"},
				GlobalRepeat: Param{Value: "50"},

				MaxPower:    Param{Value: "90"},
				QPulseWidth: Param{Value: "200"},
				Frequency:   Param{Value: "20000"},

				TabsEnabled: Param{Value: "1"},
				TabSize:     Param{Value: "0.1"},
			},
		},
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShape(0, pcb.lbrnBitmapScale(), bm),
			lbrn.NewBitmapShape(1, pcb.lbrnBitmapScale(), bm),
			lbrn.NewBitmapShape(3, pcb.lbrnBitmapScale(), bm),
		},
	}

	pcb.addEtchShapes(p)

	return p.SaveToFile(filename)
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

func (pcb *PCB) addEtchShapes(p *lbrn.LightBurnProject) {
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
