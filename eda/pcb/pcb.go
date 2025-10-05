// Copyright © 2025 Alex Temnok. All rights reserved.

package pcb

import (
	"temnok/pcbc/bitmap"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
)

var fileNamePrefix = map[bool]string{
	false: "1-",
	true:  "2-",
}

func Process(config *config.Config, component *eda.Component) error {

	componentFront := &eda.Component{
		CutsWidth:           0.1,
		CutsPerforationStep: 0.17,
		MarksWidth:          0.13,
		TracksWidth:         0.2,
		ClearWidth:          0.2,

		Nested: eda.Components{
			component,
		},
	}

	componentBack := componentFront.Arrange(transform.MirrorX)

	var copper1, copper2, mask1, mask2, silk1, silk2, stencil *bitmap.Bitmap

	err := util.RunConcurrently(
		func() error {
			var e error
			copper1, e = SaveEtch(config, componentFront, false)
			return e
		},
		func() error {
			var e error
			copper2, e = SaveEtch(config, componentBack, true)
			return e
		},
		func() error {
			var e error
			mask1, silk1, e = SaveMask(config, componentFront, false)
			return e
		},
		func() error {
			var e error
			mask2, silk2, e = SaveMask(config, componentBack, true)
			return e
		},
		func() error {
			var e error
			stencil, e = SaveStencil(config, componentFront)
			return e
		},
	)
	if err != nil {
		return err
	}

	return util.RunConcurrently(
		//func() error {
		//	return SaveAlign(config, componentFront, mask1, silk1)
		//},

		func() error {
			return saveOverview(config, "1-overview.png", copper1, mask1, silk1, stencil)
		},

		func() error {
			return saveOverview(config, "2-overview.png", copper2, mask2, silk2, nil)
		},
	)
}
