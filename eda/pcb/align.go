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
)

var (
	alignImageIndex = 0
	alignCutIndex   = 1

	alignImageSettings = []*lbrn.CutSetting{
		{
			Type:     "Image",
			Name:     &lbrn.Param{Value: "Silk"},
			Index:    &lbrn.Param{Value: strconv.Itoa(alignImageIndex)},
			Priority: &lbrn.Param{Value: strconv.Itoa(alignImageIndex)},

			MaxPower:    &lbrn.Param{Value: "5"},
			QPulseWidth: &lbrn.Param{Value: "200"},
			Frequency:   &lbrn.Param{Value: "20000"},

			NumPasses: &lbrn.Param{Value: "1"},
			Speed:     &lbrn.Param{Value: "800"},
			Interval:  &lbrn.Param{Value: "0.02"},
			DPI:       &lbrn.Param{Value: "1270"},

			CrossHatch: &lbrn.Param{Value: "1"},
			Angle:      &lbrn.Param{Value: "-90"},

			UseDotCorrection: &lbrn.Param{Value: "1"},
			DotWidth:         &lbrn.Param{Value: "0.05"},
		},
	}

	alignCutSettings = []*lbrn.CutSetting{
		{
			Type:     "Cut",
			Name:     &lbrn.Param{Value: "Cut"},
			Index:    &lbrn.Param{Value: strconv.Itoa(alignCutIndex)},
			Priority: &lbrn.Param{Value: strconv.Itoa(alignCutIndex)},

			MaxPower:    &lbrn.Param{Value: "90"},
			QPulseWidth: &lbrn.Param{Value: "200"},
			Frequency:   &lbrn.Param{Value: "20000"},

			NumPasses:    &lbrn.Param{Value: "1"},
			GlobalRepeat: &lbrn.Param{Value: "50"},
			Speed:        &lbrn.Param{Value: "1000"},

			SubLayer: &lbrn.SubLayer{
				Type:  "Cut",
				Index: "1",

				MaxPower: &lbrn.Param{Value: "0.1"},
				Speed:    &lbrn.Param{Value: "200"},

				QPulseWidth: &lbrn.Param{Value: "200"},
				Frequency:   &lbrn.Param{Value: "20000"},
			},
		},
	}
)

func SaveAlign(config *config.Config, board *eda.Component, mask, silk *bitmap.Bitmap) error {
	bm := mask.Clone()
	bm.Or(silk)

	board.Visit(func(c *eda.Component) {
	})

	filename := config.SavePath + "0-align-top.lbrn"

	bitmapImage := image.NewSingle(bm, color.Transparent, color.Black)

	p := &lbrn.LightBurnProject{
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSettingImg: alignImageSettings,
		CutSetting:    alignCutSettings,
		Shape: []*lbrn.Shape{
			lbrn.NewBitmapShapeFromImage(alignImageIndex, config.LbrnBitmapScale(), bitmapImage),
		},
	}

	return p.SaveToFile(filename)
}
