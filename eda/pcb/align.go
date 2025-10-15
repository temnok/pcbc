// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"errors"
	"github.com/temnok/pcbc/bitmap"
	"github.com/temnok/pcbc/bitmap/image"
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/pcb/config"
	"github.com/temnok/pcbc/lbrn"
	"image/color"
	"strconv"
)

const (
	alignImageIndex = 0
	alignCutIndex   = 1
)

var (
	alignImageSettings = []*lbrn.CutSetting{
		{
			Type:     "Image",
			Name:     &lbrn.Param{Value: "Outline"},
			Index:    &lbrn.Param{Value: strconv.Itoa(alignImageIndex)},
			Priority: &lbrn.Param{Value: strconv.Itoa(alignImageIndex)},

			MaxPower:    &lbrn.Param{Value: "10"},
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
			Speed:        &lbrn.Param{Value: "1100"},

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
	var topCuts, bottomCuts []*lbrn.Shape

	board.Visit(func(c *eda.Component) {
		if len(c.AlignCuts) == 0 {
			return
		}

		t := c.Transform.Multiply(config.LbrnCenterMove())

		for _, cut := range c.AlignCuts {
			if c.Bottom {
				bottomCuts = append(bottomCuts, lbrn.NewPath(alignCutIndex, t, cut))
			} else {
				topCuts = append(topCuts, lbrn.NewPath(alignCutIndex, t, cut))
			}
		}
	})

	if len(topCuts)+len(bottomCuts) == 0 {
		return nil
	}

	bounds := lbrn.NewRoundRect(alignCutIndex, config.LbrnCenterMove(),
		config.Width+5, config.Height+5, 2)

	bm := mask.Clone()
	bm.Or(silk)
	bmImage := image.NewSingle(bm, color.Transparent, color.Black)

	top := &lbrn.LightBurnProject{
		UIPrefs:       lbrn.UIPrefsDefaults,
		CutSetting:    alignCutSettings,
		CutSettingImg: alignImageSettings,
		Shape: append(
			topCuts,
			bounds,
			lbrn.NewBitmapShapeFromImage(alignImageIndex, config.LbrnBitmapScale(), bmImage),
		),
	}

	bottom := &lbrn.LightBurnProject{
		UIPrefs:    lbrn.UIPrefsDefaults,
		CutSetting: alignCutSettings,
		Shape: append(
			bottomCuts,
			bounds,
		),
	}

	return errors.Join(
		top.SaveToFile(config.SavePath+"0-align-top.lbrn"),
		bottom.SaveToFile(config.SavePath+"0-align-bottom.lbrn"),
	)
}
