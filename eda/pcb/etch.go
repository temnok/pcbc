// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"image/color"
	"strconv"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/bitmap/image"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/lbrn"
	"temnok/pcbc/path"
	"temnok/pcbc/shape"
)

const (
	etchPassIndex  = 1
	cutPassIndex   = 2
	cleanPassIndex = 3
)

var etchBitmapSettings = []*lbrn.CutSetting{
	{
		Type:     "Image",
		Name:     &lbrn.Param{Value: "Etch Cu18"},
		Index:    &lbrn.Param{Value: strconv.Itoa(etchPassIndex)},
		Priority: &lbrn.Param{Value: strconv.Itoa(etchPassIndex)},
		Negative: &lbrn.Param{Value: "1"},

		MaxPower:    &lbrn.Param{Value: "30"},
		QPulseWidth: &lbrn.Param{Value: "80"},
		Frequency:   &lbrn.Param{Value: "40000"},

		NumPasses: &lbrn.Param{Value: "10"},
		Speed:     &lbrn.Param{Value: "800"},
		Interval:  &lbrn.Param{Value: "0.02"},
		DPI:       &lbrn.Param{Value: "1270"},

		Angle:            &lbrn.Param{Value: "-90"},
		CrossHatch:       &lbrn.Param{Value: "1"},
		UseDotCorrection: &lbrn.Param{Value: "1"},
		DotWidth:         &lbrn.Param{Value: "0.05"},
	},
}

func etchCutSettings(c *eda.Component) []*lbrn.CutSetting {
	var doOutput *lbrn.Param
	if c.CutsDisabled() {
		doOutput = &lbrn.Param{Value: "0"}
	}

	return []*lbrn.CutSetting{
		{
			Type:     "Cut",
			Name:     &lbrn.Param{Value: "Cut Board"},
			Index:    &lbrn.Param{Value: strconv.Itoa(cutPassIndex)},
			Priority: &lbrn.Param{Value: strconv.Itoa(cutPassIndex)},
			DoOutput: doOutput,

			MaxPower:    &lbrn.Param{Value: "90"},
			QPulseWidth: &lbrn.Param{Value: "200"},
			Frequency:   &lbrn.Param{Value: "20000"},

			NumPasses:    &lbrn.Param{Value: "1"},
			GlobalRepeat: &lbrn.Param{Value: "150"},
			Speed:        &lbrn.Param{Value: "600"},

			SubLayer: &lbrn.SubLayer{
				Type:  "Cut",
				Index: "1",

				MaxPower: &lbrn.Param{Value: "0.1"},
				Speed:    &lbrn.Param{Value: "200"},

				QPulseWidth: &lbrn.Param{Value: "200"},
				Frequency:   &lbrn.Param{Value: "20000"},
			},
		},
		{
			Type:     "Scan",
			Name:     &lbrn.Param{Value: "Clean"},
			Index:    &lbrn.Param{Value: strconv.Itoa(cleanPassIndex)},
			Priority: &lbrn.Param{Value: strconv.Itoa(cleanPassIndex)},

			MaxPower:    &lbrn.Param{Value: "90"},
			QPulseWidth: &lbrn.Param{Value: "1"},
			Frequency:   &lbrn.Param{Value: "650000"},

			CrossHatch: &lbrn.Param{Value: "1"},
			NumPasses:  &lbrn.Param{Value: "1"},
			Speed:      &lbrn.Param{Value: "6500"},
			Interval:   &lbrn.Param{Value: "0.01"},
			DPI:        &lbrn.Param{Value: "2540"},
		},
	}
}

func SaveEtch(config *config.Config, component *eda.Component) (*bitmap.Bitmap, error) {
	copper := bitmap.New(config.BitmapSizeInPixels())
	var cuts []*lbrn.Shape

	component.Visit(func(c *eda.Component) {
		removeEtchCopper(config, c, copper)
	})

	component.Visit(func(c *eda.Component) {
		addEtchCopper(config, c, copper)
		addEtchCuts(config, c, &cuts)
	})

	component.Visit(func(c *eda.Component) {
		removeViaCopper(config, c, copper)
	})

	filename := config.SavePath + "etch.lbrn"
	im := image.NewSingle(copper, color.Black, color.White)
	bm := lbrn.NewBase64Bitmap(im)

	p := &lbrn.LightBurnProject{
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSettingImg: etchBitmapSettings,
		CutSetting:    etchCutSettings(component),
		Shape: append([]*lbrn.Shape{
			lbrn.NewBitmapShape(etchPassIndex, config.LbrnBitmapScale(), bm),
		}, cuts...),
	}

	addCleanPasses(config, p)

	return copper, p.SaveToFile(filename)
}

func removeEtchCopper(config *config.Config, component *eda.Component, copper *bitmap.Bitmap) {
	if component.ClearNone {
		return
	}

	t := component.Transform.Multiply(config.BitmapTransform())

	clearWidth := 2 * (component.ClearWidth - config.ExtraCopperWidth)

	// Cuts
	if !component.CutsDisabled() {
		cutBrush := shape.Circle(int((clearWidth / 2) * config.PixelsPerMM))
		cutBrush.ForEachPathsPixel(component.Cuts, t, copper.Set1)
	}

	// Pads
	padBrush := shape.Circle(int(clearWidth * config.PixelsPerMM))
	padBrush.ForEachPathsPixel(component.Pads, t, copper.Set1)

	// Tracks
	trackBrush := shape.Circle(int((component.TracksWidth + clearWidth) * config.PixelsPerMM))
	trackBrush.ForEachPathsPixel(component.Tracks, t, copper.Set1)
}

func addEtchCopper(config *config.Config, component *eda.Component, copper *bitmap.Bitmap) {
	t := component.Transform.Multiply(config.BitmapTransform())

	// Pads
	shape.ForEachRow(component.Pads, t, copper.Set0)

	extraCopperBrush := shape.Circle(int(config.ExtraCopperWidth * config.PixelsPerMM))
	extraCopperBrush.ForEachPathsPixel(component.Pads, t, copper.Set0)

	// Tracks
	brush := shape.Circle(int((component.TracksWidth + config.ExtraCopperWidth) * config.PixelsPerMM))
	brush.ForEachPathsPixel(component.Tracks, t, copper.Set0)
}

func addEtchCuts(config *config.Config, component *eda.Component, cuts *[]*lbrn.Shape) {
	t := component.Transform.Multiply(config.LbrnCenterMove())

	for _, cut := range component.Vias {
		*cuts = append(*cuts, lbrn.NewPath(cutPassIndex, t, cut))
	}

	for _, cut := range component.Cuts {
		*cuts = append(*cuts, lbrn.NewPath(cutPassIndex, t, cut))
	}
}

func removeViaCopper(config *config.Config, component *eda.Component, copper *bitmap.Bitmap) {
	t := component.Transform.Multiply(config.BitmapTransform())

	// Vias
	shape.ForEachRow(component.Vias, t, copper.Set1)
}

func addCleanPasses(config *config.Config, p *lbrn.LightBurnProject) {
	t := config.LbrnCenterMove()
	boardBounds := path.Rect(config.Width, config.Height)
	p.Shape = append(p.Shape,
		lbrn.NewPath(cleanPassIndex, t, boardBounds),
	)
}
