// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"fmt"
	"strings"
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/util"
)

func Process(initialConfig *config.Config, defaultComponent *eda.Component) error {
	if initialConfig == nil {
		initialConfig = config.Default()
	}

	config := *initialConfig
	setMissingConfigSize(&config, defaultComponent)

	config.SavePath = strings.ReplaceAll(initialConfig.SavePath, "{}", fmt.Sprint(defaultComponent.Layer))

	component := &eda.Component{
		TrackWidth: config.TrackWidth,
		ClearWidth: config.ClearWidth,
		Components: eda.Components{
			defaultComponent,
		},
	}

	var copper, mask, silk, stencil *bitmap.Bitmap

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
			var e error
			stencil, e = SaveStencil(&config, component)
			return e
		},
	)
	if err != nil {
		return err
	}

	return saveOverview(&config, copper, mask, silk, stencil)
}

func setMissingConfigSize(config *config.Config, component *eda.Component) {
	if config.Width > 0 && config.Height > 0 {
		return
	}

	w, h := component.Size()
	w, h = w+1, h+1

	if config.Width <= 0 {
		config.Width = w
	}

	if config.Height <= 0 {
		config.Height = h
	}
}
