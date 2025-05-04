// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/util"
)

func Process(initialConfig *config.Config, component *eda.Component) error {
	config := *initialConfig
	if config.Width == 0 || config.Height == 0 {
		w, h := component.Size()
		w, h = w+1, h+1
		if config.Width == 0 {
			config.Width = w
		}
		if config.Height == 0 {
			config.Height = h
		}
	}

	if component.TrackWidth == 0 {
		component = &eda.Component{
			TrackWidth: config.TrackWidth,
			Components: eda.Components{
				component,
			},
		}
	}

	var copper, mask, silk *bitmap.Bitmap

	err := util.RunConcurrently(
		func() error {
			var e error
			copper, e = SaveEtch(&config, component)
			return e
		},
		func() error {
			var e error
			mask, silk, e = SaveMask(&config, component)
			return e
		},
		func() error {
			return SaveStencil(&config, component)
		},
	)
	if err != nil {
		return err
	}

	return SaveOverview(&config, component, copper, mask, silk)
}

func ProcessWithDefaultConfig(component *eda.Component) error {
	return Process(config.Default(), component)
}
